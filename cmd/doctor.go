package cmd

import (
	"fmt"
	"os"

	"github.com/educlopez/duck-ai/internal/reports"
)

// DoctorArgs holds parsed flags for the doctor command.
type DoctorArgs struct {
	Fix bool
}

// ParseDoctorArgs parses the doctor subcommand flags.
func ParseDoctorArgs(args []string) (DoctorArgs, error) {
	var out DoctorArgs
	for _, a := range args {
		switch a {
		case "--fix":
			out.Fix = true
		default:
			return out, fmt.Errorf("unknown doctor flag %q", a)
		}
	}
	return out, nil
}

// RunDoctor prints the doctor report to stdout. With Fix set, it conservatively
// repairs broken/missing duck-ai-managed symlinks instead of just reporting.
// The TUI calls reports.Doctor directly with a captured writer.
func RunDoctor(repoRoot string, args DoctorArgs) error {
	if args.Fix {
		_, err := reports.DoctorFix(os.Stdout, repoRoot)
		return err
	}
	return reports.Doctor(os.Stdout, repoRoot)
}
