package backup

import (
	"os"
	"path/filepath"
	"testing"
)

func write(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func exists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}

// TestArchiveRoundTrip writes raw backup files, packs them with
// writeSessionArchive, then extracts with ExtractArchive and asserts the
// content survives the tar.gz round-trip.
func TestArchiveRoundTrip(t *testing.T) {
	root := t.TempDir()
	write(t, filepath.Join(root, "claude", "skills", "foo"), "alpha")
	write(t, filepath.Join(root, "claude", "commands", "bar.md"), "beta")

	entries := []Entry{
		{Agent: "claude", Kind: "skills", BackupPath: filepath.Join(root, "claude", "skills", "foo")},
		{Agent: "claude", Kind: "commands", BackupPath: filepath.Join(root, "claude", "commands", "bar.md")},
	}
	if err := writeSessionArchive(root, entries); err != nil {
		t.Fatalf("writeSessionArchive: %v", err)
	}
	archive := filepath.Join(root, archiveName)
	if !exists(archive) {
		t.Fatalf("archive %s not created", archive)
	}

	dest := t.TempDir()
	written, err := ExtractArchive(archive, dest)
	if err != nil {
		t.Fatalf("ExtractArchive: %v", err)
	}
	if len(written) != 2 {
		t.Fatalf("extracted %d files, want 2", len(written))
	}

	got, err := os.ReadFile(filepath.Join(dest, "claude", "skills", "foo"))
	if err != nil {
		t.Fatalf("read extracted: %v", err)
	}
	if string(got) != "alpha" {
		t.Fatalf("extracted foo = %q, want alpha", string(got))
	}
	got2, _ := os.ReadFile(filepath.Join(dest, "claude", "commands", "bar.md"))
	if string(got2) != "beta" {
		t.Fatalf("extracted bar.md = %q, want beta", string(got2))
	}
}

// TestSessionFinalizeWritesArchive verifies a full session snapshot produces
// both the raw copy (restore source) AND a tar.gz, and that restore still works.
func TestSessionFinalizeWritesArchive(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	// A real file that will be "overwritten" — snapshot it first.
	orig := filepath.Join(t.TempDir(), "config.json")
	write(t, orig, "original-content")

	s, err := NewSession()
	if err != nil {
		t.Fatalf("NewSession: %v", err)
	}
	if _, err := s.Snapshot("claude", "skills", orig); err != nil {
		t.Fatalf("Snapshot: %v", err)
	}
	if err := s.Finalize(); err != nil {
		t.Fatalf("Finalize: %v", err)
	}

	// Compressed archive must exist alongside the raw copy + manifest.
	if !exists(filepath.Join(s.Root(), archiveName)) {
		t.Fatalf("Finalize did not write %s", archiveName)
	}

	m, err := LoadManifest(s.Root())
	if err != nil {
		t.Fatalf("LoadManifest: %v", err)
	}
	if m.Archive != archiveName {
		t.Fatalf("manifest.Archive = %q, want %q", m.Archive, archiveName)
	}

	// Now delete the original and restore from the raw backup copy.
	if err := os.Remove(orig); err != nil {
		t.Fatalf("remove orig: %v", err)
	}
	items := PlanRestore(m, "")
	results := ApplyRestore(items)
	if len(results) != 1 || results[0].Class != RestoreRestore {
		t.Fatalf("restore result = %+v", results)
	}
	got, err := os.ReadFile(orig)
	if err != nil {
		t.Fatalf("read restored: %v", err)
	}
	if string(got) != "original-content" {
		t.Fatalf("restored = %q, want original-content", string(got))
	}
}

// snapshotInto runs a single-file backup session rooted at the given home and
// returns the session so callers can inspect dedup stats and the batch root.
func snapshotInto(t *testing.T, home, orig string) *Session {
	t.Helper()
	t.Setenv("HOME", home)
	s, err := NewSession()
	if err != nil {
		t.Fatalf("NewSession: %v", err)
	}
	if _, err := s.Snapshot("claude", "skills", orig); err != nil {
		t.Fatalf("Snapshot: %v", err)
	}
	if err := s.Finalize(); err != nil {
		t.Fatalf("Finalize: %v", err)
	}
	return s
}

// TestDedupHitAndMiss verifies that identical content snapshotted across two
// sessions is stored once in the CAS (a hit, no second copy of the bytes),
// while distinct content is stored fresh (a miss).
func TestDedupHitAndMiss(t *testing.T) {
	tests := []struct {
		name        string
		first       string
		second      string
		wantHit     bool
		wantCASObjs int
	}{
		{name: "identical content dedups", first: "same-bytes", second: "same-bytes", wantHit: true, wantCASObjs: 1},
		{name: "distinct content stored twice", first: "alpha", second: "beta", wantHit: false, wantCASObjs: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			home := t.TempDir()

			origA := filepath.Join(t.TempDir(), "a.json")
			write(t, origA, tt.first)
			s1 := snapshotInto(t, home, origA)
			if h, m := s1.DedupStats(); h != 0 || m != 1 {
				t.Fatalf("first session dedup hits=%d misses=%d, want 0/1", h, m)
			}

			origB := filepath.Join(t.TempDir(), "b.json")
			write(t, origB, tt.second)
			s2 := snapshotInto(t, home, origB)
			h, m := s2.DedupStats()
			if tt.wantHit && (h != 1 || m != 0) {
				t.Fatalf("second session dedup hits=%d misses=%d, want 1/0 (hit)", h, m)
			}
			if !tt.wantHit && (h != 0 || m != 1) {
				t.Fatalf("second session dedup hits=%d misses=%d, want 0/1 (miss)", h, m)
			}

			// Count distinct CAS objects on disk.
			casDir := filepath.Join(filepath.Dir(s2.Root()), casDirName)
			objs, rerr := os.ReadDir(casDir)
			if rerr != nil {
				t.Fatalf("read cas dir: %v", rerr)
			}
			if len(objs) != tt.wantCASObjs {
				t.Fatalf("CAS objects = %d, want %d", len(objs), tt.wantCASObjs)
			}
		})
	}
}

// TestDedupRestoreByteIdentical asserts a deduplicated backup still restores
// the exact original bytes (the batch path is a hardlink to the CAS object).
func TestDedupRestoreByteIdentical(t *testing.T) {
	home := t.TempDir()

	origA := filepath.Join(t.TempDir(), "first.json")
	write(t, origA, "shared-content")
	snapshotInto(t, home, origA)

	// Second session backs up identical content (dedup hit), then we restore it.
	origB := filepath.Join(t.TempDir(), "second.json")
	write(t, origB, "shared-content")
	s2 := snapshotInto(t, home, origB)
	if h, _ := s2.DedupStats(); h != 1 {
		t.Fatalf("expected dedup hit on second session")
	}

	m, err := LoadManifest(s2.Root())
	if err != nil {
		t.Fatalf("LoadManifest: %v", err)
	}

	if err := os.Remove(origB); err != nil {
		t.Fatalf("remove orig: %v", err)
	}
	results := ApplyRestore(PlanRestore(m, ""))
	if len(results) != 1 || results[0].Class != RestoreRestore {
		t.Fatalf("restore result = %+v", results)
	}
	got, err := os.ReadFile(origB)
	if err != nil {
		t.Fatalf("read restored: %v", err)
	}
	if string(got) != "shared-content" {
		t.Fatalf("restored = %q, want shared-content", string(got))
	}
}

// TestDedupGCReapsOrphanedCAS verifies CAS objects whose referencing batches
// are pruned by GC get reaped, while content still referenced by a surviving
// batch is retained.
func TestDedupGCReapsOrphanedCAS(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	root, err := BackupsRoot()
	if err != nil {
		t.Fatalf("BackupsRoot: %v", err)
	}
	casDir := filepath.Join(root, casDirName)

	// Create one referenced CAS object via a real session, plus one orphan.
	orig := filepath.Join(t.TempDir(), "live.json")
	write(t, orig, "live-content")
	snapshotInto(t, home, orig)

	// Plant an orphan CAS object with no batch referencing it (link count 1).
	if err := os.MkdirAll(casDir, 0o755); err != nil {
		t.Fatalf("mkdir cas: %v", err)
	}
	orphan := filepath.Join(casDir, "deadbeef")
	if err := os.WriteFile(orphan, []byte("orphan"), 0o644); err != nil {
		t.Fatalf("write orphan: %v", err)
	}

	gcOrphanedCAS(casDir)

	if exists(orphan) {
		t.Fatalf("orphaned CAS object was not reaped")
	}
	// The live object (hardlinked from the surviving batch) must remain.
	objs, _ := os.ReadDir(casDir)
	if len(objs) != 1 {
		t.Fatalf("CAS objects after reap = %d, want 1 (the live one)", len(objs))
	}
}

// TestPinSurvivesGC asserts a pinned batch is never pruned while unpinned
// batches beyond keep-5 are removed.
func TestPinSurvivesGC(t *testing.T) {
	parent := t.TempDir()

	// Create 8 batches with manifest-bearing dirs: ts00..ts07 (sorted oldest→newest).
	var dirs []string
	for i := 0; i < 8; i++ {
		name := filepath.Join(parent, "ts0"+string(rune('0'+i)))
		write(t, filepath.Join(name, "manifest.json"), `{"entries":[]}`)
		dirs = append(dirs, name)
	}

	// Pin the OLDEST batch (ts00) — it must survive GC despite being oldest.
	pinnedDir := dirs[0]
	if err := os.WriteFile(filepath.Join(pinnedDir, pinMarker), []byte("x"), 0o644); err != nil {
		t.Fatalf("pin: %v", err)
	}
	if !IsPinned(pinnedDir) {
		t.Fatalf("IsPinned(pinned) = false")
	}

	if err := gcOldBackups(parent, DefaultRetentionCount); err != nil {
		t.Fatalf("gc: %v", err)
	}

	// Pinned oldest must still exist.
	if !exists(pinnedDir) {
		t.Fatalf("pinned batch was pruned by GC")
	}

	// There were 7 unpinned (ts01..ts07); keep-5 means the 2 oldest unpinned
	// (ts01, ts02) are pruned.
	if exists(dirs[1]) || exists(dirs[2]) {
		t.Fatalf("oldest unpinned batches were not pruned")
	}
	// The 5 newest unpinned survive.
	for _, d := range dirs[3:] {
		if !exists(d) {
			t.Fatalf("recent unpinned batch %s was wrongly pruned", d)
		}
	}
}

// TestSetPinnedAndList verifies SetPinned toggles the marker and ListBackups
// reflects the pin state.
func TestSetPinnedAndList(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	root, err := BackupsRoot()
	if err != nil {
		t.Fatalf("BackupsRoot: %v", err)
	}
	stamp := "20250101T000000Z"
	write(t, filepath.Join(root, stamp, "manifest.json"), `{"entries":[]}`)

	if err := SetPinned(stamp, true); err != nil {
		t.Fatalf("SetPinned(true): %v", err)
	}
	summaries, err := ListBackups()
	if err != nil {
		t.Fatalf("ListBackups: %v", err)
	}
	if len(summaries) != 1 || !summaries[0].Pinned {
		t.Fatalf("ListBackups pinned state = %+v", summaries)
	}

	if err := SetPinned(stamp, false); err != nil {
		t.Fatalf("SetPinned(false): %v", err)
	}
	summaries, _ = ListBackups()
	if summaries[0].Pinned {
		t.Fatalf("backup still pinned after unpin")
	}

	if err := SetPinned("does-not-exist", true); err == nil {
		t.Fatalf("SetPinned(missing) = nil, want error")
	}
}
