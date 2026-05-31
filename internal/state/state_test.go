package state

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// TestRoundTrip verifies a saved state loads back byte-for-byte equal.
func TestRoundTrip(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	want := State{
		LastAgents: []string{"claude", "cursor"},
		LastSkills: []string{"foo", "bar"},
	}
	if err := Save(want); err != nil {
		t.Fatalf("Save: %v", err)
	}

	got := Load()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("round-trip = %+v, want %+v", got, want)
	}

	// The file must live at ~/.mallard/state.json.
	p, _ := Path()
	if want := filepath.Join(home, ".mallard", "state.json"); p != want {
		t.Fatalf("Path() = %q, want %q", p, want)
	}
	if _, err := os.Stat(p); err != nil {
		t.Fatalf("state file not written: %v", err)
	}
}

// TestLoadMissingReturnsDefaults verifies a missing file yields a zero-value
// State without error or panic.
func TestLoadMissingReturnsDefaults(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	got := Load()
	if !reflect.DeepEqual(got, State{}) {
		t.Fatalf("Load(missing) = %+v, want zero value", got)
	}
}

// TestLoadCorruptReturnsDefaults verifies a corrupt JSON file yields defaults
// and never panics.
func TestLoadCorruptReturnsDefaults(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	path := filepath.Join(home, ".mallard", "state.json")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte("{not valid json"), 0o644); err != nil {
		t.Fatalf("write corrupt: %v", err)
	}

	got := Load()
	if !reflect.DeepEqual(got, State{}) {
		t.Fatalf("Load(corrupt) = %+v, want zero value", got)
	}
}
