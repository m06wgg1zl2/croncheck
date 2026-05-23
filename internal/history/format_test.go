package history_test

import (
	"strings"
	"testing"
	"time"

	"github.com/yourorg/croncheck/internal/history"
)

func sampleEntries() []history.Entry {
	base := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	return []history.Entry{
		{Expression: "* * * * *", Timezone: "UTC", CheckedAt: base, Valid: true},
		{Expression: "0 9 * * 1", Timezone: "America/New_York", CheckedAt: base.Add(time.Hour), Valid: false},
	}
}

func TestFormatTable_ContainsHeader(t *testing.T) {
	out := history.FormatTable(sampleEntries(), time.UTC)
	if !strings.Contains(out, "Expression") {
		t.Error("expected header to contain 'Expression'")
	}
	if !strings.Contains(out, "Status") {
		t.Error("expected header to contain 'Status'")
	}
}

func TestFormatTable_ContainsExpressions(t *testing.T) {
	out := history.FormatTable(sampleEntries(), time.UTC)
	if !strings.Contains(out, "* * * * *") {
		t.Error("expected output to contain first expression")
	}
	if !strings.Contains(out, "0 9 * * 1") {
		t.Error("expected output to contain second expression")
	}
}

func TestFormatTable_ValidStatus(t *testing.T) {
	out := history.FormatTable(sampleEntries(), time.UTC)
	if !strings.Contains(out, "✓ valid") {
		t.Error("expected valid status symbol")
	}
	if !strings.Contains(out, "✗ invalid") {
		t.Error("expected invalid status symbol")
	}
}

func TestFormatTable_Empty(t *testing.T) {
	out := history.FormatTable([]history.Entry{}, time.UTC)
	if !strings.Contains(out, "No history") {
		t.Errorf("expected empty message, got: %s", out)
	}
}

func TestFormatTable_NilTimezone(t *testing.T) {
	out := history.FormatTable(sampleEntries(), nil)
	if out == "" {
		t.Error("expected non-empty output with nil timezone")
	}
}

func TestSummary_Mixed(t *testing.T) {
	s := history.Summary(sampleEntries())
	if !strings.Contains(s, "2 checks") {
		t.Errorf("expected total count, got: %s", s)
	}
	if !strings.Contains(s, "1 valid") {
		t.Errorf("expected valid count, got: %s", s)
	}
	if !strings.Contains(s, "1 invalid") {
		t.Errorf("expected invalid count, got: %s", s)
	}
}

func TestSummary_Empty(t *testing.T) {
	s := history.Summary([]history.Entry{})
	if !strings.Contains(s, "No history") {
		t.Errorf("expected empty message, got: %s", s)
	}
}
