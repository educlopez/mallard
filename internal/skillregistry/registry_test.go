package skillregistry

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/educlopez/mallard/internal/agents"
)

func TestParseFrontmatter(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want map[string]string
		// keysAbsent lists keys that must NOT be present in the result.
		keysAbsent []string
	}{
		{
			name: "simple key values",
			in:   "---\nname: foo\nversion: 1.2.3\n---\nbody",
			want: map[string]string{"name": "foo", "version": "1.2.3"},
		},
		{
			name: "no frontmatter returns empty",
			in:   "# just a heading\nno frontmatter here",
			want: map[string]string{},
		},
		{
			name: "unterminated frontmatter returns empty",
			in:   "---\nname: foo\nnever closes",
			want: map[string]string{},
		},
		{
			name: "quoted version is unquoted",
			in:   "---\nversion: \"2.0.0\"\nname: 'bar'\n---\n",
			want: map[string]string{"version": "2.0.0", "name": "bar"},
		},
		{
			name: "folded description joins on spaces",
			in:   "---\nname: foo\ndescription: >\n  Use this skill when\n  the user asks for\n  something cool\n---\n",
			want: map[string]string{
				"name":        "foo",
				"description": "Use this skill when the user asks for something cool",
			},
		},
		{
			name: "literal block description joins on spaces",
			in:   "---\nname: foo\ndescription: |\n  line one\n  line two\n---\n",
			want: map[string]string{
				"name":        "foo",
				"description": "line one line two",
			},
		},
		{
			// CRITICAL: a metadata: block with nested indented keys must not
			// corrupt parsing, and the nested keys must be ignored.
			name: "nested metadata block ignores indented keys",
			in: "---\n" +
				"name: foo\n" +
				"version: 1.0.0\n" +
				"metadata:\n" +
				"  author: Eduardo Calvo\n" +
				"  license: MIT\n" +
				"description: top level desc\n" +
				"---\n",
			want: map[string]string{
				"name":        "foo",
				"version":     "1.0.0",
				"description": "top level desc",
			},
			keysAbsent: []string{"author", "license"},
		},
		{
			name: "value with colon is preserved",
			in:   "---\nname: foo\ndescription: foo: bar baz\nversion: 1.0.0\n---\n",
			want: map[string]string{
				"name":        "foo",
				"description": "foo: bar baz",
				"version":     "1.0.0",
			},
		},
		{
			name: "quoted value with colon is unquoted",
			in:   "---\nname: foo\ndescription: \"hello: world\"\n---\n",
			want: map[string]string{
				"name":        "foo",
				"description": "hello: world",
			},
		},
		{
			// A folded description preceding a nested metadata block must not
			// swallow the nested lines past the next top-level key.
			name: "folded description then nested metadata",
			in: "---\n" +
				"name: foo\n" +
				"description: >\n" +
				"  first line\n" +
				"  second line\n" +
				"metadata:\n" +
				"  author: Eduardo Calvo\n" +
				"---\n",
			want: map[string]string{
				"name":        "foo",
				"description": "first line second line",
			},
			keysAbsent: []string{"author"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseFrontmatter([]byte(tt.in))
			for k, v := range tt.want {
				if got[k] != v {
					t.Fatalf("key %q = %q, want %q (full=%v)", k, got[k], v, got)
				}
			}
			for _, k := range tt.keysAbsent {
				if _, ok := got[k]; ok {
					t.Fatalf("key %q must be absent, got value %q", k, got[k])
				}
			}
		})
	}
}

func TestUnquote(t *testing.T) {
	cases := map[string]string{
		`"hello"`: "hello",
		`'hello'`: "hello",
		`hello`:   "hello",
		`"x`:      `"x`,
		``:        ``,
	}
	for in, want := range cases {
		if got := unquote(in); got != want {
			t.Fatalf("unquote(%q) = %q, want %q", in, got, want)
		}
	}
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
}

func findManifest(ms []Manifest, name, kind string) (Manifest, bool) {
	for _, m := range ms {
		if m.Name == name && m.Kind == kind {
			return m, true
		}
	}
	return Manifest{}, false
}

func TestParseSource(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "skills", "foo", "SKILL.md"),
		"---\nname: foo\ndescription: a skill\nversion: 1.0.0\n---\n")
	writeFile(t, filepath.Join(root, "claude", "commands", "bar.md"),
		"---\ndescription: a command\n---\n")
	writeFile(t, filepath.Join(root, "claude", "agents", "baz.md"),
		"---\ndescription: an agent\n---\n")
	// noise that must be ignored
	if err := os.MkdirAll(filepath.Join(root, "skills", "nope"), 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(root, "claude", "commands", "note.txt"), "ignore me\n")

	ms, err := ParseSource(root)
	if err != nil {
		t.Fatalf("ParseSource() error = %v", err)
	}
	if len(ms) != 3 {
		t.Fatalf("ParseSource() returned %d manifests, want 3: %+v", len(ms), ms)
	}

	skill, ok := findManifest(ms, "foo", "skill")
	if !ok {
		t.Fatalf("missing skill 'foo' in %+v", ms)
	}
	if skill.Description != "a skill" || skill.Version != "1.0.0" {
		t.Fatalf("skill fields = %+v", skill)
	}
	if _, ok := findManifest(ms, "bar", "command"); !ok {
		t.Fatalf("missing command 'bar' in %+v", ms)
	}
	if _, ok := findManifest(ms, "baz", "agent"); !ok {
		t.Fatalf("missing agent 'baz' in %+v", ms)
	}
}

func TestParseSourceEmptyRepo(t *testing.T) {
	ms, err := ParseSource(t.TempDir())
	if err != nil {
		t.Fatalf("ParseSource() on empty repo error = %v", err)
	}
	if len(ms) != 0 {
		t.Fatalf("expected no manifests, got %+v", ms)
	}
}

// fakeAdapter implements agents.Adapter against temp dirs for ParseInstalled.
type fakeAdapter struct {
	skillsDir, commandsDir, agentsDir string
}

func (f fakeAdapter) ID() string          { return "fake" }
func (f fakeAdapter) DisplayName() string { return "Fake" }
func (f fakeAdapter) Detect() bool        { return true }
func (f fakeAdapter) SkillsDir() string   { return f.skillsDir }
func (f fakeAdapter) CommandsDir() string { return f.commandsDir }
func (f fakeAdapter) AgentsDir() string   { return f.agentsDir }

func (f fakeAdapter) SkillsDirFor(_ agents.Scope, _ string) string   { return f.skillsDir }
func (f fakeAdapter) CommandsDirFor(_ agents.Scope, _ string) string { return f.commandsDir }
func (f fakeAdapter) AgentsDirFor(_ agents.Scope, _ string) string   { return f.agentsDir }

func TestParseInstalled(t *testing.T) {
	// Build a source repo.
	src := t.TempDir()
	skillSrc := filepath.Join(src, "skills", "foo")
	writeFile(t, filepath.Join(skillSrc, "SKILL.md"), "---\nname: foo\ndescription: s\n---\n")
	cmdSrc := filepath.Join(src, "claude", "commands", "bar.md")
	writeFile(t, cmdSrc, "---\ndescription: c\n---\n")
	agentSrc := filepath.Join(src, "claude", "agents", "baz.md")
	writeFile(t, agentSrc, "---\ndescription: a\n---\n")

	// Build install dirs with symlinks pointing at the sources.
	skillsDir := t.TempDir()
	commandsDir := t.TempDir()
	agentsDir := t.TempDir()
	if err := os.Symlink(skillSrc, filepath.Join(skillsDir, "foo")); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(cmdSrc, filepath.Join(commandsDir, "bar.md")); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(agentSrc, filepath.Join(agentsDir, "baz.md")); err != nil {
		t.Fatal(err)
	}
	// A non-symlink entry must be skipped (drift surfaced elsewhere).
	writeFile(t, filepath.Join(skillsDir, "real-dir", "SKILL.md"), "---\nname: real\n---\n")

	adapter := fakeAdapter{skillsDir: skillsDir, commandsDir: commandsDir, agentsDir: agentsDir}
	ms, err := ParseInstalled(adapter)
	if err != nil {
		t.Fatalf("ParseInstalled() error = %v", err)
	}
	if len(ms) != 3 {
		t.Fatalf("ParseInstalled() = %d manifests, want 3 (real-dir must be skipped): %+v", len(ms), ms)
	}
	if _, ok := findManifest(ms, "foo", "skill"); !ok {
		t.Fatalf("missing installed skill foo: %+v", ms)
	}
	if _, ok := findManifest(ms, "bar", "command"); !ok {
		t.Fatalf("missing installed command bar: %+v", ms)
	}
	if _, ok := findManifest(ms, "baz", "agent"); !ok {
		t.Fatalf("missing installed agent baz: %+v", ms)
	}
}

func TestParseInstalledEmptyDirs(t *testing.T) {
	// All adapter dirs empty string -> no walking, no error.
	ms, err := ParseInstalled(fakeAdapter{})
	if err != nil {
		t.Fatalf("ParseInstalled() error = %v", err)
	}
	if len(ms) != 0 {
		t.Fatalf("expected no manifests, got %+v", ms)
	}
}
