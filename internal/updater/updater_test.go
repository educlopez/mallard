package updater

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/educlopez/mallard/internal/agents"
	"github.com/educlopez/mallard/internal/backup"
	"github.com/educlopez/mallard/internal/skills"
)

func TestClassify(t *testing.T) {
	t.Run("missing", func(t *testing.T) {
		src := t.TempDir()
		dst := filepath.Join(t.TempDir(), "foo")
		if got := classify(src, dst); got != ClassMissing {
			t.Fatalf("classify = %q, want missing", got)
		}
	})

	t.Run("noop when symlink points at src", func(t *testing.T) {
		src := t.TempDir()
		dst := filepath.Join(t.TempDir(), "foo")
		if err := os.Symlink(src, dst); err != nil {
			t.Fatal(err)
		}
		if got := classify(src, dst); got != ClassNoop {
			t.Fatalf("classify = %q, want noop", got)
		}
	})

	t.Run("update when symlink points elsewhere", func(t *testing.T) {
		src := t.TempDir()
		other := t.TempDir()
		dst := filepath.Join(t.TempDir(), "foo")
		if err := os.Symlink(other, dst); err != nil {
			t.Fatal(err)
		}
		if got := classify(src, dst); got != ClassUpdate {
			t.Fatalf("classify = %q, want update", got)
		}
	})

	t.Run("replace when real file exists", func(t *testing.T) {
		src := t.TempDir()
		dstDir := t.TempDir()
		dst := filepath.Join(dstDir, "foo")
		if err := os.WriteFile(dst, []byte("real"), 0o644); err != nil {
			t.Fatal(err)
		}
		if got := classify(src, dst); got != ClassReplace {
			t.Fatalf("classify = %q, want replace", got)
		}
	})

	t.Run("replace when real dir exists", func(t *testing.T) {
		src := t.TempDir()
		dstDir := t.TempDir()
		dst := filepath.Join(dstDir, "foo")
		if err := os.MkdirAll(dst, 0o755); err != nil {
			t.Fatal(err)
		}
		if got := classify(src, dst); got != ClassReplace {
			t.Fatalf("classify = %q, want replace", got)
		}
	})
}

// stubAdapter implements agents.Adapter for planFor tests.
type stubAdapter struct {
	id                                string
	skillsDir, commandsDir, agentsDir string
}

func (s stubAdapter) ID() string          { return s.id }
func (s stubAdapter) DisplayName() string { return s.id }
func (s stubAdapter) Detect() bool        { return true }
func (s stubAdapter) SkillsDir() string   { return s.skillsDir }
func (s stubAdapter) CommandsDir() string { return s.commandsDir }
func (s stubAdapter) AgentsDir() string   { return s.agentsDir }

func (s stubAdapter) SkillsDirFor(_ agents.Scope, _ string) string   { return s.skillsDir }
func (s stubAdapter) CommandsDirFor(_ agents.Scope, _ string) string { return s.commandsDir }
func (s stubAdapter) AgentsDirFor(_ agents.Scope, _ string) string   { return s.agentsDir }

func TestPlanFor(t *testing.T) {
	srcRoot := t.TempDir()
	skillSrc := filepath.Join(srcRoot, "skills", "foo")
	if err := os.MkdirAll(skillSrc, 0o755); err != nil {
		t.Fatal(err)
	}
	items := []skills.Skill{{Name: "foo", SrcPath: skillSrc}}

	t.Run("produces items for non-empty dstDir", func(t *testing.T) {
		dstDir := t.TempDir()
		a := stubAdapter{id: "claude", skillsDir: dstDir}
		got := planFor(a, items, "skills", a.SkillsDir())
		if len(got) != 1 {
			t.Fatalf("planFor returned %d items, want 1", len(got))
		}
		it := got[0]
		if it.Agent != "claude" || it.Kind != "skills" || it.Name != "foo" {
			t.Fatalf("planItem = %+v", it)
		}
		if it.Src != skillSrc {
			t.Fatalf("Src = %q, want %q", it.Src, skillSrc)
		}
		if it.Dst != filepath.Join(dstDir, "foo") {
			t.Fatalf("Dst = %q", it.Dst)
		}
		if it.Class != ClassMissing {
			t.Fatalf("Class = %q, want missing", it.Class)
		}
	})

	t.Run("skips empty dstDir", func(t *testing.T) {
		a := stubAdapter{id: "codex"}
		if got := planFor(a, items, "agents", a.AgentsDir()); got != nil {
			t.Fatalf("planFor with empty dstDir = %v, want nil", got)
		}
	})

	t.Run("skips when no items", func(t *testing.T) {
		a := stubAdapter{id: "claude", skillsDir: t.TempDir()}
		if got := planFor(a, nil, "skills", a.SkillsDir()); got != nil {
			t.Fatalf("planFor with no items = %v, want nil", got)
		}
	})
}

// TestPlanForAllKinds asserts plan items are produced for skills, commands and
// agents when all three destination dirs are configured.
func TestPlanForAllKinds(t *testing.T) {
	srcRoot := t.TempDir()
	skillSrc := filepath.Join(srcRoot, "skills", "s1")
	if err := os.MkdirAll(skillSrc, 0o755); err != nil {
		t.Fatal(err)
	}

	skillItems := []skills.Skill{{Name: "s1", SrcPath: skillSrc}}
	cmdItems := []skills.Skill{{Name: "c1.md", SrcPath: filepath.Join(srcRoot, "c1.md")}}
	agentItems := []skills.Skill{{Name: "a1.md", SrcPath: filepath.Join(srcRoot, "a1.md")}}

	a := stubAdapter{
		id:          "claude",
		skillsDir:   t.TempDir(),
		commandsDir: t.TempDir(),
		agentsDir:   t.TempDir(),
	}

	var all []PlanItem
	all = append(all, planFor(a, skillItems, "skills", a.SkillsDir())...)
	all = append(all, planFor(a, cmdItems, "commands", a.CommandsDir())...)
	all = append(all, planFor(a, agentItems, "agents", a.AgentsDir())...)

	kinds := map[string]bool{}
	for _, it := range all {
		kinds[it.Kind] = true
	}
	for _, want := range []string{"skills", "commands", "agents"} {
		if !kinds[want] {
			t.Fatalf("expected a plan item of kind %q, got kinds %v", want, kinds)
		}
	}
}

// --- ensureLink tests ---

func TestEnsureLink(t *testing.T) {
	t.Run("creates symlink when missing", func(t *testing.T) {
		src := t.TempDir()
		dst := filepath.Join(t.TempDir(), "sub", "link")
		if err := ensureLink(src, dst); err != nil {
			t.Fatalf("ensureLink: %v", err)
		}
		target, err := os.Readlink(dst)
		if err != nil {
			t.Fatalf("readlink: %v", err)
		}
		if target != src {
			t.Fatalf("link target = %q, want %q", target, src)
		}
	})

	t.Run("returns error when dst already exists as wrong symlink", func(t *testing.T) {
		src := t.TempDir()
		dstDir := t.TempDir()
		dst := filepath.Join(dstDir, "link")
		other := t.TempDir()
		// Pre-create a symlink pointing elsewhere.
		if err := os.Symlink(other, dst); err != nil {
			t.Fatalf("setup symlink: %v", err)
		}
		// ensureLink will fail because the destination already exists.
		err := ensureLink(src, dst)
		if err == nil {
			t.Fatal("expected error when dst already exists, got nil")
		}
	})

	t.Run("creates intermediate directories", func(t *testing.T) {
		src := t.TempDir()
		dst := filepath.Join(t.TempDir(), "a", "b", "c", "link")
		if err := ensureLink(src, dst); err != nil {
			t.Fatalf("ensureLink with deep path: %v", err)
		}
		if _, err := os.Lstat(dst); err != nil {
			t.Fatalf("dst not created: %v", err)
		}
	})
}

// --- applyItem tests ---

// newSession creates a backup session that stores under a temp dir, using the
// exported NewSessionAt helper. Since backup.NewSession uses the real home dir
// we call a constructor that accepts an explicit root.
func newTestSession(t *testing.T) *backup.Session {
	t.Helper()
	s, err := backup.NewSessionAt(t.TempDir())
	if err != nil {
		t.Fatalf("backup.NewSessionAt: %v", err)
	}
	return s
}

func TestApplyItem(t *testing.T) {
	t.Run("ClassNoop is a no-op", func(t *testing.T) {
		src := t.TempDir()
		dst := filepath.Join(t.TempDir(), "link")
		it := &PlanItem{Agent: "a", Kind: "skills", Name: "x", Src: src, Dst: dst, Class: ClassNoop}
		applyItem(it, newTestSession(t))
		if it.Err != nil {
			t.Fatalf("ClassNoop should not set Err: %v", it.Err)
		}
		if _, err := os.Lstat(dst); err == nil {
			t.Fatal("ClassNoop must not create dst")
		}
	})

	t.Run("ClassMissing creates symlink", func(t *testing.T) {
		src := t.TempDir()
		dst := filepath.Join(t.TempDir(), "link")
		it := &PlanItem{Agent: "a", Kind: "skills", Name: "x", Src: src, Dst: dst, Class: ClassMissing}
		applyItem(it, newTestSession(t))
		if it.Err != nil {
			t.Fatalf("ClassMissing err: %v", it.Err)
		}
		target, err := os.Readlink(dst)
		if err != nil {
			t.Fatalf("readlink: %v", err)
		}
		if target != src {
			t.Fatalf("link target = %q, want %q", target, src)
		}
	})

	t.Run("ClassUpdate replaces stale symlink", func(t *testing.T) {
		src := t.TempDir()
		dstDir := t.TempDir()
		dst := filepath.Join(dstDir, "link")
		old := t.TempDir()
		// Create a symlink pointing at the old target.
		if err := os.Symlink(old, dst); err != nil {
			t.Fatalf("setup: %v", err)
		}
		it := &PlanItem{Agent: "a", Kind: "skills", Name: "x", Src: src, Dst: dst, Class: ClassUpdate}
		applyItem(it, newTestSession(t))
		if it.Err != nil {
			t.Fatalf("ClassUpdate err: %v", it.Err)
		}
		target, err := os.Readlink(dst)
		if err != nil {
			t.Fatalf("readlink after update: %v", err)
		}
		if target != src {
			t.Fatalf("link target after update = %q, want %q", target, src)
		}
	})

	t.Run("ClassReplace backs up real file and creates symlink", func(t *testing.T) {
		src := t.TempDir()
		dstDir := t.TempDir()
		dst := filepath.Join(dstDir, "foo")
		// Write a regular file at dst to simulate a non-managed entry.
		if err := os.WriteFile(dst, []byte("original content"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		sess := newTestSession(t)
		it := &PlanItem{Agent: "a", Kind: "skills", Name: "foo", Src: src, Dst: dst, Class: ClassReplace}
		applyItem(it, sess)
		if it.Err != nil {
			t.Fatalf("ClassReplace err: %v", it.Err)
		}
		// dst must now be a symlink to src.
		target, err := os.Readlink(dst)
		if err != nil {
			t.Fatalf("readlink: %v", err)
		}
		if target != src {
			t.Fatalf("link target = %q, want %q", target, src)
		}
		// Session must have recorded the backup.
		if sess.Count() != 1 {
			t.Fatalf("session count = %d, want 1", sess.Count())
		}
	})
}

// --- Run end-to-end tests ---

// buildFakeRepo creates a minimal mallard source repo with one skill, one
// command, and one agent definition. Returns the repo root.
func buildFakeRepo(t *testing.T) string {
	t.Helper()
	repo := t.TempDir()
	skillDir := filepath.Join(repo, "skills", "my-skill")
	if err := os.MkdirAll(skillDir, 0o755); err != nil {
		t.Fatalf("mkdir skill: %v", err)
	}
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("---\nname: my-skill\n---\n"), 0o644); err != nil {
		t.Fatalf("write SKILL.md: %v", err)
	}
	cmdDir := filepath.Join(repo, "claude", "commands")
	if err := os.MkdirAll(cmdDir, 0o755); err != nil {
		t.Fatalf("mkdir commands: %v", err)
	}
	if err := os.WriteFile(filepath.Join(cmdDir, "my-cmd.md"), []byte("---\nname: my-cmd\n---\n"), 0o644); err != nil {
		t.Fatalf("write command: %v", err)
	}
	agentDir := filepath.Join(repo, "claude", "agents")
	if err := os.MkdirAll(agentDir, 0o755); err != nil {
		t.Fatalf("mkdir agents: %v", err)
	}
	if err := os.WriteFile(filepath.Join(agentDir, "my-agent.md"), []byte("---\nname: my-agent\n---\n"), 0o644); err != nil {
		t.Fatalf("write agent: %v", err)
	}
	return repo
}

func TestRunDryRun(t *testing.T) {
	repo := buildFakeRepo(t)

	skillsDir := t.TempDir()
	commandsDir := t.TempDir()
	agentsDir := t.TempDir()
	a := stubAdapter{
		id:          "claude",
		skillsDir:   skillsDir,
		commandsDir: commandsDir,
		agentsDir:   agentsDir,
	}
	// Override agents.Detected by using the plan directly via a lower-level test.
	// We exercise planFor + classify directly to confirm dry-run never writes.
	allSkills, _ := skills.DiscoverSkills(repo)
	allCmds, _ := skills.DiscoverCommands(repo)
	allAgents, _ := skills.DiscoverAgents(repo)

	var items []PlanItem
	items = append(items, planFor(a, allSkills, "skills", a.SkillsDir())...)
	items = append(items, planFor(a, allCmds, "commands", a.CommandsDir())...)
	items = append(items, planFor(a, allAgents, "agents", a.AgentsDir())...)

	// In dry-run mode no applyItem is called; all items must be ClassMissing.
	for _, it := range items {
		if it.Class != ClassMissing {
			t.Fatalf("item %q class = %q, want missing before any apply", it.Name, it.Class)
		}
	}
	// Confirm nothing was written.
	for _, it := range items {
		if _, err := os.Lstat(it.Dst); err == nil {
			t.Fatalf("dry-run wrote %s", it.Dst)
		}
	}
}

func TestRunApplyAndReclassify(t *testing.T) {
	repo := buildFakeRepo(t)

	skillsDir := t.TempDir()
	commandsDir := t.TempDir()
	agentsDir := t.TempDir()
	a := stubAdapter{
		id:          "claude",
		skillsDir:   skillsDir,
		commandsDir: commandsDir,
		agentsDir:   agentsDir,
	}

	allSkills, _ := skills.DiscoverSkills(repo)
	allCmds, _ := skills.DiscoverCommands(repo)
	allAgents, _ := skills.DiscoverAgents(repo)

	sess := newTestSession(t)

	var items []PlanItem
	items = append(items, planFor(a, allSkills, "skills", a.SkillsDir())...)
	items = append(items, planFor(a, allCmds, "commands", a.CommandsDir())...)
	items = append(items, planFor(a, allAgents, "agents", a.AgentsDir())...)

	for i := range items {
		applyItem(&items[i], sess)
		if items[i].Err != nil {
			t.Fatalf("applyItem %q: %v", items[i].Name, items[i].Err)
		}
	}

	// After apply all items should be symlinks. Re-planning should yield ClassNoop.
	for _, it := range items {
		newClass := classify(it.Src, it.Dst)
		if newClass != ClassNoop {
			t.Fatalf("after apply, classify(%q) = %q, want noop", it.Name, newClass)
		}
	}
}

func TestRunReplaceRealFile(t *testing.T) {
	repo := buildFakeRepo(t)

	skillsDir := t.TempDir()
	a := stubAdapter{id: "claude", skillsDir: skillsDir}

	allSkills, _ := skills.DiscoverSkills(repo)
	if len(allSkills) == 0 {
		t.Skip("no skills in fake repo")
	}

	// Pre-populate dst with a regular file to trigger ClassReplace.
	dst := filepath.Join(skillsDir, allSkills[0].Name)
	if err := os.WriteFile(dst, []byte("original"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	items := planFor(a, allSkills, "skills", a.SkillsDir())
	if items[0].Class != ClassReplace {
		t.Fatalf("class = %q, want replace", items[0].Class)
	}

	sess := newTestSession(t)
	applyItem(&items[0], sess)
	if items[0].Err != nil {
		t.Fatalf("applyItem: %v", items[0].Err)
	}
	if sess.Count() != 1 {
		t.Fatalf("session count = %d, want 1", sess.Count())
	}
	// dst must now be a symlink.
	if _, err := os.Readlink(dst); err != nil {
		t.Fatalf("dst is not a symlink after replace: %v", err)
	}
}
