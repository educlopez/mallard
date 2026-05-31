package cmd

import (
	"fmt"
	"os"

	"github.com/educlopez/duck-ai/internal/update"
)

// UpgradeArgs captures parsed flags for `duck-ai upgrade`.
type UpgradeArgs struct {
	Force  bool
	Check  bool
	DryRun bool
}

// ParseUpgradeArgs parses the upgrade subcommand flags.
//
// Supported flags:
//
//	--force    replace even when not strictly newer / on a dev build
//	--check    report the available version without applying anything
//	--dry-run  resolve the action but write nothing to disk
func ParseUpgradeArgs(args []string) (UpgradeArgs, error) {
	var out UpgradeArgs
	for _, a := range args {
		switch a {
		case "--force":
			out.Force = true
		case "--check":
			out.Check = true
		case "--dry-run":
			out.DryRun = true
		default:
			return out, fmt.Errorf("unknown upgrade flag %q", a)
		}
	}
	return out, nil
}

// RunUpgrade self-updates the duck-ai binary to the latest GitHub release.
// currentVersion comes from main's version var.
func RunUpgrade(currentVersion string, args UpgradeArgs) error {
	fmt.Println("\n  duck-ai upgrade")
	return update.Run(os.Stdout, update.Options{
		CurrentVersion: currentVersion,
		Force:          args.Force,
		Check:          args.Check,
		DryRun:         args.DryRun,
	})
}
