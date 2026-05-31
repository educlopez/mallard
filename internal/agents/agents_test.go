package agents

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestAgentsDir(t *testing.T) {
	tests := []struct {
		id         string
		wantSuffix string // non-empty: AgentsDir must end with this
		wantEmpty  bool   // true: AgentsDir must be ""
	}{
		{id: IDClaude, wantSuffix: filepath.Join(".claude", "agents")},
		{id: IDCodex, wantEmpty: true},
		{id: IDOpenCode, wantEmpty: true},
		{id: IDAgents, wantEmpty: true},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			a, ok := ByID(tt.id)
			if !ok {
				t.Fatalf("ByID(%q) not found", tt.id)
			}
			got := a.AgentsDir()
			if tt.wantEmpty {
				if got != "" {
					t.Fatalf("%s AgentsDir() = %q, want empty", tt.id, got)
				}
				return
			}
			if !strings.HasSuffix(got, tt.wantSuffix) {
				t.Fatalf("%s AgentsDir() = %q, want suffix %q", tt.id, got, tt.wantSuffix)
			}
		})
	}
}

func TestByID(t *testing.T) {
	if _, ok := ByID("does-not-exist"); ok {
		t.Fatalf("ByID(unknown) returned ok=true")
	}
	for _, id := range []string{IDClaude, IDCodex, IDOpenCode, IDAgents} {
		a, ok := ByID(id)
		if !ok {
			t.Fatalf("ByID(%q) not found", id)
		}
		if a.ID() != id {
			t.Fatalf("adapter ID() = %q, want %q", a.ID(), id)
		}
	}
}

func TestAllReturnsCopy(t *testing.T) {
	got := All()
	if len(got) != len(all) {
		t.Fatalf("All() len = %d, want %d", len(got), len(all))
	}
	// Mutating the returned slice must not affect the package-level registry.
	if len(got) > 0 {
		got[0] = nil
		if All()[0] == nil {
			t.Fatalf("All() returned a slice aliasing the internal registry")
		}
	}
}
