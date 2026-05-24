package scheduler

import (
	"testing"
	"time"
)

func TestPrevRuns_EveryMinute(t *testing.T) {
	expr := mustParse(t, "* * * * *")
	ref := time.Date(2024, 6, 1, 12, 10, 0, 0, time.UTC)

	runs := PrevRuns(expr, ref, 5)
	if len(runs) != 5 {
		t.Fatalf("expected 5 runs, got %d", len(runs))
	}

	// Each run should be exactly one minute before the previous.
	for i := 1; i < len(runs); i++ {
		diff := runs[i-1].Sub(runs[i])
		if diff != time.Minute {
			t.Errorf("expected 1m gap between run %d and %d, got %v", i-1, i, diff)
		}
	}
}

func TestPrevRuns_AllBeforeRef(t *testing.T) {
	expr := mustParse(t, "* * * * *")
	ref := time.Date(2024, 6, 1, 12, 10, 30, 0, time.UTC)

	runs := PrevRuns(expr, ref, 3)
	for _, r := range runs {
		if !r.Before(ref) {
			t.Errorf("expected run %v to be before ref %v", r, ref)
		}
	}
}

func TestPrevRuns_HourlyAt30(t *testing.T) {
	expr := mustParse(t, "30 * * * *")
	ref := time.Date(2024, 6, 1, 15, 45, 0, 0, time.UTC)

	runs := PrevRuns(expr, ref, 3)
	if len(runs) != 3 {
		t.Fatalf("expected 3 runs, got %d", len(runs))
	}

	for _, r := range runs {
		if r.Minute() != 30 {
			t.Errorf("expected minute 30, got %d", r.Minute())
		}
	}
}

func TestPrevRuns_NilExpression(t *testing.T) {
	runs := PrevRuns(nil, time.Now(), 5)
	if runs != nil {
		t.Errorf("expected nil for nil expression, got %v", runs)
	}
}

func TestPrevRuns_ZeroCount(t *testing.T) {
	expr := mustParse(t, "* * * * *")
	runs := PrevRuns(expr, time.Now(), 0)
	if runs != nil {
		t.Errorf("expected nil for zero count, got %v", runs)
	}
}

func TestPrevRuns_DescendingOrder(t *testing.T) {
	expr := mustParse(t, "0 * * * *")
	ref := time.Date(2024, 6, 1, 18, 0, 0, 0, time.UTC)

	runs := PrevRuns(expr, ref, 4)
	for i := 1; i < len(runs); i++ {
		if !runs[i].Before(runs[i-1]) {
			t.Errorf("runs not in descending order at index %d: %v >= %v", i, runs[i], runs[i-1])
		}
	}
}
