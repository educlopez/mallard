package skills

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const claudeAgent = `---
name: panda-expert
description: Use for "Panda theme" questions: modules, layouts.
tools: Read, Grep, Glob, Bash
model: sonnet
color: orange
metadata:
  author: Eduardo Calvo
---

# Panda Expert

Say "hi" and use C:\path.
`

func renderRepo(t *testing.T) (string, []Skill) {
	t.Helper()
	root := t.TempDir()
	src := filepath.Join(root, "claude", "agents", "panda-expert.md")
	writeFile(t, src, claudeAgent)
	return root, []Skill{{Name: "panda-expert.md", SrcPath: src}}
}

func TestAgentsForPassthrough(t *testing.T) {
	root, src := renderRepo(t)
	for _, id := range []string{"claude", "gemini", "windsurf", "agents"} {
		got, err := AgentsFor(root, id, src, true)
		if err != nil {
			t.Fatal(err)
		}
		if got[0].SrcPath != src[0].SrcPath {
			t.Fatalf("%s: SrcPath = %q, want source as-is", id, got[0].SrcPath)
		}
	}
	if _, err := os.Stat(filepath.Join(root, GeneratedDir)); !os.IsNotExist(err) {
		t.Fatalf("passthrough must not create %s", GeneratedDir)
	}
}

func TestAgentsForCodex(t *testing.T) {
	root, src := renderRepo(t)
	got, err := AgentsFor(root, "codex", src, true)
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(root, GeneratedDir, "codex", "agents", "panda-expert.toml")
	if got[0].Name != "panda-expert.toml" || got[0].SrcPath != want {
		t.Fatalf("got %+v, want %s", got[0], want)
	}
	data, _ := os.ReadFile(want)
	out := string(data)
	for _, s := range []string{
		`name = "panda_expert"`,
		`description = "Use for \"Panda theme\" questions: modules, layouts."`,
		"developer_instructions = \"\"\"\n# Panda Expert\n",
		`Say \"hi\" and use C:\\path.`,
	} {
		if !strings.Contains(out, s) {
			t.Fatalf("codex output missing %q:\n%s", s, out)
		}
	}
	if strings.Contains(out, "sonnet") || strings.Contains(out, "orange") {
		t.Fatalf("codex output must drop Claude model/color:\n%s", out)
	}
}

func TestAgentsForOpenCode(t *testing.T) {
	root, src := renderRepo(t)
	got, err := AgentsFor(root, "opencode", src, true)
	if err != nil {
		t.Fatal(err)
	}
	data, _ := os.ReadFile(got[0].SrcPath)
	out := string(data)
	wantHead := "---\ndescription: \"Use for \\\"Panda theme\\\" questions: modules, layouts.\"\nmode: subagent\npermission:\n  edit: deny\n---\n# Panda Expert\n"
	if !strings.HasPrefix(out, wantHead) {
		t.Fatalf("opencode output:\n%s\nwant prefix:\n%s", out, wantHead)
	}
	if strings.Contains(out, "bash: deny") {
		t.Fatalf("agent has Bash, must not deny it:\n%s", out)
	}
}

func TestAgentsForCursor(t *testing.T) {
	root, src := renderRepo(t)
	got, err := AgentsFor(root, "cursor", src, true)
	if err != nil {
		t.Fatal(err)
	}
	data, _ := os.ReadFile(got[0].SrcPath)
	out := string(data)
	for _, s := range []string{`name: "panda-expert"`, "model: inherit\n", "readonly: true\n", "# Panda Expert"} {
		if !strings.Contains(out, s) {
			t.Fatalf("cursor output missing %q:\n%s", s, out)
		}
	}
}

func TestAgentsForDryRunWritesNothing(t *testing.T) {
	root, src := renderRepo(t)
	got, err := AgentsFor(root, "codex", src, false)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(got[0].SrcPath); !os.IsNotExist(err) {
		t.Fatalf("dry run wrote %s", got[0].SrcPath)
	}
}

func TestAgentsForRemovesStaleRenderings(t *testing.T) {
	root, src := renderRepo(t)
	stale := filepath.Join(root, GeneratedDir, "codex", "agents", "gone.toml")
	writeFile(t, stale, "x")
	if _, err := AgentsFor(root, "codex", src, true); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatalf("stale rendering %s survived", stale)
	}
}
