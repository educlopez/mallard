package updater

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/educlopez/duck-ai/internal/skills"
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
