package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/educlopez/mallard/internal/agents"
)

// buildSourceRepo creates a minimal mallard source repo containing a single
// skill, and returns the repo root.
func buildSourceRepo(t *testing.T) string {
	t.Helper()
	repo := t.TempDir()
	skillDir := filepath.Join(repo, "skills", "demo")
	mkdir(t, skillDir)
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("---\nname: demo\n---\n"), 0o644); err != nil {
		t.Fatalf("write SKILL.md: %v", err)
	}
	return repo
}

// chdir switches into dir for the duration of the test and restores afterward.
func chdir(t *testing.T, dir string) {
	t.Helper()
	prev, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir %s: %v", dir, err)
	}
	t.Cleanup(func() { _ = os.Chdir(prev) })
}

// TestInstallWorkspaceScope verifies workspace-scope install links into the
// project-local <cwd>/.claude/skills dir and NOT into the agent's global dir.
func TestInstallWorkspaceScope(t *testing.T) {
	repo := buildSourceRepo(t)
	ws := t.TempDir()
	chdir(t, ws)

	// Global dir is a distinct temp dir; workspace scope must not write here.
	globalSkills := t.TempDir()
	a := fakeAdapter{
		id:        "claude",
		detect:    true,
		skillsDir: globalSkills,
	}

	if err := installToAgents(repo, []agents.Adapter{a}, agents.ScopeWorkspace); err != nil {
		t.Fatalf("installToAgents: %v", err)
	}

	wantLink := filepath.Join(ws, ".claude", "skills", "demo")
	if !existsLstat(wantLink) {
		t.Fatalf("workspace link %s was not created", wantLink)
	}
	target, err := os.Readlink(wantLink)
	if err != nil {
		t.Fatalf("readlink: %v", err)
	}
	if want := filepath.Join(repo, "skills", "demo"); target != want {
		t.Fatalf("link target = %q, want %q", target, want)
	}

	// The global dir must remain empty in workspace scope.
	if existsLstat(filepath.Join(globalSkills, "demo")) {
		t.Fatalf("workspace install leaked into global skills dir")
	}
}

// TestInstallGlobalScopeUnchanged verifies global scope still links into the
// adapter's configured (home-based) dir and not into any workspace dir.
func TestInstallGlobalScopeUnchanged(t *testing.T) {
	repo := buildSourceRepo(t)
	ws := t.TempDir()
	chdir(t, ws)

	globalSkills := t.TempDir()
	a := fakeAdapter{id: "claude", detect: true, skillsDir: globalSkills}

	if err := installToAgents(repo, []agents.Adapter{a}, agents.ScopeGlobal); err != nil {
		t.Fatalf("installToAgents: %v", err)
	}

	if !existsLstat(filepath.Join(globalSkills, "demo")) {
		t.Fatalf("global install did not create link in global dir")
	}
	if existsLstat(filepath.Join(ws, ".claude", "skills", "demo")) {
		t.Fatalf("global install leaked into workspace dir")
	}
}

// TestUninstallWorkspaceScope round-trips install+uninstall in workspace scope.
func TestUninstallWorkspaceScope(t *testing.T) {
	repo := buildSourceRepo(t)
	ws := t.TempDir()
	chdir(t, ws)

	a := fakeAdapter{id: "claude", detect: true, skillsDir: t.TempDir()}
	if err := installToAgents(repo, []agents.Adapter{a}, agents.ScopeWorkspace); err != nil {
		t.Fatalf("install: %v", err)
	}
	link := filepath.Join(ws, ".claude", "skills", "demo")
	if !existsLstat(link) {
		t.Fatalf("precondition: link not created")
	}

	var buf bytes.Buffer
	if err := uninstall(&buf, repo, []agents.Adapter{a}, false, agents.ScopeWorkspace, ws); err != nil {
		t.Fatalf("uninstall: %v", err)
	}
	if existsLstat(link) {
		t.Fatalf("workspace uninstall did not remove managed link")
	}
}
