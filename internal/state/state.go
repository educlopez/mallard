// Package state persists small, non-sensitive user selections for mallard
// across runs in ~/.mallard/state.json. It is deliberately minimal and always
// graceful: a missing or corrupt file yields zero-value defaults and never
// blocks a run.
package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const stateDir = ".mallard"
const stateFile = "state.json"

// State holds persisted user selections from the last interactive run. New
// fields must be optional (omitempty) so older state files keep loading.
type State struct {
	// LastAgents is the set of agent IDs the user last selected in the TUI.
	LastAgents []string `json:"last_agents,omitempty"`
	// LastSkills is the set of skill names the user last selected in the TUI.
	LastSkills []string `json:"last_skills,omitempty"`
}

// Path returns the absolute path to the state file under the user's home dir.
func Path() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home dir: %w", err)
	}
	return filepath.Join(home, stateDir, stateFile), nil
}

// Load reads and decodes the state file. It is intentionally forgiving: a
// missing file, an unreadable file, or a corrupt/undecodable file all yield a
// zero-value State with no error, so a bad state file can never block a run.
func Load() State {
	path, err := Path()
	if err != nil {
		return State{}
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return State{} // missing or unreadable → defaults
	}
	var s State
	if err := json.Unmarshal(data, &s); err != nil {
		return State{} // corrupt → defaults, never panic
	}
	return s
}

// Save persists the state to ~/.mallard/state.json, creating the directory if
// needed. Errors are returned for callers that care, but a failed save must
// never be treated as fatal by the TUI.
func Save(s State) error {
	path, err := Path()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir state dir: %w", err)
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal state: %w", err)
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}
