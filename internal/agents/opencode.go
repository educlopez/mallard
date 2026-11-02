package agents

import (
	"os"
	"os/exec"
	"path/filepath"
)

// openCodeAdapter targets OpenCode. Skills and commands are plain markdown and
// are symlinked straight from the source. Agents live in
// ~/.config/opencode/agents (.opencode/agents in a project) but use their own
// frontmatter (mode, permission, provider/model), so mallard renders an
// OpenCode variant of each claude/agents/*.md and links that.
type openCodeAdapter struct{}

func (openCodeAdapter) ID() string          { return IDOpenCode }
func (openCodeAdapter) DisplayName() string { return "OpenCode" }

func (openCodeAdapter) Detect() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	if dirExists(filepath.Join(home, ".config", "opencode")) {
		return true
	}
	_, err = exec.LookPath("opencode")
	return err == nil
}

func (openCodeAdapter) SkillsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".config", "opencode", "skills")
}

func (openCodeAdapter) CommandsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".config", "opencode", "commands")
}

func (openCodeAdapter) AgentsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".config", "opencode", "agents")
}

func (a openCodeAdapter) SkillsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.SkillsDir, ".opencode", "skills")
}

func (a openCodeAdapter) CommandsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.CommandsDir, ".opencode", "commands")
}

func (a openCodeAdapter) AgentsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.AgentsDir, ".opencode", "agents")
}
