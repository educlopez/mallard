package agents

import (
	"os"
	"os/exec"
	"path/filepath"
)

// geminiAdapter targets Google's Gemini CLI, whose global config lives under
// ~/.gemini. It hosts markdown skills in ~/.gemini/skills. Gemini exposes no
// claude-style slash-command or sub-agent directories, so CommandsDir and
// AgentsDir are empty (consistent with codex/opencode).
type geminiAdapter struct{}

func (geminiAdapter) ID() string          { return IDGemini }
func (geminiAdapter) DisplayName() string { return "Gemini CLI" }

func (geminiAdapter) Detect() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	if dirExists(filepath.Join(home, ".gemini")) {
		return true
	}
	_, err = exec.LookPath("gemini")
	return err == nil
}

func (geminiAdapter) SkillsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".gemini", "skills")
}

func (geminiAdapter) CommandsDir() string { return "" }

func (geminiAdapter) AgentsDir() string { return "" }

func (a geminiAdapter) SkillsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.SkillsDir, ".gemini", "skills")
}

func (a geminiAdapter) CommandsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.CommandsDir)
}

func (a geminiAdapter) AgentsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.AgentsDir)
}
