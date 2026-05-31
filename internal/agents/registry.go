package agents

import (
	"os"
	"path/filepath"
)

var all = []Adapter{
	claudeAdapter{},
	codexAdapter{},
	openCodeAdapter{},
	genericAdapter{},
	geminiAdapter{},
	cursorAdapter{},
	windsurfAdapter{},
}

func All() []Adapter {
	out := make([]Adapter, len(all))
	copy(out, all)
	return out
}

func ByID(id string) (Adapter, bool) {
	for _, a := range all {
		if a.ID() == id {
			return a, true
		}
	}
	return nil, false
}

func Detected() []Adapter {
	var out []Adapter
	for _, a := range all {
		if a.Detect() {
			out = append(out, a)
		}
	}
	return out
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// scopedDir resolves a destination dir for the given scope.
//
//   - global scope returns the adapter's home-based path via globalFn.
//   - workspace scope joins workspaceRoot with the supplied path segments
//     (e.g. ".claude", "skills"), yielding <ws>/.claude/skills. An empty
//     workspaceRoot in workspace scope yields "" (caller should skip).
//
// If globalFn returns "" (the agent has no such dir at all), the workspace
// variant also returns "" so the dir is skipped consistently in both scopes.
func scopedDir(scope Scope, workspaceRoot string, globalFn func() string, segments ...string) string {
	if scope == ScopeWorkspace {
		// No segments means this agent has no workspace location for this kind.
		if workspaceRoot == "" || len(segments) == 0 {
			return ""
		}
		parts := append([]string{workspaceRoot}, segments...)
		return filepath.Join(parts...)
	}
	return globalFn()
}
