package scheduler

import (
	"testing"
	"time"

	"github.com/user/croncheck/internal/parser"
)

func mustParseExpr(t *testing.T, expr string) *parser.Expression {
	t.Helper()
	e, err := parser.Parse(expr)
	if err != nil {
		t.Fatalf("mustParseExpr(%q): %v", expr, err)
	}
	return e
}

func TestRunsInWindow_EveryMinute(t *testing.T) {
	expr := mustParseExpr(t, "* * * * *")
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	runs, err := WindowDuration(expr, start, 5*time.Minute)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(runs) != 5 {
		t.Errorf("expected 5 runs, got %d", len(runs))
	}
}

func TestRunsInWindow_HourlyAt30(t *testing.T) {
	expr := mustParseExpr(t, "30 * * * *")
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	runs, err := WindowDuration(expr, start, 2*time.Hour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Expect runs at 00:30 and 01:30
	if len(runs) != 2 {
		t.Errorf("expected 2 runs, got %d", len(runs))
	}
	for _, r := range runs {
		if r.Minute() != 30 {
			t.Errorf("expected minute 30, got %d", r.Minute())
		}
	}
}

func TestRunsInWindow_NilExpression(t *testing.T) {
	_, err := RunsInWindow(nil, Window{
		Start: time.Now(),
		End:   time.Now().Add(time.Hour),
	})
	if err == nil {
		t.Error("expected error for nil expression, got nil")
	}
}

func TestRunsInWindow_EndBeforeStart(t *testing.T) {
	expr := mustParseExpr(t, "* * * * *")
	now := time.Now()
	runs, err := RunsInWindow(expr, Window{Start: now, End: now.Add(-time.Minute)})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(runs) != 0 {
		t.Errorf("expected 0 runs for inverted window, got %d", len(runs))
	}
}

func TestRunsInWindow_AllWithinBounds(t *testing.T) {
	expr := mustParseExpr(t, "* * * * *")
	start := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	end := start.Add(10 * time.Minute)
	runs, err := RunsInWindow(expr, Window{Start: start, End: end})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, r := range runs {
		if r.Before(start) || !r.Before(end) {
			t.Errorf("run %v is outside window [%v, %v)", r, start, end)
		}
	}
}
