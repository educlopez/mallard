package agents

import (
	"os"
	"os/exec"
	"path/filepath"
)

// codexAdapter targets OpenAI's Codex CLI.
//
//   - Skills: Codex scans the vendor-neutral $HOME/.agents/skills (user) and
//     <repo>/.agents/skills (project). ~/.codex/skills is no longer a documented
//     location, so it is only kept as a legacy dir to clean up (LegacySkillsDir).
//   - Agents: custom subagents are TOML files in ~/.codex/agents (user) and
//     .codex/agents (project). mallard renders them from claude/agents/*.md.
//   - Commands: Claude commands map 1:1 onto Codex custom prompts
//     (~/.codex/prompts/*.md, invoked as /prompts:<name>, same $ARGUMENTS),
//     so they are symlinked as-is. Custom prompts are user-level only, so the
//     workspace scope has no commands dir. OpenAI marks custom prompts as
//     deprecated in favor of skills; they still load.
type codexAdapter struct{}

func (codexAdapter) ID() string          { return IDCodex }
func (codexAdapter) DisplayName() string { return "Codex" }

func (codexAdapter) Detect() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	if dirExists(filepath.Join(home, ".codex")) {
		return true
	}
	_, err = exec.LookPath("codex")
	return err == nil
}

func (codexAdapter) SkillsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".agents", "skills")
}

func (codexAdapter) CommandsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".codex", "prompts")
}

func (codexAdapter) AgentsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".codex", "agents")
}

func (a codexAdapter) SkillsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.SkillsDir, ".agents", "skills")
}

func (a codexAdapter) CommandsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.CommandsDir)
}

func (a codexAdapter) AgentsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.AgentsDir, ".codex", "agents")
}

// LegacySkillsDirFor is where older mallard releases linked Codex skills.
func (codexAdapter) LegacySkillsDirFor(scope Scope, ws string) string {
	if scope == ScopeWorkspace {
		if ws == "" {
			return ""
		}
		return filepath.Join(ws, ".codex", "skills")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".codex", "skills")
}
