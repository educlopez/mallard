package agents

import (
	"os"
	"path/filepath"
)

// cursorAdapter targets the Cursor IDE, a desktop app whose config lives under
// ~/.cursor. Skills (~/.cursor/skills) and slash commands (~/.cursor/commands)
// are plain markdown and are symlinked straight from the source. Sub-agents
// (~/.cursor/agents) use their own frontmatter (model: inherit, readonly), so
// mallard renders a Cursor variant of each claude/agents/*.md and links that.
//
// Cursor 2.4 had a regression where global ~/.cursor/commands were ignored;
// current releases (3.x) load them again.
//
// Cursor is a desktop app with no binary on PATH, so detection is purely by the
// presence of the ~/.cursor config directory.
type cursorAdapter struct{}

func (cursorAdapter) ID() string          { return IDCursor }
func (cursorAdapter) DisplayName() string { return "Cursor" }

func (cursorAdapter) Detect() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	return dirExists(filepath.Join(home, ".cursor"))
}

func (cursorAdapter) SkillsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".cursor", "skills")
}

func (cursorAdapter) CommandsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".cursor", "commands")
}

func (cursorAdapter) AgentsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".cursor", "agents")
}

func (a cursorAdapter) SkillsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.SkillsDir, ".cursor", "skills")
}

func (a cursorAdapter) CommandsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.CommandsDir, ".cursor", "commands")
}

func (a cursorAdapter) AgentsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.AgentsDir, ".cursor", "agents")
}
