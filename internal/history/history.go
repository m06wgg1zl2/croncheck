package history

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Entry represents a single history record of a validated cron expression.
type Entry struct {
	Expression string    `json:"expression"`
	Timezone   string    `json:"timezone"`
	CheckedAt  time.Time `json:"checked_at"`
	Valid      bool      `json:"valid"`
}

// Store manages persistence of cron check history.
type Store struct {
	path string
}

// NewStore creates a Store backed by the given file path.
func NewStore(path string) *Store {
	return &Store{path: path}
}

// DefaultStorePath returns the default path for the history file.
func DefaultStorePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".croncheck", "history.json"), nil
}

// Add appends a new entry to the history file, creating it if necessary.
func (s *Store) Add(entry Entry) error {
	entries, err := s.Load()
	if err != nil {
		entries = []Entry{}
	}
	entries = append(entries, entry)
	return s.save(entries)
}

// Load reads all entries from the history file.
func (s *Store) Load() ([]Entry, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}
		return nil, err
	}
	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}

// Clear removes all entries from the history file.
func (s *Store) Clear() error {
	return s.save([]Entry{})
}

func (s *Store) save(entries []Entry) error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0o644)
}
