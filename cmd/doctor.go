package cmd

import (
	"fmt"
	"os"

	"github.com/educlopez/duck-ai/internal/agents"
	"github.com/educlopez/duck-ai/internal/reports"
)

// DoctorArgs holds parsed flags for the doctor command.
type DoctorArgs struct {
	Fix   bool
	Scope agents.Scope
}

// ParseDoctorArgs parses the doctor subcommand flags.
func ParseDoctorArgs(args []string) (DoctorArgs, error) {
	var out DoctorArgs
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--fix":
			out.Fix = true
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
			return out, fmt.Errorf("unknown doctor flag %q", args[i])
		}
	}
	return out, nil
}

// RunDoctor prints the doctor report to stdout. With Fix set, it conservatively
// repairs broken/missing duck-ai-managed symlinks instead of just reporting.
// The TUI calls reports.Doctor directly with a captured writer.
func RunDoctor(repoRoot string, args DoctorArgs) error {
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
	if args.Fix {
		_, err := reports.DoctorFix(os.Stdout, repoRoot, scope, ws)
		return err
	}
	return reports.Doctor(os.Stdout, repoRoot, scope, ws)
}
