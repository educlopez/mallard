package agents

import (
	"os"
	"path/filepath"
)

// windsurfAdapter targets the Windsurf IDE (by Codeium). Its global AI config
// directory is ~/.codeium/windsurf (created on first launch), where it hosts
// markdown skills under ~/.codeium/windsurf/skills.
//
// Windsurf's native workflow format (.windsurf/workflows/*.md) is a distinct,
// non-skill format, so we do NOT map it here — only the claude-style markdown
// skills directory is supported. Windsurf has no slash-command or sub-agent
// directory, so CommandsDir and AgentsDir are empty.
//
// Windsurf is a desktop app with no binary on PATH, so detection is by the
// presence of the ~/.codeium/windsurf config directory.
type windsurfAdapter struct{}

func (windsurfAdapter) ID() string          { return IDWindsurf }
func (windsurfAdapter) DisplayName() string { return "Windsurf" }

func (windsurfAdapter) Detect() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	return dirExists(filepath.Join(home, ".codeium", "windsurf"))
}

func (windsurfAdapter) SkillsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".codeium", "windsurf", "skills")
}

func (windsurfAdapter) CommandsDir() string { return "" }

func (windsurfAdapter) AgentsDir() string { return "" }

func (a windsurfAdapter) SkillsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.SkillsDir, ".windsurf", "skills")
}

func (a windsurfAdapter) CommandsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.CommandsDir)
}

func (a windsurfAdapter) AgentsDirFor(scope Scope, ws string) string {
	return scopedDir(scope, ws, a.AgentsDir)
}
