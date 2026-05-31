package agents

import (
	"os"
	"os/exec"
	"path/filepath"
)

type claudeAdapter struct{}

func (claudeAdapter) ID() string          { return IDClaude }
func (claudeAdapter) DisplayName() string { return "Claude Code" }

func (claudeAdapter) Detect() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	if dirExists(filepath.Join(home, ".claude")) {
		return true
	}
	_, err = exec.LookPath("claude")
	return err == nil
}

func (claudeAdapter) SkillsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".claude", "skills")
}

func (claudeAdapter) CommandsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".claude", "commands")
}

func (claudeAdapter) AgentsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".claude", "agents")
}

func (a claudeAdapter) SkillsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.SkillsDir, ".claude", "skills")
}

func (a claudeAdapter) CommandsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.CommandsDir, ".claude", "commands")
}

func (a claudeAdapter) AgentsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.AgentsDir, ".claude", "agents")
}
