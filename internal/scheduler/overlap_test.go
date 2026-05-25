package scheduler

import (
	"testing"
	"time"
)

func TestFindOverlaps_SameExpression(t *testing.T) {
	expr := mustParse("*/5 * * * *")
	from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC)

	result, err := FindOverlaps(expr, expr, from, to)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.TotalA != result.TotalB {
		t.Errorf("expected TotalA == TotalB, got %d vs %d", result.TotalA, result.TotalB)
	}
	if len(result.Overlaps) != result.TotalA {
		t.Errorf("expected all runs to overlap, got %d overlaps out of %d", len(result.Overlaps), result.TotalA)
	}
}

func TestFindOverlaps_NoOverlap(t *testing.T) {
	// every even minute vs every odd minute — no overlap
	exprA := mustParse("*/2 * * * *")
	exprB := mustParse("1-59/2 * * * *")
	from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 1, 1, 0, 10, 0, 0, time.UTC)

	result, err := FindOverlaps(exprA, exprB, from, to)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Overlaps) != 0 {
		t.Errorf("expected 0 overlaps, got %d", len(result.Overlaps))
	}
}

func TestFindOverlaps_PartialOverlap(t *testing.T) {
	exprA := mustParse("*/15 * * * *") // 0,15,30,45
	exprB := mustParse("*/30 * * * *") // 0,30
	from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC)

	result, err := FindOverlaps(exprA, exprB, from, to)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Expect overlaps at :00 and :30
	if len(result.Overlaps) != 2 {
		t.Errorf("expected 2 overlaps, got %d", len(result.Overlaps))
	}
}

func TestFindOverlaps_NilExpression(t *testing.T) {
	expr := mustParse("* * * * *")
	from := time.Now()
	to := from.Add(time.Hour)

	result, err := FindOverlaps(nil, expr, from, to)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != nil {
		t.Error("expected nil result when expression is nil")
	}
}

func TestFindOverlaps_ExprFieldsPopulated(t *testing.T) {
	exprA := mustParse("0 * * * *")
	exprB := mustParse("30 * * * *")
	from := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 6, 1, 2, 0, 0, 0, time.UTC)

	result, err := FindOverlaps(exprA, exprB, from, to)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ExprA != exprA.Raw {
		t.Errorf("expected ExprA %q, got %q", exprA.Raw, result.ExprA)
	}
	if result.ExprB != exprB.Raw {
		t.Errorf("expected ExprB %q, got %q", exprB.Raw, result.ExprB)
	}
	if len(result.Overlaps) != 0 {
		t.Errorf("expected 0 overlaps for non-overlapping expressions, got %d", len(result.Overlaps))
	}
}
