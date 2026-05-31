package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/educlopez/duck-ai/internal/agents"
)

// UninstallArgs captures parsed flags for `duck-ai uninstall`.
type UninstallArgs struct {
	AgentID string
	All     bool
	DryRun  bool
	Scope   agents.Scope
}

// ParseUninstallArgs parses the uninstall subcommand flags.
//
// Supported flags:
//
//	--agent NAME  uninstall from a single agent
//	--all         uninstall from every detected agent (default behaviour)
//	--dry-run     report what would be removed without touching disk
func ParseUninstallArgs(args []string) (UninstallArgs, error) {
	var out UninstallArgs
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--agent":
			if i+1 >= len(args) {
				return out, fmt.Errorf("--agent requires a value")
			}
			out.AgentID = args[i+1]
			i++
		case "--all":
			out.All = true
		case "--dry-run":
			out.DryRun = true
		case "--scope":
			if i+1 >= len(args) {
				return out, fmt.Errorf("--scope requires a value")
			}
			sc, err := agents.ParseScope(args[i+1])
			if err != nil {
				return out, err
			}
			out.Scope = sc
			i++
		default:
			return out, fmt.Errorf("unknown uninstall flag %q", args[i])
		}
	}
	if out.AgentID != "" && out.All {
		return out, fmt.Errorf("--agent and --all are mutually exclusive")
	}
	return out, nil
}

// RunUninstall removes ONLY duck-ai-managed symlinks (those whose resolved
// target points into repoRoot) from the detected agents' skills/commands/agents
// directories. It is the mirror of install: real files/dirs and symlinks
// pointing outside the repo are unmanaged and always left untouched.
//
// It prints its report to os.Stdout; the testable core is uninstall().
func RunUninstall(repoRoot string, args UninstallArgs) error {
	adapters, err := selectUninstallAgents(args)
	if err != nil {
		return err
	}
	scope := args.Scope
	if scope == "" {
		scope = agents.ScopeGlobal
	}
	ws := ""
	if scope == agents.ScopeWorkspace {
		if cwd, cerr := os.Getwd(); cerr == nil {
			ws = cwd
		}
	}
	return uninstall(os.Stdout, repoRoot, adapters, args.DryRun, scope, ws)
}

// selectUninstallAgents resolves which adapters to operate on from the parsed
// flags. With no flag (or --all) it returns every detected agent; with --agent
// it returns just that one (still required to be detected).
func selectUninstallAgents(args UninstallArgs) ([]agents.Adapter, error) {
	if args.AgentID != "" {
		a, ok := agents.ByID(args.AgentID)
		if !ok {
			return nil, fmt.Errorf("unknown agent %q (supported: claude, codex, opencode, agents)", args.AgentID)
		}
		return []agents.Adapter{a}, nil
	}
	// Default and --all both mean: every detected agent.
	return agents.Detected(), nil
}

// uninstall is the pure-ish core: it removes managed symlinks across the given
// adapters and writes a per-agent summary to w. With dryRun set it changes
// nothing on disk and only reports what it would remove.
func uninstall(w io.Writer, repoRoot string, adapters []agents.Adapter, dryRun bool, scope agents.Scope, workspaceRoot string) error {
	if scope == "" {
		scope = agents.ScopeGlobal
	}
	absRepo, err := filepath.Abs(repoRoot)
	if err != nil {
		absRepo = repoRoot
	}

	if dryRun {
		fmt.Fprintln(w, "\n  duck-ai uninstall (dry run) — no files were modified.")
	} else {
		fmt.Fprintln(w, "\n  duck-ai uninstall")
	}

	if len(adapters) == 0 {
		fmt.Fprintln(w, "  No agents detected.")
		return nil
	}

	totalRemoved, totalSkipped := 0, 0
	for _, a := range adapters {
		fmt.Fprintf(w, "\n  Agent: %s\n", a.ID())
		if !a.Detect() {
			fmt.Fprintf(w, "    -  not detected, skipping\n")
			continue
		}

		removed, skipped := 0, 0
		for _, dir := range []string{
			a.SkillsDirFor(scope, workspaceRoot),
			a.CommandsDirFor(scope, workspaceRoot),
			a.AgentsDirFor(scope, workspaceRoot),
		} {
			r, s := unlinkManagedInDir(w, dir, absRepo, dryRun)
			removed += r
			skipped += s
		}
		fmt.Fprintf(w, "    removed: %d  skipped (unmanaged): %d\n", removed, skipped)
		totalRemoved += removed
		totalSkipped += skipped
	}

	verb := "removed"
	if dryRun {
		verb = "would remove"
	}
	fmt.Fprintf(w, "\n  Done — %s %d managed symlink(s); left %d unmanaged entr(ies) untouched.\n",
		verb, totalRemoved, totalSkipped)
	return nil
}

// unlinkManagedInDir scans dir one level deep and removes every duck-ai-managed
// symlink (one resolving to a target inside repoRoot). It returns the count
// removed and the count of unmanaged entries left in place. Hidden entries are
// skipped. With dryRun set it reports but does not remove.
//
// "Managed" is determined exactly like reports.scanDir: resolve the link target
// (making relative targets absolute against the link's own dir) and check it is
// repoRoot itself or lives under repoRoot.
func unlinkManagedInDir(w io.Writer, dir, repoRoot string, dryRun bool) (removed, skipped int) {
	if dir == "" {
		return 0, 0
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, 0
	}
	prefix := filepath.Base(dir)
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
		if info.Mode()&os.ModeSymlink == 0 {
			// Real file or dir — unmanaged, never touch.
			skipped++
			continue
		}
		target, err := os.Readlink(full)
		if err != nil {
			skipped++
			continue
		}
		absTarget := target
		if !filepath.IsAbs(absTarget) {
			absTarget = filepath.Join(filepath.Dir(full), absTarget)
		}
		managed := absTarget == repoRoot ||
			strings.HasPrefix(absTarget, repoRoot+string(filepath.Separator))
		if !managed {
			// Symlink pointing outside the repo — unmanaged, leave it.
			skipped++
			continue
		}
		rel := filepath.Join(prefix, name)
		if dryRun {
			fmt.Fprintf(w, "    -  would remove %s\n", rel)
			removed++
			continue
		}
		if err := os.Remove(full); err != nil {
			fmt.Fprintf(w, "    x  %s (error: %v)\n", rel, err)
			continue
		}
		fmt.Fprintf(w, "    -  removed %s\n", rel)
		removed++
	}
	return removed, skipped
}
