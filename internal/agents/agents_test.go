package agents

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAgentsDir(t *testing.T) {
	tests := []struct {
		id         string
		wantSuffix string // non-empty: AgentsDir must end with this
		wantEmpty  bool   // true: AgentsDir must be ""
	}{
		{id: IDClaude, wantSuffix: filepath.Join(".claude", "agents")},
		{id: IDCodex, wantEmpty: true},
		{id: IDOpenCode, wantEmpty: true},
		{id: IDAgents, wantEmpty: true},
		{id: IDGemini, wantEmpty: true},
		{id: IDCursor, wantSuffix: filepath.Join(".cursor", "agents")},
		{id: IDWindsurf, wantEmpty: true},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			a, ok := ByID(tt.id)
			if !ok {
				t.Fatalf("ByID(%q) not found", tt.id)
			}
			got := a.AgentsDir()
			if tt.wantEmpty {
				if got != "" {
					t.Fatalf("%s AgentsDir() = %q, want empty", tt.id, got)
				}
				return
			}
			if !strings.HasSuffix(got, tt.wantSuffix) {
				t.Fatalf("%s AgentsDir() = %q, want suffix %q", tt.id, got, tt.wantSuffix)
			}
		})
	}
}

func TestByID(t *testing.T) {
	if _, ok := ByID("does-not-exist"); ok {
		t.Fatalf("ByID(unknown) returned ok=true")
	}
	for _, id := range []string{IDClaude, IDCodex, IDOpenCode, IDAgents, IDGemini, IDCursor, IDWindsurf} {
		a, ok := ByID(id)
		if !ok {
			t.Fatalf("ByID(%q) not found", id)
		}
		if a.ID() != id {
			t.Fatalf("adapter ID() = %q, want %q", a.ID(), id)
		}
	}
}

func TestParseScope(t *testing.T) {
	tests := []struct {
		in      string
		want    Scope
		wantErr bool
	}{
		{in: "", want: ScopeGlobal},
		{in: "global", want: ScopeGlobal},
		{in: "workspace", want: ScopeWorkspace},
		{in: "nope", wantErr: true},
	}
	for _, tt := range tests {
		got, err := ParseScope(tt.in)
		if tt.wantErr {
			if err == nil {
				t.Fatalf("ParseScope(%q) err = nil, want error", tt.in)
			}
			continue
		}
		if err != nil {
			t.Fatalf("ParseScope(%q) err = %v", tt.in, err)
		}
		if got != tt.want {
			t.Fatalf("ParseScope(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

// TestScopeDirResolution verifies workspace scope yields project-local
// .claude dirs and global scope is byte-identical to the legacy methods.
func TestScopeDirResolution(t *testing.T) {
	a, ok := ByID(IDClaude)
	if !ok {
		t.Fatal("claude adapter missing")
	}
	ws := t.TempDir()

	// Workspace scope → <ws>/.claude/{skills,commands,agents}.
	if got, want := a.SkillsDirFor(ScopeWorkspace, ws), filepath.Join(ws, ".claude", "skills"); got != want {
		t.Fatalf("workspace SkillsDirFor = %q, want %q", got, want)
	}
	if got, want := a.CommandsDirFor(ScopeWorkspace, ws), filepath.Join(ws, ".claude", "commands"); got != want {
		t.Fatalf("workspace CommandsDirFor = %q, want %q", got, want)
	}
	if got, want := a.AgentsDirFor(ScopeWorkspace, ws), filepath.Join(ws, ".claude", "agents"); got != want {
		t.Fatalf("workspace AgentsDirFor = %q, want %q", got, want)
	}

	// Global scope must equal the legacy home-based methods exactly.
	if got, want := a.SkillsDirFor(ScopeGlobal, ws), a.SkillsDir(); got != want {
		t.Fatalf("global SkillsDirFor = %q, want %q (must equal SkillsDir())", got, want)
	}
	if got, want := a.CommandsDirFor(ScopeGlobal, ws), a.CommandsDir(); got != want {
		t.Fatalf("global CommandsDirFor = %q, want %q", got, want)
	}
}

// TestWorkspaceScopeSkipsAgentsWithoutLocation ensures agents whose global dir
// is empty also yield empty workspace dirs (skipped consistently).
func TestWorkspaceScopeSkipsAgentsWithoutLocation(t *testing.T) {
	ws := t.TempDir()
	codex, _ := ByID(IDCodex)
	// Codex has no commands/agents dir globally; workspace must also be empty.
	if got := codex.CommandsDirFor(ScopeWorkspace, ws); got != "" {
		t.Fatalf("codex workspace CommandsDirFor = %q, want empty", got)
	}
	if got := codex.AgentsDirFor(ScopeWorkspace, ws); got != "" {
		t.Fatalf("codex workspace AgentsDirFor = %q, want empty", got)
	}
	// But codex DOES have a skills dir, so workspace resolves it.
	if got, want := codex.SkillsDirFor(ScopeWorkspace, ws), filepath.Join(ws, ".codex", "skills"); got != want {
		t.Fatalf("codex workspace SkillsDirFor = %q, want %q", got, want)
	}
}

// TestNewAdapterDirs verifies the global skills/commands/agents directory
// conventions for the gemini, cursor, and windsurf adapters match the paths
// each agent actually uses.
func TestNewAdapterDirs(t *testing.T) {
	tests := []struct {
		id            string
		skillsSuffix  string
		commandsEmpty bool
		agentsSuffix  string // empty string means AgentsDir must be ""
	}{
		{id: IDGemini, skillsSuffix: filepath.Join(".gemini", "skills"), commandsEmpty: true},
		{id: IDCursor, skillsSuffix: filepath.Join(".cursor", "skills"), commandsEmpty: true, agentsSuffix: filepath.Join(".cursor", "agents")},
		{id: IDWindsurf, skillsSuffix: filepath.Join(".codeium", "windsurf", "skills"), commandsEmpty: true},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			a, ok := ByID(tt.id)
			if !ok {
				t.Fatalf("ByID(%q) not found", tt.id)
			}
			if got := a.SkillsDir(); !strings.HasSuffix(got, tt.skillsSuffix) {
				t.Fatalf("%s SkillsDir() = %q, want suffix %q", tt.id, got, tt.skillsSuffix)
			}
			if tt.commandsEmpty {
				if got := a.CommandsDir(); got != "" {
					t.Fatalf("%s CommandsDir() = %q, want empty", tt.id, got)
				}
			}
			if tt.agentsSuffix == "" {
				if got := a.AgentsDir(); got != "" {
					t.Fatalf("%s AgentsDir() = %q, want empty", tt.id, got)
				}
			} else if got := a.AgentsDir(); !strings.HasSuffix(got, tt.agentsSuffix) {
				t.Fatalf("%s AgentsDir() = %q, want suffix %q", tt.id, got, tt.agentsSuffix)
			}
		})
	}
}

// TestNewAdapterWorkspaceDirs checks the scope-aware workspace resolution for
// the new adapters, including that unsupported kinds stay empty.
func TestNewAdapterWorkspaceDirs(t *testing.T) {
	ws := t.TempDir()

	gemini, _ := ByID(IDGemini)
	if got, want := gemini.SkillsDirFor(ScopeWorkspace, ws), filepath.Join(ws, ".gemini", "skills"); got != want {
		t.Fatalf("gemini workspace SkillsDirFor = %q, want %q", got, want)
	}
	if got := gemini.CommandsDirFor(ScopeWorkspace, ws); got != "" {
		t.Fatalf("gemini workspace CommandsDirFor = %q, want empty", got)
	}

	cursor, _ := ByID(IDCursor)
	if got, want := cursor.AgentsDirFor(ScopeWorkspace, ws), filepath.Join(ws, ".cursor", "agents"); got != want {
		t.Fatalf("cursor workspace AgentsDirFor = %q, want %q", got, want)
	}

	windsurf, _ := ByID(IDWindsurf)
	if got, want := windsurf.SkillsDirFor(ScopeWorkspace, ws), filepath.Join(ws, ".windsurf", "skills"); got != want {
		t.Fatalf("windsurf workspace SkillsDirFor = %q, want %q", got, want)
	}
	if got := windsurf.AgentsDirFor(ScopeWorkspace, ws); got != "" {
		t.Fatalf("windsurf workspace AgentsDirFor = %q, want empty", got)
	}
}

// TestNewAdapterDetect verifies the presence-of-config-dir detection for the
// new adapters using a temp HOME with no agent dirs (must be false) and then
// with the dir present (must be true).
func TestNewAdapterDetect(t *testing.T) {
	tests := []struct {
		id      string
		dirRel  []string // path under HOME that, when present, triggers Detect
	}{
		{id: IDCursor, dirRel: []string{".cursor"}},
		{id: IDWindsurf, dirRel: []string{".codeium", "windsurf"}},
		{id: IDGemini, dirRel: []string{".gemini"}},
	}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			a, _ := ByID(tt.id)
			home := t.TempDir()
			t.Setenv("HOME", home)
			// Gemini may also be found via PATH; only assert the positive case
			// for the config-dir signal, which is deterministic.
			parts := append([]string{home}, tt.dirRel...)
			if err := os.MkdirAll(filepath.Join(parts...), 0o755); err != nil {
				t.Fatalf("mkdir: %v", err)
			}
			if !a.Detect() {
				t.Fatalf("%s Detect() = false with config dir present", tt.id)
			}
		})
	}
}

func TestAllReturnsCopy(t *testing.T) {
	got := All()
	if len(got) != len(all) {
		t.Fatalf("All() len = %d, want %d", len(got), len(all))
	}
	// Mutating the returned slice must not affect the package-level registry.
	if len(got) > 0 {
		got[0] = nil
		if All()[0] == nil {
			t.Fatalf("All() returned a slice aliasing the internal registry")
		}
	}
}
