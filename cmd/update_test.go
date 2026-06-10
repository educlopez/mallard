package cmd

import (
	"strings"
	"testing"

	"github.com/educlopez/mallard/internal/agents"
)

func TestParseUpdateArgs(t *testing.T) {
	t.Run("dry-run flag", func(t *testing.T) {
		got, err := ParseUpdateArgs([]string{"--dry-run"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.DryRun {
			t.Fatalf("DryRun = false, want true")
		}
	})

	t.Run("yes flag", func(t *testing.T) {
		got, err := ParseUpdateArgs([]string{"--yes"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.Yes {
			t.Fatalf("Yes = false, want true")
		}
	})

	t.Run("scope global", func(t *testing.T) {
		got, err := ParseUpdateArgs([]string{"--scope", "global"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Scope != agents.ScopeGlobal {
			t.Fatalf("Scope = %q, want global", got.Scope)
		}
	})

	t.Run("scope workspace", func(t *testing.T) {
		got, err := ParseUpdateArgs([]string{"--scope", "workspace"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Scope != agents.ScopeWorkspace {
			t.Fatalf("Scope = %q, want workspace", got.Scope)
		}
	})

	t.Run("scope requires value", func(t *testing.T) {
		_, err := ParseUpdateArgs([]string{"--scope"})
		if err == nil {
			t.Fatal("expected error for missing --scope value, got nil")
		}
	})

	t.Run("restore flag", func(t *testing.T) {
		got, err := ParseUpdateArgs([]string{"--restore", "20240101T120000"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Restore != "20240101T120000" {
			t.Fatalf("Restore = %q, want 20240101T120000", got.Restore)
		}
	})

	t.Run("list-backups flag", func(t *testing.T) {
		got, err := ParseUpdateArgs([]string{"--list-backups"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.ListBackups {
			t.Fatalf("ListBackups = false, want true")
		}
	})

	t.Run("restore and list-backups are mutually exclusive", func(t *testing.T) {
		_, err := ParseUpdateArgs([]string{"--restore", "foo", "--list-backups"})
		if err == nil {
			t.Fatal("expected error for mutual exclusion, got nil")
		}
	})

	t.Run("unknown flag returns error", func(t *testing.T) {
		_, err := ParseUpdateArgs([]string{"--bogus"})
		if err == nil {
			t.Fatal("expected error for unknown flag, got nil")
		}
		if !strings.Contains(err.Error(), "--bogus") {
			t.Fatalf("error message %q does not mention --bogus", err.Error())
		}
	})

	t.Run("unknown short flag returns error", func(t *testing.T) {
		_, err := ParseUpdateArgs([]string{"-z"})
		if err == nil {
			t.Fatal("expected error for unknown short flag, got nil")
		}
	})

	t.Run("non-flag arguments are ignored", func(t *testing.T) {
		_, err := ParseUpdateArgs([]string{"positional"})
		if err != nil {
			t.Fatalf("unexpected error for positional arg: %v", err)
		}
	})
}
