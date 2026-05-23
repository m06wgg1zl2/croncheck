package formatter

import (
	"strings"
	"testing"
	"time"
)

func fixedRuns(base time.Time, count int, step time.Duration) []time.Time {
	runs := make([]time.Time, count)
	for i := range runs {
		runs[i] = base.Add(step * time.Duration(i+1))
	}
	return runs
}

func TestBuildTable_Length(t *testing.T) {
	base := time.Now().UTC()
	runs := fixedRuns(base, 5, time.Minute)
	rows := BuildTable(runs, time.UTC)
	if len(rows) != 5 {
		t.Fatalf("expected 5 rows, got %d", len(rows))
	}
}

func TestBuildTable_IndexStartsAtOne(t *testing.T) {
	base := time.Now().UTC()
	runs := fixedRuns(base, 3, time.Minute)
	rows := BuildTable(runs, time.UTC)
	if rows[0].Index != 1 {
		t.Errorf("expected first index 1, got %d", rows[0].Index)
	}
}

func TestBuildTable_NilTimezone(t *testing.T) {
	base := time.Now().UTC()
	runs := fixedRuns(base, 2, time.Hour)
	// should not panic with nil tz
	rows := BuildTable(runs, nil)
	if len(rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(rows))
	}
}

func TestRenderTable_ContainsHeader(t *testing.T) {
	base := time.Now().UTC()
	runs := fixedRuns(base, 3, time.Minute)
	rows := BuildTable(runs, time.UTC)
	out := RenderTable(rows, "UTC")
	if !strings.Contains(out, "UTC") {
		t.Error("expected table to contain timezone header 'UTC'")
	}
	if !strings.Contains(out, "#") {
		t.Error("expected table to contain '#' column")
	}
}

func TestRenderTable_EmptyRows(t *testing.T) {
	out := RenderTable([]TableRow{}, "UTC")
	if out != "No upcoming runs." {
		t.Errorf("unexpected output for empty rows: %q", out)
	}
}

func TestRenderTable_RowCount(t *testing.T) {
	base := time.Now().UTC()
	runs := fixedRuns(base, 4, time.Hour)
	rows := BuildTable(runs, time.UTC)
	out := RenderTable(rows, "UTC")
	// count data lines (non-separator, non-header)
	lines := strings.Split(strings.TrimSpace(out), "\n")
	// separator + header + separator + 4 data + separator = 8 lines
	if len(lines) != 8 {
		t.Errorf("expected 8 lines in table output, got %d: %v", len(lines), lines)
	}
}

func TestRelativeTime_Future(t *testing.T) {
	now := time.Now().UTC()
	future := now.Add(90 * time.Second)
	r := relativeTime(now, future)
	if !strings.HasPrefix(r, "in") {
		t.Errorf("expected relative time to start with 'in', got %q", r)
	}
}

func TestRelativeTime_Past(t *testing.T) {
	now := time.Now().UTC()
	past := now.Add(-5 * time.Minute)
	r := relativeTime(now, past)
	if r != "in the past" {
		t.Errorf("expected 'in the past', got %q", r)
	}
}
