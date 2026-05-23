package scheduler

import (
	"errors"
	"testing"
	"time"

	"github.com/croncheck/internal/parser"
)

func mustParse(t *testing.T, expr string) *parser.Expression {
	t.Helper()
	e, err := parser.Parse(expr)
	if err != nil {
		t.Fatalf("failed to parse %q: %v", expr, err)
	}
	return e
}

func TestNextRuns_EveryMinute(t *testing.T) {
	expr := mustParse(t, "* * * * *")
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	runs, err := NextRuns(expr, start, 3, time.UTC)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(runs) != 3 {
		t.Fatalf("expected 3 runs, got %d", len(runs))
	}
	expected := []time.Time{
		time.Date(2024, 1, 1, 0, 1, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 2, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 3, 0, 0, time.UTC),
	}
	for i, r := range runs {
		if !r.Equal(expected[i]) {
			t.Errorf("run[%d]: expected %v, got %v", i, expected[i], r)
		}
	}
}

func TestNextRuns_HourlyAt30(t *testing.T) {
	expr := mustParse(t, "30 * * * *")
	start := time.Date(2024, 6, 15, 10, 0, 0, 0, time.UTC)
	runs, err := NextRuns(expr, start, 2, time.UTC)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(runs) != 2 {
		t.Fatalf("expected 2 runs, got %d", len(runs))
	}
	if runs[0].Minute() != 30 || runs[0].Hour() != 10 {
		t.Errorf("unexpected first run: %v", runs[0])
	}
	if runs[1].Minute() != 30 || runs[1].Hour() != 11 {
		t.Errorf("unexpected second run: %v", runs[1])
	}
}

func TestNextRuns_NilExpression(t *testing.T) {
	_, err := NextRuns(nil, time.Now(), 5, time.UTC)
	if !errors.Is(err, ErrNilExpression) {
		t.Errorf("expected ErrNilExpression, got %v", err)
	}
}

func TestNextRuns_ZeroCount(t *testing.T) {
	expr := mustParse(t, "* * * * *")
	runs, err := NextRuns(expr, time.Now(), 0, time.UTC)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(runs) != 0 {
		t.Errorf("expected 0 runs, got %d", len(runs))
	}
}

func TestNextRuns_Timezone(t *testing.T) {
	expr := mustParse(t, "0 9 * * *")
	loc, _ := time.LoadLocation("America/New_York")
	start := time.Date(2024, 3, 1, 0, 0, 0, 0, loc)
	runs, err := NextRuns(expr, start, 1, loc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(runs) != 1 {
		t.Fatalf("expected 1 run, got %d", len(runs))
	}
	if runs[0].Hour() != 9 || runs[0].Minute() != 0 {
		t.Errorf("expected 09:00, got %v", runs[0])
	}
}
