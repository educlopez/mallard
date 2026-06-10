package backup

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// DefaultRetentionCount is the number of unpinned backup batches the GC keeps.
// Pinned batches are never pruned and do not count toward this limit.
const DefaultRetentionCount = 5

// pinMarker is the sentinel file written inside a batch dir to pin it.
const pinMarker = ".pinned"

// archiveName is the tar.gz snapshot written alongside the raw copies for each
// batch. It mirrors the raw files and allows compressed round-trip restore.
const archiveName = "archive.tar.gz"

// casDirName is the content-addressed store directory under BackupsRoot. It
// holds one copy of each distinct file content keyed by SHA-256. Backup batches
// reference content here via hardlinks, so identical files snapshotted across
// many sessions occupy a single inode on disk (dedup) while each batch path
// still restores byte-identically. The leading underscore keeps it sorted
// before timestamped batch dirs and excluded from batch enumeration.
const casDirName = "_cas"

// Entry describes a single file or directory that was backed up before being
// overwritten by mallard.
type Entry struct {
	Agent        string `json:"agent"`
	Kind         string `json:"kind"`
	OriginalPath string `json:"original_path"`
	BackupPath   string `json:"backup_path"`
	Sha256       string `json:"sha256,omitempty"`
}

// Manifest is the JSON manifest written alongside each backup batch.
type Manifest struct {
	Timestamp time.Time `json:"timestamp"`
	Entries   []Entry   `json:"entries"`
	// Archive is the relative name of the tar.gz snapshot for this batch, if any.
	Archive string `json:"archive,omitempty"`
	// Pinned records whether the batch is pinned (never pruned). The on-disk
	// .pinned marker file is the source of truth; this field is informational.
	Pinned bool `json:"pinned,omitempty"`
}

// Session groups all snapshots taken under a single timestamped backup dir.
type Session struct {
	mu      sync.Mutex
	rootDir string
	stamp   string
	entries []Entry

	// dedupHits / dedupMisses count, for the lifetime of the session, how many
	// snapshotted regular files matched content already present in the CAS
	// (hits) versus how many had to be freshly stored (misses). Used by tests
	// and informational only.
	dedupHits   int
	dedupMisses int
}

// DedupStats reports content-addressed dedup activity for this session: hits
// (file content already present in the store, not re-stored) and misses
// (content stored for the first time).
func (s *Session) DedupStats() (hits, misses int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.dedupHits, s.dedupMisses
}

// NewSession opens (lazily) a new backup session under ~/.mallard/backups/<RFC3339>.
// The directory is NOT created until the first Snapshot call.
func NewSession() (*Session, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("home dir: %w", err)
	}
	stamp := time.Now().UTC().Format("20060102T150405Z")
	root := filepath.Join(home, ".mallard", "backups", stamp)
	return &Session{rootDir: root, stamp: stamp}, nil
}

// NewSessionAt creates a backup session rooted at an explicit directory.
// Intended for testing; production code uses NewSession.
func NewSessionAt(rootDir string) (*Session, error) {
	stamp := time.Now().UTC().Format("20060102T150405Z")
	return &Session{rootDir: rootDir, stamp: stamp}, nil
}

// Root returns the backup directory path. Empty until first snapshot.
func (s *Session) Root() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.entries) == 0 {
		return ""
	}
	return s.rootDir
}

// Count returns how many entries have been snapshotted.
func (s *Session) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.entries)
}

// DiscardEntry removes the last entry whose BackupPath matches e.BackupPath from
// the session entry list. It is used to roll back a snapshot when a subsequent
// operation (e.g. RemoveAll) fails and the snapshot should not count as a backup
// hit. The backup file on disk is left in place (harmless orphan).
func (s *Session) DiscardEntry(e Entry) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := len(s.entries) - 1; i >= 0; i-- {
		if s.entries[i].BackupPath == e.BackupPath {
			s.entries = append(s.entries[:i], s.entries[i+1:]...)
			return
		}
	}
}

// Snapshot copies the file or directory at originalPath into the session's
// backup root, organized by agent/kind. Symlinks are NOT followed: a symlink
// is recorded with its target string instead of copying through.
func (s *Session) Snapshot(agentID, kind, originalPath string) (Entry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	info, err := os.Lstat(originalPath)
	if err != nil {
		return Entry{}, fmt.Errorf("lstat %s: %w", originalPath, err)
	}

	dstDir := filepath.Join(s.rootDir, agentID, kind)
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return Entry{}, fmt.Errorf("mkdir %s: %w", dstDir, err)
	}
	dst := filepath.Join(dstDir, filepath.Base(originalPath))

	entry := Entry{
		Agent:        agentID,
		Kind:         kind,
		OriginalPath: originalPath,
		BackupPath:   dst,
	}

	switch {
	case info.Mode()&os.ModeSymlink != 0:
		target, rerr := os.Readlink(originalPath)
		if rerr != nil {
			return Entry{}, fmt.Errorf("readlink %s: %w", originalPath, rerr)
		}
		if werr := os.WriteFile(dst+".symlink", []byte(target), 0o644); werr != nil {
			return Entry{}, fmt.Errorf("write symlink record: %w", werr)
		}
		entry.BackupPath = dst + ".symlink"

	case info.IsDir():
		if err := copyTree(originalPath, dst); err != nil {
			return Entry{}, err
		}

	default:
		sum, hit, err := s.storeDeduped(originalPath, dst)
		if err != nil {
			return Entry{}, err
		}
		entry.Sha256 = sum
		if hit {
			s.dedupHits++
		} else {
			s.dedupMisses++
		}
	}

	s.entries = append(s.entries, entry)
	return entry, nil
}

// storeDeduped content-addresses the regular file at src into the shared CAS
// (under BackupsRoot/_cas/<sha>) and links the batch destination dst to it.
//
// On a dedup HIT (content already in the CAS) no second copy of the bytes is
// written: dst becomes a hardlink to the existing CAS object. On a MISS the
// content is copied into the CAS once, then dst is hardlinked to it. If
// hardlinking is unsupported (e.g. cross-device or a filesystem without link
// support) it falls back to a plain copy so the batch dir is always usable for
// restore. Returns the content SHA-256 and whether it was a dedup hit.
func (s *Session) storeDeduped(src, dst string) (sum string, hit bool, err error) {
	sum, err = hashFile(src)
	if err != nil {
		return "", false, err
	}

	casDir := filepath.Join(filepath.Dir(s.rootDir), casDirName)
	if err := os.MkdirAll(casDir, 0o755); err != nil {
		return "", false, fmt.Errorf("mkdir cas: %w", err)
	}
	casPath := filepath.Join(casDir, sum)

	if _, statErr := os.Stat(casPath); statErr == nil {
		hit = true // content already stored — reference it, do not re-store.
	} else if os.IsNotExist(statErr) {
		if _, cerr := copyFile(src, casPath); cerr != nil {
			return "", false, cerr
		}
	} else {
		return "", false, fmt.Errorf("stat cas object: %w", statErr)
	}

	// Link the batch destination to the CAS object so identical content shares
	// one inode. Remove any stale dst first so the link call is deterministic.
	_ = os.Remove(dst)
	if lerr := os.Link(casPath, dst); lerr != nil {
		// Hardlinks may be unavailable (cross-device, unsupported fs). Fall back
		// to a plain copy. Prefer copying from the CAS object; if it has been
		// concurrently removed (e.g. by a GC race), copy directly from src.
		if _, cerr := copyFile(casPath, dst); cerr != nil {
			if !os.IsNotExist(cerr) {
				return "", false, cerr
			}
			// CAS object disappeared (concurrent GC). Copy from original src;
			// no longer a dedup hit since the content was re-read, not shared.
			hit = false
			if _, cerr2 := copyFile(src, dst); cerr2 != nil {
				return "", false, cerr2
			}
		}
	}
	return sum, hit, nil
}

// hashFile returns the hex SHA-256 of the file at path without copying it.
func hashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("hash %s: %w", path, err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Finalize writes manifest.json and runs the keep-latest-5 garbage collector.
// No-op if no snapshots were taken.
func (s *Session) Finalize() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.entries) == 0 {
		return nil
	}

	// Write a compressed tar.gz mirror of the raw backup copies. This is a
	// best-effort companion to the raw files (which remain the primary restore
	// source). A failure here must not lose the manifest, so it is non-fatal.
	archive := ""
	if err := writeSessionArchive(s.rootDir, s.entries); err == nil {
		archive = archiveName
	}

	m := Manifest{Timestamp: time.Now().UTC(), Entries: s.entries, Archive: archive}
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal manifest: %w", err)
	}
	if err := os.WriteFile(filepath.Join(s.rootDir, "manifest.json"), data, 0o644); err != nil {
		return fmt.Errorf("write manifest: %w", err)
	}

	return gcOldBackups(filepath.Dir(s.rootDir), DefaultRetentionCount)
}

// writeSessionArchive packs every raw backup file under rootDir (excluding the
// manifest and the archive itself) into a single tar.gz. Paths inside the
// archive are relative to rootDir so extraction can faithfully reproduce the
// batch layout. Symlink-record (.symlink) files are stored verbatim like any
// regular file. Symlink directory entries are skipped.
func writeSessionArchive(rootDir string, entries []Entry) error {
	if len(entries) == 0 {
		return nil
	}
	f, err := os.Create(filepath.Join(rootDir, archiveName))
	if err != nil {
		return fmt.Errorf("create archive: %w", err)
	}
	defer f.Close()

	gw := gzip.NewWriter(f)
	tw := tar.NewWriter(gw)

	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, werr error) error {
		if werr != nil {
			return werr
		}
		if d.IsDir() {
			return nil
		}
		// Skip symlink directory entries; regular .symlink record files are
		// plain files and will be included by the regular-file branch below.
		if d.Type()&os.ModeSymlink != 0 {
			return nil
		}
		base := filepath.Base(path)
		if base == archiveName || base == "manifest.json" || base == pinMarker {
			return nil
		}
		rel, rerr := filepath.Rel(rootDir, path)
		if rerr != nil {
			return rerr
		}
		return addFileToTar(tw, path, filepath.ToSlash(rel))
	})
	if err != nil {
		_ = tw.Close()
		_ = gw.Close()
		return err
	}
	if err := tw.Close(); err != nil {
		return fmt.Errorf("close tar: %w", err)
	}
	if err := gw.Close(); err != nil {
		return fmt.Errorf("close gzip: %w", err)
	}
	return nil
}

// addFileToTar writes the tar header and streams the contents of path into tw.
// The header size comes from a Stat of the already-open handle so it cannot
// disagree with the bytes actually copied. Using a helper function rather than
// an inline open/defer avoids deferring a Close inside a loop.
func addFileToTar(tw *tar.Writer, path, relName string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return err
	}
	hdr := &tar.Header{
		Typeflag: tar.TypeReg,
		Name:     relName,
		Mode:     int64(info.Mode().Perm()),
		Size:     info.Size(),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	if _, err := io.CopyN(tw, f, info.Size()); err != nil {
		return fmt.Errorf("archive %s: %w", relName, err)
	}
	return nil
}

// ExtractArchive unpacks the tar.gz batch archive at archivePath into destDir,
// recreating the relative file layout. It rejects entries that would escape
// destDir. Returns the absolute paths of the files written.
func ExtractArchive(archivePath, destDir string) ([]string, error) {
	f, err := os.Open(archivePath)
	if err != nil {
		return nil, fmt.Errorf("open archive: %w", err)
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("gzip reader: %w", err)
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	var written []string
	cleanBase := filepath.Clean(destDir) + string(filepath.Separator)
	for {
		hdr, rerr := tr.Next()
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			return nil, fmt.Errorf("read tar entry: %w", rerr)
		}
		if hdr.Typeflag != tar.TypeReg {
			continue
		}
		dst := filepath.Join(destDir, filepath.FromSlash(hdr.Name))
		if !strings.HasPrefix(filepath.Clean(dst), cleanBase) {
			return nil, fmt.Errorf("archive entry %q escapes destination", hdr.Name)
		}
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return nil, err
		}
		out, oerr := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(hdr.Mode).Perm())
		if oerr != nil {
			return nil, fmt.Errorf("create %q: %w", dst, oerr)
		}
		if _, cerr := io.Copy(out, tr); cerr != nil {
			_ = out.Close()
			return nil, fmt.Errorf("write %q: %w", dst, cerr)
		}
		if cerr := out.Close(); cerr != nil {
			return nil, cerr
		}
		written = append(written, dst)
	}
	return written, nil
}

// SetPinned pins or unpins a backup batch identified by its full stamp
// (directory name under BackupsRoot). A pinned batch is never pruned by the
// keep-latest GC. Pinning writes a .pinned marker file and updates the
// manifest's Pinned field when a manifest is present.
func SetPinned(stamp string, pinned bool) error {
	root, err := BackupsRoot()
	if err != nil {
		return err
	}
	dir := filepath.Join(root, stamp)
	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("backup %q not found", stamp)
	}

	marker := filepath.Join(dir, pinMarker)
	if pinned {
		if err := os.WriteFile(marker, []byte(time.Now().UTC().Format(time.RFC3339)+"\n"), 0o644); err != nil {
			return fmt.Errorf("write pin marker: %w", err)
		}
	} else {
		if err := os.Remove(marker); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("remove pin marker: %w", err)
		}
	}

	// Best-effort: keep the manifest's Pinned field in sync. Absence of a
	// manifest is not an error — the marker file is authoritative.
	if m, lerr := LoadManifest(dir); lerr == nil {
		m.Pinned = pinned
		if data, merr := json.MarshalIndent(m, "", "  "); merr == nil {
			_ = os.WriteFile(filepath.Join(dir, "manifest.json"), data, 0o644)
		}
	}
	return nil
}

// IsPinned reports whether the backup batch dir is pinned (has a .pinned marker).
func IsPinned(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, pinMarker))
	return err == nil
}

func copyFile(src, dst string) (string, error) {
	in, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("open %s: %w", src, err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", fmt.Errorf("create %s: %w", dst, err)
	}
	defer out.Close()

	h := sha256.New()
	if _, err := io.Copy(io.MultiWriter(out, h), in); err != nil {
		return "", fmt.Errorf("copy %s -> %s: %w", src, dst, err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func copyTree(srcRoot, dstRoot string) error {
	return filepath.Walk(srcRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, rerr := filepath.Rel(srcRoot, path)
		if rerr != nil {
			return rerr
		}
		target := filepath.Join(dstRoot, rel)

		// Lstat to avoid following symlinks inside the tree.
		linfo, lerr := os.Lstat(path)
		if lerr != nil {
			return lerr
		}
		switch {
		case linfo.Mode()&os.ModeSymlink != 0:
			t, rrerr := os.Readlink(path)
			if rrerr != nil {
				return rrerr
			}
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			return os.WriteFile(target+".symlink", []byte(t), 0o644)
		case linfo.IsDir():
			return os.MkdirAll(target, 0o755)
		default:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			_, err := copyFile(path, target)
			return err
		}
	})
}

// BackupsRoot returns the parent directory that holds every backup batch.
func BackupsRoot() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home dir: %w", err)
	}
	return filepath.Join(home, ".mallard", "backups"), nil
}

// Summary describes a single backup batch on disk.
type Summary struct {
	Timestamp  string
	Dir        string
	EntryCount int
	TotalBytes int64
	ByAgent    map[string]int
	Pinned     bool
}

// ListBackups returns every backup batch under ~/.mallard/backups, newest first.
func ListBackups() ([]Summary, error) {
	root, err := BackupsRoot()
	if err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var stamps []string
	for _, e := range entries {
		if e.IsDir() && e.Name() != casDirName {
			stamps = append(stamps, e.Name())
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(stamps)))

	out := make([]Summary, 0, len(stamps))
	for _, stamp := range stamps {
		dir := filepath.Join(root, stamp)
		m, err := LoadManifest(dir)
		s := Summary{Timestamp: stamp, Dir: dir, ByAgent: map[string]int{}, Pinned: IsPinned(dir)}
		if err != nil {
			out = append(out, s)
			continue
		}
		s.EntryCount = len(m.Entries)
		for _, en := range m.Entries {
			s.ByAgent[en.Agent]++
			s.TotalBytes += entryBytes(en.BackupPath)
		}
		out = append(out, s)
	}
	return out, nil
}

// entryBytes returns the on-disk size of a backup entry. For files it is the
// file size; for directories it is the recursive sum of contained file sizes.
// Errors are silently treated as zero — list output should never fail just
// because one stale entry on disk is unreadable.
func entryBytes(path string) int64 {
	info, err := os.Lstat(path)
	if err != nil {
		return 0
	}
	if !info.IsDir() {
		return info.Size()
	}
	var total int64
	_ = filepath.Walk(path, func(_ string, fi os.FileInfo, werr error) error {
		if werr != nil || fi == nil {
			return nil
		}
		if !fi.IsDir() {
			total += fi.Size()
		}
		return nil
	})
	return total
}

// ResolveTimestamp accepts either a full stamp or a unique prefix and returns
// the canonical timestamp directory name.
func ResolveTimestamp(prefix string) (string, error) {
	root, err := BackupsRoot()
	if err != nil {
		return "", err
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("no backups found")
		}
		return "", err
	}

	var matches []string
	for _, e := range entries {
		if !e.IsDir() || e.Name() == casDirName {
			continue
		}
		if e.Name() == prefix {
			return e.Name(), nil
		}
		if strings.HasPrefix(e.Name(), prefix) {
			matches = append(matches, e.Name())
		}
	}
	switch len(matches) {
	case 0:
		return "", fmt.Errorf("no backup matches %q", prefix)
	case 1:
		return matches[0], nil
	default:
		sort.Strings(matches)
		return "", fmt.Errorf("ambiguous prefix: matches %s", strings.Join(matches, ", "))
	}
}

// LoadManifest reads manifest.json from a backup batch directory.
func LoadManifest(dir string) (*Manifest, error) {
	data, err := os.ReadFile(filepath.Join(dir, "manifest.json"))
	if err != nil {
		return nil, fmt.Errorf("read manifest: %w", err)
	}
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("parse manifest: %w", err)
	}
	return &m, nil
}

// RestoreClass classifies what a restore would do for a single entry.
type RestoreClass string

const (
	RestoreRestore RestoreClass = "restore" // copy backup over the target
	RestoreRelink  RestoreClass = "relink"  // target is a mallard symlink; safe to drop
	RestoreSkip    RestoreClass = "skip"    // target was modified by user; refuse to clobber
	RestoreFailed  RestoreClass = "failed"  // restore attempted but failed (sha mismatch, IO)
)

// RestoreItem is the per-entry result of restore planning or execution.
type RestoreItem struct {
	Entry  Entry
	Class  RestoreClass
	Reason string
	Err    error
}

// PlanRestore walks the manifest and classifies each entry against the current
// state of the filesystem. It performs no mutation.
func PlanRestore(m *Manifest, agentFilter string) []RestoreItem {
	out := make([]RestoreItem, 0, len(m.Entries))
	for _, e := range m.Entries {
		if agentFilter != "" && e.Agent != agentFilter {
			continue
		}
		item := RestoreItem{Entry: e, Class: RestoreRestore}

		info, lerr := os.Lstat(e.OriginalPath)
		switch {
		case lerr != nil && os.IsNotExist(lerr):
			// target gone — just restore
		case lerr != nil:
			item.Class = RestoreFailed
			item.Err = lerr
			item.Reason = "lstat failed"
		case info.Mode()&os.ModeSymlink != 0:
			item.Class = RestoreRelink
			item.Reason = "target is symlink"
		default:
			item.Class = RestoreSkip
			item.Reason = "target modified"
		}

		out = append(out, item)
	}
	return out
}

// ApplyRestore executes a previously-planned restore. Items already marked
// RestoreSkip or RestoreFailed are passed through unchanged.
func ApplyRestore(items []RestoreItem) []RestoreItem {
	for i := range items {
		it := &items[i]
		if it.Class == RestoreSkip || it.Class == RestoreFailed {
			continue
		}
		if it.Class == RestoreRelink {
			if err := os.Remove(it.Entry.OriginalPath); err != nil && !os.IsNotExist(err) {
				it.Class = RestoreFailed
				it.Err = fmt.Errorf("remove symlink: %w", err)
				continue
			}
		}
		if err := restoreOne(it.Entry); err != nil {
			it.Class = RestoreFailed
			it.Err = err
			continue
		}
		it.Class = RestoreRestore
	}
	return items
}

func restoreOne(e Entry) error {
	if strings.HasSuffix(e.BackupPath, ".symlink") {
		target, err := os.ReadFile(e.BackupPath)
		if err != nil {
			return fmt.Errorf("read symlink record: %w", err)
		}
		if err := os.MkdirAll(filepath.Dir(e.OriginalPath), 0o755); err != nil {
			return err
		}
		return os.Symlink(strings.TrimSpace(string(target)), e.OriginalPath)
	}

	info, err := os.Lstat(e.BackupPath)
	if err != nil {
		return fmt.Errorf("lstat backup: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(e.OriginalPath), 0o755); err != nil {
		return err
	}
	if info.IsDir() {
		return copyTree(e.BackupPath, e.OriginalPath)
	}

	sum, err := copyFileWithMode(e.BackupPath, e.OriginalPath, info.Mode())
	if err != nil {
		return err
	}
	if e.Sha256 != "" && sum != e.Sha256 {
		return fmt.Errorf("sha mismatch: got %s want %s", sum, e.Sha256)
	}
	return nil
}

func copyFileWithMode(src, dst string, mode os.FileMode) (string, error) {
	in, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("open %s: %w", src, err)
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode.Perm())
	if err != nil {
		return "", fmt.Errorf("create %s: %w", dst, err)
	}
	defer out.Close()

	h := sha256.New()
	if _, err := io.Copy(io.MultiWriter(out, h), in); err != nil {
		return "", fmt.Errorf("copy %s -> %s: %w", src, dst, err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// gcOldBackups prunes the oldest UNPINNED backup batches, keeping at most
// `keep` unpinned batches. Pinned batches (those with a .pinned marker) are
// never pruned and do not count toward the limit.
func gcOldBackups(parent string, keep int) error {
	entries, err := os.ReadDir(parent)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	var unpinned []string
	for _, e := range entries {
		if !e.IsDir() || e.Name() == casDirName {
			continue
		}
		if IsPinned(filepath.Join(parent, e.Name())) {
			continue // pinned — never prune, doesn't count toward keep
		}
		unpinned = append(unpinned, e.Name())
	}
	if len(unpinned) > keep {
		sort.Strings(unpinned)
		for _, d := range unpinned[:len(unpinned)-keep] {
			_ = os.RemoveAll(filepath.Join(parent, d))
		}
	}

	// Reap orphaned content-addressed objects. After batch dirs are removed,
	// any CAS object no longer hardlinked from a surviving batch has a link
	// count of 1 (only the CAS entry itself) and is safe to delete. This keeps
	// the store bounded as old batches age out.
	gcOrphanedCAS(filepath.Join(parent, casDirName))
	return nil
}

// gcOrphanedCAS deletes CAS objects whose only remaining hardlink is the CAS
// entry itself (link count <= 1), meaning no live backup batch references them.
// Best-effort: errors are ignored so GC never fails on an unreadable object.
func gcOrphanedCAS(casDir string) {
	entries, err := os.ReadDir(casDir)
	if err != nil {
		return
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		path := filepath.Join(casDir, e.Name())
		if nlink(path) <= 1 {
			_ = os.Remove(path)
		}
	}
}
