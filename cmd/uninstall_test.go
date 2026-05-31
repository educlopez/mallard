package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/educlopez/duck-ai/internal/agents"
)

// fakeAdapter is a test double for agents.Adapter pointing its three dirs
// wherever the test wants.
type fakeAdapter struct {
	id          string
	detect      bool
	skillsDir   string
	commandsDir string
	agentsDir   string
}

func (f fakeAdapter) ID() string          { return f.id }
func (f fakeAdapter) DisplayName() string { return f.id }
func (f fakeAdapter) Detect() bool        { return f.detect }
func (f fakeAdapter) SkillsDir() string   { return f.skillsDir }
func (f fakeAdapter) CommandsDir() string { return f.commandsDir }
func (f fakeAdapter) AgentsDir() string   { return f.agentsDir }

func mkdir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
}

// existsLstat reports whether anything (incl. broken symlink) exists at path.
func existsLstat(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}

func TestUnlinkManagedInDir(t *testing.T) {
	t.Run("removes managed symlink pointing into repo", func(t *testing.T) {
		repo := t.TempDir()
		src := filepath.Join(repo, "skills", "foo")
		mkdir(t, src)
		dst := t.TempDir()
		link := filepath.Join(dst, "foo")
		if err := os.Symlink(src, link); err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		removed, skipped := unlinkManagedInDir(&buf, dst, repo, false)
		if removed != 1 || skipped != 0 {
			t.Fatalf("removed=%d skipped=%d, want 1/0", removed, skipped)
		}
		if existsLstat(link) {
			t.Fatalf("managed link was not removed")
		}
	})

	t.Run("leaves symlink pointing outside repo untouched", func(t *testing.T) {
		repo := t.TempDir()
		outside := filepath.Join(t.TempDir(), "elsewhere")
		mkdir(t, outside)
		dst := t.TempDir()
		link := filepath.Join(dst, "foo")
		if err := os.Symlink(outside, link); err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		removed, skipped := unlinkManagedInDir(&buf, dst, repo, false)
		if removed != 0 || skipped != 1 {
			t.Fatalf("removed=%d skipped=%d, want 0/1", removed, skipped)
		}
		if !existsLstat(link) {
			t.Fatalf("unmanaged link was removed")
		}
		target, _ := os.Readlink(link)
		if target != outside {
			t.Fatalf("unmanaged link target changed: %q", target)
		}
	})

	t.Run("leaves real file untouched", func(t *testing.T) {
		repo := t.TempDir()
		dst := t.TempDir()
		real := filepath.Join(dst, "foo")
		if err := os.WriteFile(real, []byte("real"), 0o644); err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		removed, skipped := unlinkManagedInDir(&buf, dst, repo, false)
		if removed != 0 || skipped != 1 {
			t.Fatalf("removed=%d skipped=%d, want 0/1", removed, skipped)
		}
		if !existsLstat(real) {
			t.Fatalf("real file was removed")
		}
	})

	t.Run("dry-run changes nothing", func(t *testing.T) {
		repo := t.TempDir()
		src := filepath.Join(repo, "skills", "foo")
		mkdir(t, src)
		dst := t.TempDir()
		link := filepath.Join(dst, "foo")
		if err := os.Symlink(src, link); err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		removed, skipped := unlinkManagedInDir(&buf, dst, repo, true)
		if removed != 1 || skipped != 0 {
			t.Fatalf("removed=%d skipped=%d, want 1/0", removed, skipped)
		}
		if !existsLstat(link) {
			t.Fatalf("dry-run removed the link")
		}
	})

	t.Run("empty dir is a no-op", func(t *testing.T) {
		repo := t.TempDir()
		var buf bytes.Buffer
		removed, skipped := unlinkManagedInDir(&buf, "", repo, false)
		if removed != 0 || skipped != 0 {
			t.Fatalf("removed=%d skipped=%d, want 0/0", removed, skipped)
		}
	})

	t.Run("skips hidden entries", func(t *testing.T) {
		repo := t.TempDir()
		src := filepath.Join(repo, "skills", "foo")
		mkdir(t, src)
		dst := t.TempDir()
		hidden := filepath.Join(dst, ".keep")
		if err := os.Symlink(src, hidden); err != nil {
			t.Fatal(err)
		}
		var buf bytes.Buffer
		removed, skipped := unlinkManagedInDir(&buf, dst, repo, false)
		if removed != 0 || skipped != 0 {
			t.Fatalf("removed=%d skipped=%d, want 0/0 (hidden skipped)", removed, skipped)
		}
		if !existsLstat(hidden) {
			t.Fatalf("hidden entry was removed")
		}
	})
}

func TestUninstallAcrossAgent(t *testing.T) {
	repo := t.TempDir()
	managedSrc := filepath.Join(repo, "skills", "foo")
	mkdir(t, managedSrc)
	outside := filepath.Join(t.TempDir(), "ext")
	mkdir(t, outside)

	skillsDir := t.TempDir()
	commandsDir := t.TempDir()

	managedLink := filepath.Join(skillsDir, "foo")    // managed -> repo
	unmanagedLink := filepath.Join(skillsDir, "bar")  // -> outside repo
	realFile := filepath.Join(commandsDir, "real.md") // real file
	if err := os.Symlink(managedSrc, managedLink); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, unmanagedLink); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(realFile, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}

	a := fakeAdapter{
		id:          "claude",
		detect:      true,
		skillsDir:   skillsDir,
		commandsDir: commandsDir,
	}

	var buf bytes.Buffer
	if err := uninstall(&buf, repo, []agents.Adapter{a}, false); err != nil {
		t.Fatalf("uninstall: %v", err)
	}

	if existsLstat(managedLink) {
		t.Fatalf("managed link should have been removed")
	}
	if !existsLstat(unmanagedLink) {
		t.Fatalf("unmanaged link should have been left")
	}
	if !existsLstat(realFile) {
		t.Fatalf("real file should have been left")
	}
}

func TestUninstallNoAgents(t *testing.T) {
	var buf bytes.Buffer
	if err := uninstall(&buf, t.TempDir(), nil, false); err != nil {
		t.Fatalf("uninstall: %v", err)
	}
	if !bytes.Contains(buf.Bytes(), []byte("No agents detected")) {
		t.Fatalf("expected no-agents message, got: %s", buf.String())
	}
}

func TestParseUninstallArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    UninstallArgs
		wantErr bool
	}{
		{name: "default", args: nil, want: UninstallArgs{}},
		{name: "all", args: []string{"--all"}, want: UninstallArgs{All: true}},
		{name: "agent", args: []string{"--agent", "claude"}, want: UninstallArgs{AgentID: "claude"}},
		{name: "dry-run", args: []string{"--dry-run"}, want: UninstallArgs{DryRun: true}},
		{name: "agent missing value", args: []string{"--agent"}, wantErr: true},
		{name: "unknown flag", args: []string{"--nope"}, wantErr: true},
		{name: "agent and all conflict", args: []string{"--agent", "claude", "--all"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUninstallArgs(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}
