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
		sum, err := copyFile(originalPath, dst)
		if err != nil {
			return Entry{}, err
		}
		entry.Sha256 = sum
	}

	s.entries = append(s.entries, entry)
	return entry, nil
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
// regular file.
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

	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, werr error) error {
		if werr != nil {
			return werr
		}
		if info.IsDir() {
			return nil
		}
		base := filepath.Base(path)
		if base == archiveName || base == "manifest.json" || base == pinMarker {
			return nil
		}
		// Only archive regular files (the raw copies and .symlink records).
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}
		rel, rerr := filepath.Rel(rootDir, path)
		if rerr != nil {
			return rerr
		}
		data, rerr := os.ReadFile(path)
		if rerr != nil {
			return rerr
		}
		hdr := &tar.Header{
			Typeflag: tar.TypeReg,
			Name:     filepath.ToSlash(rel),
			Mode:     int64(info.Mode().Perm()),
			Size:     int64(len(data)),
		}
		if werr := tw.WriteHeader(hdr); werr != nil {
			return werr
		}
		_, werr = tw.Write(data)
		return werr
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
		if e.IsDir() {
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
		if !e.IsDir() {
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
		if !e.IsDir() {
			continue
		}
		if IsPinned(filepath.Join(parent, e.Name())) {
			continue // pinned — never prune, doesn't count toward keep
		}
		unpinned = append(unpinned, e.Name())
	}
	if len(unpinned) <= keep {
		return nil
	}
	sort.Strings(unpinned)
	for _, d := range unpinned[:len(unpinned)-keep] {
		_ = os.RemoveAll(filepath.Join(parent, d))
	}
	return nil
}
