package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/educlopez/mallard/cmd"
	"github.com/educlopez/mallard/internal/agents"
)

// version is the mallard release version.
// "dev" is overridden at build time by goreleaser via:
//
//	-ldflags "-X main.version={{.Version}}"
var version = "dev"

func main() {
	repoRoot := repoRootFromEnvOrBinary(version)

	args := os.Args[1:]

	// No args or "install" → TUI
	if len(args) == 0 || args[0] == "install" {
		installArgs := args
		if len(args) > 0 {
			installArgs = args[1:]
		}
		if err := handleInstall(repoRoot, installArgs, version); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	switch args[0] {
	case "doctor":
		dargs, err := cmd.ParseDoctorArgs(args[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if err := cmd.RunDoctor(repoRoot, dargs); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

	case "update":
		uargs, err := cmd.ParseUpdateArgs(args[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if err := cmd.RunUpdate(repoRoot, uargs); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

	case "uninstall":
		unargs, err := cmd.ParseUninstallArgs(args[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if err := cmd.RunUninstall(repoRoot, unargs); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

	case "upgrade":
		upargs, err := cmd.ParseUpgradeArgs(args[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if err := cmd.RunUpgrade(version, upargs); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

	case "registry":
		rargs, err := cmd.ParseRegistryArgs(args[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if err := cmd.RunRegistry(repoRoot, rargs); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

	case "version", "--version", "-v":
		fmt.Printf("mallard %s\n", version)

	case "help", "--help", "-h":
		printHelp()

	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n", args[0])
		printHelp()
		os.Exit(1)
	}
}

func handleInstall(repoRoot string, args []string, version string) error {
	// Parse flags
	agentFlag := ""
	allFlag := false
	scope := agents.ScopeGlobal

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--agent":
			if i+1 < len(args) {
				agentFlag = args[i+1]
				i++
			} else {
				return fmt.Errorf("--agent requires a value")
			}
		case "--all":
			allFlag = true
		case "--scope":
			if i+1 >= len(args) {
				return fmt.Errorf("--scope requires a value")
			}
			sc, err := agents.ParseScope(args[i+1])
			if err != nil {
				return err
			}
			scope = sc
			i++
		}
	}

	if agentFlag != "" {
		return cmd.RunInstallAgent(repoRoot, agentFlag, scope)
	}
	if allFlag {
		return cmd.RunInstallAll(repoRoot, scope)
	}
	// Default: TUI (workspace scope is not offered interactively)
	return cmd.RunInstallTUI(repoRoot, version)
}

// repoRootFromEnvOrBinary resolves the mallard repo root.
//
// Priority:
//  1. MALLARD_DIR env var (explicit override)
//  2. Walk up from the binary looking for a skills/ sibling (dev mode:
//     running `go run .` or `./mallard` from the cloned repo)
//  3. Materialize the embedded source tree into ~/.mallard/source/<version>/
//     (release mode: binary installed via curl-pipe with no sibling repo)
//  4. Fallback to cwd if all of the above fail
func repoRootFromEnvOrBinary(version string) string {
	if dir := os.Getenv("MALLARD_DIR"); dir != "" {
		return dir
	}
	if exe, err := os.Executable(); err == nil {
		// Walk up from the binary looking for a skills/ directory.
		dir := filepath.Dir(exe)
		for {
			if _, err := os.Stat(filepath.Join(dir, "skills")); err == nil {
				return dir
			}
			parent := filepath.Dir(dir)
			if parent == dir {
				break
			}
			dir = parent
		}
	}
	// No sibling repo — materialize the embedded source.
	if dir, err := materializeEmbeddedSource(version); err == nil {
		return dir
	} else {
		fmt.Fprintf(os.Stderr, "warning: could not materialize embedded source: %v\n", err)
	}
	// Last-resort fallback: cwd.
	cwd, _ := os.Getwd()
	return cwd
}

func printHelp() {
	fmt.Print(`mallard — personal Claude Code toolkit

Usage:
  mallard                        Launch interactive TUI installer
  mallard install                Launch interactive TUI installer
  mallard install --agent NAME   Install only to NAME (claude|agents|codex|opencode)
  mallard install --all          Install to all detected agents non-interactively
  mallard install --scope SCOPE  Link into global (default) or workspace (<cwd>/.claude) dirs
  mallard update                 Re-link skills/commands, backing up any conflicting files
  mallard update --dry-run       Show what update would change without touching disk
  mallard update --agent NAME    Update only NAME
  mallard update --yes           Skip confirmation prompts
  mallard update --scope SCOPE   Operate on global (default) or workspace (<cwd>/.claude) dirs
  mallard update --list-backups  List backup batches under ~/.mallard/backups
  mallard update --restore TS    Restore files from backup TS (full stamp or unique prefix)
  mallard update --pin-backup TS Pin backup TS so it is never pruned by the keep-latest GC
  mallard uninstall              Remove mallard-managed symlinks from all detected agents
  mallard uninstall --agent NAME Remove managed symlinks from NAME only
  mallard uninstall --all        Remove managed symlinks from all detected agents
  mallard uninstall --dry-run    Show what would be removed without touching disk
  mallard uninstall --scope SCOPE Remove from global (default) or workspace (<cwd>/.claude) dirs
  mallard upgrade                Self-update the mallard binary to the latest release
  mallard upgrade --check        Report whether a newer release is available
  mallard upgrade --dry-run      Show what upgrade would download/replace
  mallard upgrade --force        Upgrade even on a dev build or non-newer release
  mallard doctor                 Check symlink health per detected agent
  mallard doctor --fix           Repair broken/missing mallard-managed links only
  mallard doctor --scope SCOPE   Check global (default) or workspace (<cwd>/.claude) dirs
  mallard registry               List skills/commands with versions per agent
  mallard registry --source      List source entries from the repo
  mallard registry --json        Emit machine-readable JSON
  mallard version                Print version

Environment:
  MALLARD_DIR              Override repo root (defaults to binary location)
  MALLARD_NO_SELF_UPDATE   Set to 1 to disable 'mallard upgrade' self-replace
`)
}
