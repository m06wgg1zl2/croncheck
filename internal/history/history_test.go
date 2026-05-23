package history_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/yourorg/croncheck/internal/history"
)

func tempStore(t *testing.T) *history.Store {
	t.Helper()
	dir := t.TempDir()
	return history.NewStore(filepath.Join(dir, "history.json"))
}

func TestAdd_And_Load(t *testing.T) {
	s := tempStore(t)
	entry := history.Entry{
		Expression: "* * * * *",
		Timezone:   "UTC",
		CheckedAt:  time.Now().UTC(),
		Valid:      true,
	}
	if err := s.Add(entry); err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	entries, err := s.Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
	if entries[0].Expression != "* * * * *" {
		t.Errorf("unexpected expression: %s", entries[0].Expression)
	}
}

func TestLoad_NonExistent(t *testing.T) {
	s := history.NewStore("/tmp/does_not_exist_croncheck/history.json")
	entries, err := s.Load()
	if err != nil {
		t.Fatalf("expected no error for missing file, got: %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("expected empty slice, got %d entries", len(entries))
	}
}

func TestAdd_Multiple(t *testing.T) {
	s := tempStore(t)
	for i := 0; i < 3; i++ {
		_ = s.Add(history.Entry{Expression: "0 * * * *", Valid: true})
	}
	entries, _ := s.Load()
	if len(entries) != 3 {
		t.Errorf("expected 3 entries, got %d", len(entries))
	}
}

func TestClear(t *testing.T) {
	s := tempStore(t)
	_ = s.Add(history.Entry{Expression: "* * * * *", Valid: true})
	if err := s.Clear(); err != nil {
		t.Fatalf("Clear failed: %v", err)
	}
	entries, _ := s.Load()
	if len(entries) != 0 {
		t.Errorf("expected 0 entries after clear, got %d", len(entries))
	}
}

func TestDefaultStorePath(t *testing.T) {
	path, err := history.DefaultStorePath()
	if err != nil {
		t.Fatalf("DefaultStorePath failed: %v", err)
	}
	if path == "" {
		t.Error("expected non-empty path")
	}
	if filepath.Base(path) != "history.json" {
		t.Errorf("unexpected filename: %s", filepath.Base(path))
	}
}

func TestAdd_CreatesDirectory(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sub", "dir", "history.json")
	s := history.NewStore(path)
	if err := s.Add(history.Entry{Expression: "* * * * *", Valid: true}); err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Errorf("expected file to exist: %v", err)
	}
}
