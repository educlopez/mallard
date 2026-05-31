package skills

import (
	"os"
	"path/filepath"
	"testing"
)

// writeFile creates a file (and parent dirs) with the given content.
func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

// buildRepo builds a fake mallard repo tree in a temp dir and returns its root.
func buildRepo(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	// skills/foo/SKILL.md  -> valid skill
	writeFile(t, filepath.Join(root, "skills", "foo", "SKILL.md"), "---\nname: foo\n---\n")
	// skills/bar/SKILL.md  -> valid skill
	writeFile(t, filepath.Join(root, "skills", "bar", "SKILL.md"), "---\nname: bar\n---\n")
	// skills/nope/ -> dir without SKILL.md, must be ignored
	if err := os.MkdirAll(filepath.Join(root, "skills", "nope"), 0o755); err != nil {
		t.Fatal(err)
	}
	// commands
	writeFile(t, filepath.Join(root, "claude", "commands", "bar.md"), "# bar\n")
	writeFile(t, filepath.Join(root, "claude", "commands", "qux.md"), "# qux\n")
	writeFile(t, filepath.Join(root, "claude", "commands", "ignore.txt"), "not a command\n")
	// agents
	writeFile(t, filepath.Join(root, "claude", "agents", "baz.md"), "# baz\n")
	return root
}

func names(skills []Skill) []string {
	out := make([]string, len(skills))
	for i, s := range skills {
		out[i] = s.Name
	}
	return out
}

func hasName(skills []Skill, name string) (Skill, bool) {
	for _, s := range skills {
		if s.Name == name {
			return s, true
		}
	}
	return Skill{}, false
}

func TestDiscoverSkills(t *testing.T) {
	root := buildRepo(t)

	got, err := DiscoverSkills(root)
	if err != nil {
		t.Fatalf("DiscoverSkills() error = %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("DiscoverSkills() returned %d skills, want 2 (%v)", len(got), names(got))
	}
	foo, ok := hasName(got, "foo")
	if !ok {
		t.Fatalf("expected skill 'foo' in %v", names(got))
	}
	wantSrc := filepath.Join(root, "skills", "foo")
	if foo.SrcPath != wantSrc {
		t.Fatalf("foo.SrcPath = %q, want %q", foo.SrcPath, wantSrc)
	}
	if _, ok := hasName(got, "nope"); ok {
		t.Fatalf("dir without SKILL.md must not be discovered: %v", names(got))
	}
}

func TestDiscoverSkillsMissingDir(t *testing.T) {
	// skills/ does not exist -> error (DiscoverSkills does not swallow ENOENT).
	root := t.TempDir()
	_, err := DiscoverSkills(root)
	if err == nil {
		t.Fatalf("DiscoverSkills() on missing dir: expected error, got nil")
	}
}

func TestDiscoverCommands(t *testing.T) {
	tests := []struct {
		name  string
		build func(t *testing.T) string
		want  []string
	}{
		{
			name:  "populated",
			build: buildRepo,
			want:  []string{"bar.md", "qux.md"},
		},
		{
			name: "missing dir returns nil not error",
			build: func(t *testing.T) string {
				return t.TempDir()
			},
			want: nil,
		},
		{
			name: "empty dir returns nil",
			build: func(t *testing.T) string {
				root := t.TempDir()
				if err := os.MkdirAll(filepath.Join(root, "claude", "commands"), 0o755); err != nil {
					t.Fatal(err)
				}
				return root
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.build(t)
			got, err := DiscoverCommands(root)
			if err != nil {
				t.Fatalf("DiscoverCommands() error = %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("DiscoverCommands() = %v, want %v", names(got), tt.want)
			}
			for _, w := range tt.want {
				if _, ok := hasName(got, w); !ok {
					t.Fatalf("missing command %q in %v", w, names(got))
				}
			}
		})
	}
}

func TestDiscoverCommandsSrcPath(t *testing.T) {
	root := buildRepo(t)
	got, err := DiscoverCommands(root)
	if err != nil {
		t.Fatal(err)
	}
	bar, ok := hasName(got, "bar.md")
	if !ok {
		t.Fatalf("expected bar.md in %v", names(got))
	}
	want := filepath.Join(root, "claude", "commands", "bar.md")
	if bar.SrcPath != want {
		t.Fatalf("bar.md SrcPath = %q, want %q", bar.SrcPath, want)
	}
}

func TestDiscoverAgents(t *testing.T) {
	tests := []struct {
		name  string
		build func(t *testing.T) string
		want  []string
	}{
		{
			name:  "populated",
			build: buildRepo,
			want:  []string{"baz.md"},
		},
		{
			name: "missing dir returns nil not error",
			build: func(t *testing.T) string {
				return t.TempDir()
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.build(t)
			got, err := DiscoverAgents(root)
			if err != nil {
				t.Fatalf("DiscoverAgents() error = %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("DiscoverAgents() = %v, want %v", names(got), tt.want)
			}
			for _, w := range tt.want {
				if _, ok := hasName(got, w); !ok {
					t.Fatalf("missing agent %q in %v", w, names(got))
				}
			}
		})
	}
}

func TestLink(t *testing.T) {
	t.Run("linked when missing", func(t *testing.T) {
		src := t.TempDir()
		dstDir := t.TempDir()
		skill := Skill{Name: "foo", SrcPath: src}

		res := Link("claude", skill, dstDir)
		if res.Status != "linked" {
			t.Fatalf("Status = %q (err=%v), want linked", res.Status, res.Err)
		}
		target, err := os.Readlink(filepath.Join(dstDir, "foo"))
		if err != nil {
			t.Fatalf("readlink: %v", err)
		}
		if target != src {
			t.Fatalf("symlink target = %q, want %q", target, src)
		}
	})

	t.Run("already_linked when symlink points at src", func(t *testing.T) {
		src := t.TempDir()
		dstDir := t.TempDir()
		skill := Skill{Name: "foo", SrcPath: src}

		if res := Link("claude", skill, dstDir); res.Status != "linked" {
			t.Fatalf("first Link Status = %q", res.Status)
		}
		res := Link("claude", skill, dstDir)
		if res.Status != "already_linked" {
			t.Fatalf("second Link Status = %q (err=%v), want already_linked", res.Status, res.Err)
		}
	})

	t.Run("relink when symlink points elsewhere", func(t *testing.T) {
		oldSrc := t.TempDir()
		newSrc := t.TempDir()
		dstDir := t.TempDir()

		if err := os.Symlink(oldSrc, filepath.Join(dstDir, "foo")); err != nil {
			t.Fatal(err)
		}
		res := Link("claude", Skill{Name: "foo", SrcPath: newSrc}, dstDir)
		if res.Status != "linked" {
			t.Fatalf("Status = %q (err=%v), want linked", res.Status, res.Err)
		}
		target, _ := os.Readlink(filepath.Join(dstDir, "foo"))
		if target != newSrc {
			t.Fatalf("symlink target = %q, want %q", target, newSrc)
		}
	})

	t.Run("skipped when real file exists", func(t *testing.T) {
		src := t.TempDir()
		dstDir := t.TempDir()
		writeFile(t, filepath.Join(dstDir, "foo"), "real file\n")

		res := Link("claude", Skill{Name: "foo", SrcPath: src}, dstDir)
		if res.Status != "skipped" {
			t.Fatalf("Status = %q, want skipped", res.Status)
		}
		if res.Err == nil {
			t.Fatalf("expected an explanatory error on skip")
		}
	})

	t.Run("creates parent dir when dstDir missing", func(t *testing.T) {
		src := t.TempDir()
		dstDir := filepath.Join(t.TempDir(), "nested", "dir")
		res := Link("claude", Skill{Name: "foo", SrcPath: src}, dstDir)
		if res.Status != "linked" {
			t.Fatalf("Status = %q (err=%v), want linked", res.Status, res.Err)
		}
	})
}

func TestCheckLink(t *testing.T) {
	t.Run("missing", func(t *testing.T) {
		dstDir := t.TempDir()
		if got := CheckLink(dstDir, "foo"); got != "missing" {
			t.Fatalf("CheckLink = %q, want missing", got)
		}
	})

	t.Run("ok", func(t *testing.T) {
		src := t.TempDir()
		dstDir := t.TempDir()
		if err := os.Symlink(src, filepath.Join(dstDir, "foo")); err != nil {
			t.Fatal(err)
		}
		if got := CheckLink(dstDir, "foo"); got != "ok" {
			t.Fatalf("CheckLink = %q, want ok", got)
		}
	})

	t.Run("broken when target removed", func(t *testing.T) {
		gone := filepath.Join(t.TempDir(), "gone")
		dstDir := t.TempDir()
		if err := os.Symlink(gone, filepath.Join(dstDir, "foo")); err != nil {
			t.Fatal(err)
		}
		if got := CheckLink(dstDir, "foo"); got != "broken" {
			t.Fatalf("CheckLink = %q, want broken", got)
		}
	})

	t.Run("not_symlink when real file", func(t *testing.T) {
		dstDir := t.TempDir()
		writeFile(t, filepath.Join(dstDir, "foo"), "x\n")
		if got := CheckLink(dstDir, "foo"); got != "not_symlink" {
			t.Fatalf("CheckLink = %q, want not_symlink", got)
		}
	})
}
