package agents

import (
	"os"
	"path/filepath"
)

// cursorAdapter targets the Cursor IDE, a desktop app whose config lives under
// ~/.cursor. It hosts markdown skills in ~/.cursor/skills and native sub-agents
// in ~/.cursor/agents. Cursor has no claude-style slash-command directory, so
// CommandsDir is empty. Cursor is a desktop app with no binary on PATH, so
// detection is purely by the presence of the ~/.cursor config directory.
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

func (cursorAdapter) CommandsDir() string { return "" }

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
	return scopedDir(scope, ws, a.CommandsDir)
}

func (a cursorAdapter) AgentsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.AgentsDir, ".cursor", "agents")
}
