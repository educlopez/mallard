package reports

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/educlopez/mallard/internal/agents"
	"github.com/educlopez/mallard/internal/skillregistry"
)

// RegistryArgs mirrors cmd.RegistryArgs so callers can pass flag state into
// the renderer. cmd.RunRegistry adapts cmd.RegistryArgs to this struct.
type RegistryArgs struct {
	Source bool
	JSON   bool
	All    bool
	Help   bool
}

// ParseRegistryArgs parses the same flags accepted by `mallard registry`.
// Kept here so the TUI can also build args without depending on cmd.
func ParseRegistryArgs(args []string) (RegistryArgs, error) {
	var out RegistryArgs
	for _, a := range args {
		switch a {
		case "--source":
			out.Source = true
		case "--json":
			out.JSON = true
		case "--all":
			out.All = true
		case "--help", "-h":
			out.Help = true
		default:
			if strings.HasPrefix(a, "-") {
				return out, fmt.Errorf("unknown flag %q", a)
			}
		}
	}
	return out, nil
}

// Registry writes the mallard registry report to w.
func Registry(w io.Writer, repoRoot string, args RegistryArgs) error {
	if args.Help {
		PrintRegistryHelp(w)
		return nil
	}

	source, err := skillregistry.ParseSource(repoRoot)
	if err != nil {
		return err
	}

	if args.Source && !args.JSON {
		return printSourceText(w, source)
	}

	installed := map[string][]skillregistry.Manifest{}
	for _, a := range agents.All() {
		if !a.Detect() {
			continue
		}
		ms, err := skillregistry.ParseInstalled(a)
		if err != nil {
			return err
		}
		installed[a.ID()] = ms
	}

	sourceVersions := map[string]string{}
	for _, m := range source {
		sourceVersions[m.Kind+"/"+m.Name] = m.Version
	}

	// Default behavior: filter out orphan/unversioned entries so only
	// mallard-managed entries are shown. --all disables the filter.
	if !args.All {
		filtered := map[string][]skillregistry.Manifest{}
		for id, ms := range installed {
			kept := make([]skillregistry.Manifest, 0, len(ms))
			for _, m := range ms {
				if isManaged(m, sourceVersions) {
					kept = append(kept, m)
				}
			}
			filtered[id] = kept
		}
		installed = filtered
	}

	if args.JSON {
		payload := map[string]any{
			"source":    source,
			"installed": installed,
		}
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		return enc.Encode(payload)
	}

	return printInstalledText(w, sourceVersions, installed)
}

func printSourceText(w io.Writer, source []skillregistry.Manifest) error {
	fmt.Fprintln(w, "\n  mallard registry — source")
	skills, commands, agentDefs := splitByKind(source)
	if len(skills) > 0 {
		fmt.Fprintln(w, "    skills:")
		for _, m := range skills {
			fmt.Fprintf(w, "      %-32s %s\n", m.Name, versionLabel(m.Version))
		}
	}
	if len(commands) > 0 {
		fmt.Fprintln(w, "    commands:")
		for _, m := range commands {
			fmt.Fprintf(w, "      %-32s %s\n", m.Name, versionLabel(m.Version))
		}
	}
	if len(agentDefs) > 0 {
		fmt.Fprintln(w, "    agents:")
		for _, m := range agentDefs {
			fmt.Fprintf(w, "      %-32s %s\n", m.Name, versionLabel(m.Version))
		}
	}
	return nil
}

func printInstalledText(w io.Writer, sourceVersions map[string]string, installed map[string][]skillregistry.Manifest) error {
	fmt.Fprintln(w, "\nmallard registry")

	if len(installed) == 0 {
		fmt.Fprintln(w, "  No agents detected.")
		return nil
	}

	for _, a := range agents.All() {
		ms, ok := installed[a.ID()]
		if !ok {
			continue
		}
		fmt.Fprintf(w, "\n  Agent: %s\n", a.ID())
		skills, commands, agentDefs := splitByKind(ms)
		if len(skills) == 0 && len(commands) == 0 && len(agentDefs) == 0 {
			fmt.Fprintln(w, "    (none managed)")
			continue
		}
		if len(skills) > 0 {
			fmt.Fprintln(w, "    skills:")
			for _, m := range skills {
				fmt.Fprintf(w, "      %-32s %-8s %s\n",
					m.Name, versionLabel(m.Version), statusFor(m, sourceVersions))
			}
		}
		if len(commands) > 0 {
			fmt.Fprintln(w, "    commands:")
			for _, m := range commands {
				fmt.Fprintf(w, "      %-32s %-8s %s\n",
					m.Name, versionLabel(m.Version), statusFor(m, sourceVersions))
			}
		}
		if len(agentDefs) > 0 {
			fmt.Fprintln(w, "    agents:")
			for _, m := range agentDefs {
				fmt.Fprintf(w, "      %-32s %-8s %s\n",
					m.Name, versionLabel(m.Version), statusFor(m, sourceVersions))
			}
		}
	}
	return nil
}

func splitByKind(ms []skillregistry.Manifest) (skills, commands, agentDefs []skillregistry.Manifest) {
	for _, m := range ms {
		switch m.Kind {
		case "skill":
			skills = append(skills, m)
		case "command":
			commands = append(commands, m)
		case "agent":
			agentDefs = append(agentDefs, m)
		}
	}
	return
}

func versionLabel(v string) string {
	if v == "" {
		return "(no ver)"
	}
	return "v" + v
}

// isManaged reports whether an installed manifest corresponds to a mallard
// source entry (matched by kind + name).
func isManaged(m skillregistry.Manifest, sourceVersions map[string]string) bool {
	if m.Version == "" {
		return false
	}
	_, ok := sourceVersions[m.Kind+"/"+m.Name]
	return ok
}

func statusFor(m skillregistry.Manifest, sourceVersions map[string]string) string {
	if m.Version == "" {
		return "unversioned"
	}
	srcVer, ok := sourceVersions[m.Kind+"/"+m.Name]
	if !ok {
		return "orphan"
	}
	if srcVer == "" {
		return "unversioned"
	}
	if srcVer != m.Version {
		return "drift"
	}
	return "ok"
}

// PrintRegistryHelp writes the registry --help text to w.
func PrintRegistryHelp(w io.Writer) {
	fmt.Fprint(w, `mallard registry — list installed skills/commands per agent

Usage:
  mallard registry             Show only mallard-managed entries (default)
  mallard registry --all       Show every entry, including orphans and
                               unversioned items from other tooling
  mallard registry --source    List source entries from the mallard repo
  mallard registry --json      Emit machine-readable JSON (respects --all)
  mallard registry --help      Show this help

By default, only entries that match a mallard source skill/command by name
are shown. Use --all to include orphan and unversioned entries written by
other tooling.
`)
}
