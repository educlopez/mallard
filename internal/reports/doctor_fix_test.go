package reports

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/educlopez/duck-ai/internal/agents"
	"github.com/educlopez/duck-ai/internal/skills"
)

func mkdir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
}

func TestFixDir(t *testing.T) {
	t.Run("relinks broken managed link to moved source", func(t *testing.T) {
		repo := t.TempDir()
		// current source lives at repo/skills/foo
		newSrc := filepath.Join(repo, "skills", "foo")
		mkdir(t, newSrc)
		// stale source path (where the link used to point) under repo, now gone
		oldSrc := filepath.Join(repo, "skills", "foo-old")

		dstDir := t.TempDir()
		link := filepath.Join(dstDir, "foo")
		if err := os.Symlink(oldSrc, link); err != nil {
			t.Fatal(err)
		}

		fixes := fixDir("claude", "skills",
			dstDir, []skills.Skill{{Name: "foo", SrcPath: newSrc}}, repo)

		if len(fixes) != 1 || fixes[0].Action != "relinked-broken" {
			t.Fatalf("fixes = %+v, want one relinked-broken", fixes)
		}
		target, err := os.Readlink(link)
		if err != nil {
			t.Fatalf("readlink: %v", err)
		}
		if target != newSrc {
			t.Fatalf("link target = %q, want %q", target, newSrc)
		}
	})

	t.Run("creates missing link for current source", func(t *testing.T) {
		repo := t.TempDir()
		src := filepath.Join(repo, "skills", "bar")
		mkdir(t, src)
		dstDir := t.TempDir()

		fixes := fixDir("claude", "skills",
			dstDir, []skills.Skill{{Name: "bar", SrcPath: src}}, repo)

		if len(fixes) != 1 || fixes[0].Action != "created-missing" {
			t.Fatalf("fixes = %+v, want one created-missing", fixes)
		}
		target, err := os.Readlink(filepath.Join(dstDir, "bar"))
		if err != nil {
			t.Fatalf("readlink: %v", err)
		}
		if target != src {
			t.Fatalf("link target = %q, want %q", target, src)
		}
	})

	t.Run("leaves healthy managed link untouched", func(t *testing.T) {
		repo := t.TempDir()
		src := filepath.Join(repo, "skills", "ok")
		mkdir(t, src)
		dstDir := t.TempDir()
		if err := os.Symlink(src, filepath.Join(dstDir, "ok")); err != nil {
			t.Fatal(err)
		}
		fixes := fixDir("claude", "skills",
			dstDir, []skills.Skill{{Name: "ok", SrcPath: src}}, repo)
		if len(fixes) != 0 {
			t.Fatalf("healthy link should not be fixed: %+v", fixes)
		}
	})

	t.Run("never touches unmanaged broken link pointing outside repo", func(t *testing.T) {
		repo := t.TempDir()
		dstDir := t.TempDir()
		outside := filepath.Join(t.TempDir(), "elsewhere") // does not exist, outside repo
		link := filepath.Join(dstDir, "foo")
		if err := os.Symlink(outside, link); err != nil {
			t.Fatal(err)
		}
		// A source item named "foo" also exists, but the existing link is
		// unmanaged (points outside repo) so fix must NOT relink it; and since
		// an entry named "foo" is already present, it must NOT create either.
		src := filepath.Join(repo, "skills", "foo")
		mkdir(t, src)

		fixes := fixDir("claude", "skills",
			dstDir, []skills.Skill{{Name: "foo", SrcPath: src}}, repo)

		if len(fixes) != 0 {
			t.Fatalf("unmanaged link must be left alone: %+v", fixes)
		}
		// link still points at the original outside target
		target, _ := os.Readlink(link)
		if target != outside {
			t.Fatalf("unmanaged link was modified: target = %q", target)
		}
	})

	t.Run("never touches real file in dst", func(t *testing.T) {
		repo := t.TempDir()
		dstDir := t.TempDir()
		real := filepath.Join(dstDir, "foo")
		if err := os.WriteFile(real, []byte("real"), 0o644); err != nil {
			t.Fatal(err)
		}
		src := filepath.Join(repo, "skills", "foo")
		mkdir(t, src)

		fixes := fixDir("claude", "skills",
			dstDir, []skills.Skill{{Name: "foo", SrcPath: src}}, repo)
		if len(fixes) != 0 {
			t.Fatalf("real file must be left alone: %+v", fixes)
		}
		info, err := os.Lstat(real)
		if err != nil {
			t.Fatal(err)
		}
		if info.Mode()&os.ModeSymlink != 0 {
			t.Fatalf("real file was replaced by a symlink")
		}
	})

	t.Run("does not remove broken managed link when source is gone", func(t *testing.T) {
		repo := t.TempDir()
		dstDir := t.TempDir()
		gone := filepath.Join(repo, "skills", "gone") // under repo but never created
		link := filepath.Join(dstDir, "gone")
		if err := os.Symlink(gone, link); err != nil {
			t.Fatal(err)
		}
		// No source item named "gone" supplied -> genuine drift, leave it.
		fixes := fixDir("claude", "skills", dstDir, nil, repo)
		if len(fixes) != 0 {
			t.Fatalf("broken managed link with no source must be left: %+v", fixes)
		}
		if _, err := os.Lstat(link); err != nil {
			t.Fatalf("link was removed: %v", err)
		}
	})

	t.Run("empty dstDir is a no-op", func(t *testing.T) {
		repo := t.TempDir()
		src := filepath.Join(repo, "skills", "x")
		mkdir(t, src)
		fixes := fixDir("codex", "agents", "", []skills.Skill{{Name: "x", SrcPath: src}}, repo)
		if fixes != nil {
			t.Fatalf("empty dstDir should produce no fixes: %+v", fixes)
		}
	})
}

func TestDoctorFixSmoke(t *testing.T) {
	// Smoke test: DoctorFix against a temp repo root captures output into a
	// buffer and must not error. It walks whichever agents are detected on the
	// host; any non-nil error would indicate a discovery bug.
	repo := t.TempDir()
	mkdir(t, filepath.Join(repo, "skills"))
	var buf bytes.Buffer
	if _, err := DoctorFix(&buf, repo, agents.ScopeGlobal, ""); err != nil {
		t.Fatalf("DoctorFix() error = %v", err)
	}
	if buf.Len() == 0 {
		t.Fatalf("DoctorFix() wrote no output")
	}
}
