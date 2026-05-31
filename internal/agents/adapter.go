package agents

const (
	IDClaude   = "claude"
	IDCodex    = "codex"
	IDOpenCode = "opencode"
	IDAgents   = "agents"
	IDGemini   = "gemini"
	IDCursor   = "cursor"
	IDWindsurf = "windsurf"
)

// Scope selects where skills/commands/agents are linked.
//
//   - ScopeGlobal links into the agent's home-based dir (e.g. ~/.claude/skills).
//   - ScopeWorkspace links into a project-local .claude/ dir under the current
//     working directory (e.g. <cwd>/.claude/skills).
type Scope string

const (
	ScopeGlobal    Scope = "global"
	ScopeWorkspace Scope = "workspace"
)

// ParseScope validates a user-supplied scope string. Empty defaults to global.
func ParseScope(s string) (Scope, error) {
	switch s {
	case "", string(ScopeGlobal):
		return ScopeGlobal, nil
	case string(ScopeWorkspace):
		return ScopeWorkspace, nil
	default:
		return "", errInvalidScope(s)
	}
}

type errInvalidScope string

func (e errInvalidScope) Error() string {
	return "invalid scope " + string(e) + " (supported: global, workspace)"
}

// Adapter resolves the destination directories for a single agent.
//
// The zero-argument SkillsDir/CommandsDir/AgentsDir methods are retained for
// backwards compatibility and return the GLOBAL (home-based) paths. The *For
// variants are scope-aware: in workspace scope they resolve a project-local
// .claude-style dir under workspaceRoot, returning "" for agents that have no
// sensible workspace location.
type Adapter interface {
	ID() string
	DisplayName() string
	Detect() bool

	// Global home-based directories (unchanged behaviour).
	SkillsDir() string
	CommandsDir() string
	AgentsDir() string

	// Scope-aware directories. workspaceRoot is only consulted in workspace
	// scope; in global scope it is ignored and the home-based path is returned.
	SkillsDirFor(scope Scope, workspaceRoot string) string
	CommandsDirFor(scope Scope, workspaceRoot string) string
	AgentsDirFor(scope Scope, workspaceRoot string) string
}
