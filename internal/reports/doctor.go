// Package reports holds the writer-based renderers for `duck-ai doctor` and
// `duck-ai registry`. Both the cmd CLI wrappers and the TUI screens consume
// these functions. Keeping them here breaks the cycle between cmd and
// internal/tui (cmd imports internal/tui via install.go).
package reports

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/educlopez/duck-ai/internal/agents"
	"github.com/educlopez/duck-ai/internal/skills"
)

// Doctor writes the duck-ai doctor report to w. repoRoot is the absolute or
// relative path to the duck-ai source repo.
func Doctor(w io.Writer, repoRoot string) error {
	absRepo, err := filepath.Abs(repoRoot)
	if err != nil {
		absRepo = repoRoot
	}

	for _, a := range agents.All() {
		detected := a.Detect()
		skillsDir := a.SkillsDir()
		commandsDir := a.CommandsDir()
		agentsDir := a.AgentsDir()

		fmt.Fprintf(w, "\n  %s (%s)\n", a.DisplayName(), a.ID())
		fmt.Fprintf(w, "    detected:     %s\n", yesNo(detected))
		fmt.Fprintf(w, "    skills dir:   %s\n", displayPath(skillsDir))
		fmt.Fprintf(w, "    commands dir: %s\n", displayPath(commandsDir))
		fmt.Fprintf(w, "    agents dir:   %s\n", displayPath(agentsDir))

		if !detected {
			continue
		}

		skillsManaged, skillsUnmanaged := scanDir(skillsDir, absRepo)
		commandsManaged, commandsUnmanaged := scanDir(commandsDir, absRepo)
		agentsManaged, agentsUnmanaged := scanDir(agentsDir, absRepo)
		managed := skillsManaged + commandsManaged + agentsManaged
		fmt.Fprintf(w, "    managed:      %d duck-ai symlinks\n", managed)

		unmanaged := append([]driftEntry{}, skillsUnmanaged...)
		unmanaged = append(unmanaged, commandsUnmanaged...)
		unmanaged = append(unmanaged, agentsUnmanaged...)
		if len(unmanaged) > 0 {
			fmt.Fprintf(w, "    unmanaged:    %d entries (not managed by duck-ai)\n", len(unmanaged))
			for _, u := range unmanaged {
				fmt.Fprintf(w, "      - %s (%s)\n", u.relPath, u.kind)
			}
			fmt.Fprintf(w, "    hint: run `duck-ai update` to absorb colliding entries; non-colliding ones will be left alone.\n")
		}
	}

	return nil
}

// FixResult records a single repair performed by DoctorFix.
type FixResult struct {
	Agent  string
	Kind   string // "skills" | "commands" | "agents"
	Name   string
	Action string // "relinked-broken" | "created-missing"
	Dst    string
	Src    string
}

// DoctorFix conservatively repairs ONLY duck-ai-managed breakage in detected
// agents' directories and writes a summary of what it fixed to w. It performs
// exactly two kinds of repair:
//
//  1. relinked-broken — a symlink whose target points inside the duck-ai repo
//     but no longer resolves (the source moved/renamed). It is removed and, if
//     a current source item of the same name exists, re-linked to it.
//  2. created-missing — a current source item that has no entry at all in the
//     destination dir. A fresh symlink is created.
//
// It never touches real files/dirs, never touches symlinks pointing outside
// the repo, and never removes a broken link when no matching source exists
// (that is genuine drift the user must resolve). Returns the repairs made.
func DoctorFix(w io.Writer, repoRoot string) ([]FixResult, error) {
	absRepo, err := filepath.Abs(repoRoot)
	if err != nil {
		absRepo = repoRoot
	}

	allSkills, err := skills.DiscoverSkills(absRepo)
	if err != nil {
		return nil, err
	}
	allCommands, err := skills.DiscoverCommands(absRepo)
	if err != nil {
		return nil, err
	}
	allAgents, err := skills.DiscoverAgents(absRepo)
	if err != nil {
		return nil, err
	}

	var fixes []FixResult
	for _, a := range agents.All() {
		if !a.Detect() {
			continue
		}
		fixes = append(fixes, fixDir(a.ID(), "skills", a.SkillsDir(), allSkills, absRepo)...)
		fixes = append(fixes, fixDir(a.ID(), "commands", a.CommandsDir(), allCommands, absRepo)...)
		fixes = append(fixes, fixDir(a.ID(), "agents", a.AgentsDir(), allAgents, absRepo)...)
	}

	if len(fixes) == 0 {
		fmt.Fprintf(w, "\n  Nothing to fix — no broken or missing duck-ai links found.\n")
		return fixes, nil
	}
	fmt.Fprintf(w, "\n  Fixed %d duck-ai-managed link(s):\n", len(fixes))
	for _, f := range fixes {
		fmt.Fprintf(w, "    [%s] %s/%s — %s\n", f.Agent, f.Kind, f.Name, f.Action)
	}
	return fixes, nil
}

// fixDir repairs broken and missing duck-ai-managed links in a single agent
// destination dir. src is the list of current source items for this kind.
func fixDir(agentID, kind, dstDir string, src []skills.Skill, repoRoot string) []FixResult {
	if dstDir == "" {
		return nil
	}

	// Map source name -> source path for quick lookup.
	srcByName := make(map[string]string, len(src))
	for _, s := range src {
		srcByName[s.Name] = s.SrcPath
	}

	var fixes []FixResult

	// 1. Repair broken managed symlinks already present in dstDir.
	present := map[string]bool{}
	if entries, err := os.ReadDir(dstDir); err == nil {
		for _, e := range entries {
			name := e.Name()
			if strings.HasPrefix(name, ".") {
				continue
			}
			full := filepath.Join(dstDir, name)
			info, err := os.Lstat(full)
			if err != nil {
				continue
			}
			if info.Mode()&os.ModeSymlink == 0 {
				continue // real file/dir — never touch
			}
			present[name] = true

			target, err := os.Readlink(full)
			if err != nil {
				continue
			}
			absTarget := target
			if !filepath.IsAbs(absTarget) {
				absTarget = filepath.Join(filepath.Dir(full), absTarget)
			}
			// Only consider links that point INTO the duck-ai repo (managed).
			if !(absTarget == repoRoot || strings.HasPrefix(absTarget, repoRoot+string(filepath.Separator))) {
				continue
			}
			// Healthy? leave it.
			if _, err := os.Stat(absTarget); err == nil {
				continue
			}
			// Broken managed link. Only repair if a current source still exists.
			newSrc, ok := srcByName[name]
			if !ok {
				continue // source genuinely gone — user must resolve drift
			}
			if err := os.Remove(full); err != nil {
				continue
			}
			if err := os.Symlink(newSrc, full); err != nil {
				continue
			}
			fixes = append(fixes, FixResult{
				Agent: agentID, Kind: kind, Name: name,
				Action: "relinked-broken", Dst: full, Src: newSrc,
			})
		}
	}

	// 2. Create links for current source items missing from dstDir entirely.
	for _, s := range src {
		if present[s.Name] {
			continue
		}
		full := filepath.Join(dstDir, s.Name)
		// Guard: skip if anything exists at the path (e.g. a non-symlink we
		// did not enumerate due to a race) — never clobber.
		if _, err := os.Lstat(full); err == nil {
			continue
		}
		if err := os.MkdirAll(dstDir, 0o755); err != nil {
			continue
		}
		if err := os.Symlink(s.SrcPath, full); err != nil {
			continue
		}
		fixes = append(fixes, FixResult{
			Agent: agentID, Kind: kind, Name: s.Name,
			Action: "created-missing", Dst: full, Src: s.SrcPath,
		})
	}

	return fixes
}

func yesNo(b bool) string {
	if b {
		return "y"
	}
	return "n"
}

func displayPath(p string) string {
	if p == "" {
		return "(none)"
	}
	return p
}

// driftEntry describes a single unmanaged file or directory found in an agent
// directory.
type driftEntry struct {
	relPath string
	kind    string // "file" or "dir"
}

// scanDir walks dir one level deep and returns the count of duck-ai-managed
// symlinks plus a slice of unmanaged entries (real files/dirs, or symlinks
// pointing outside repoRoot). Hidden entries are skipped.
func scanDir(dir, repoRoot string) (managed int, unmanaged []driftEntry) {
	if dir == "" {
		return 0, nil
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, nil
	}
	prefix := filepath.Base(dir) // "skills" or "commands"
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		full := filepath.Join(dir, name)
		info, err := os.Lstat(full)
		if err != nil {
			continue
		}
		if info.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(full)
			if err == nil {
				absTarget := target
				if !filepath.IsAbs(absTarget) {
					absTarget = filepath.Join(filepath.Dir(full), absTarget)
				}
				if strings.HasPrefix(absTarget, repoRoot+string(filepath.Separator)) || absTarget == repoRoot {
					managed++
					continue
				}
			}
		}
		kind := "file"
		if info.IsDir() {
			kind = "dir"
		}
		unmanaged = append(unmanaged, driftEntry{
			relPath: filepath.Join(prefix, name),
			kind:    kind,
		})
	}
	return managed, unmanaged
}
