package diff_test

import (
	"testing"

	"github.com/user/croncheck/internal/diff"
)

func TestCompare_SameExpression(t *testing.T) {
	r, err := diff.Compare("* * * * *", "* * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(r.Changed()) != 0 {
		t.Errorf("expected no changes, got %d", len(r.Changed()))
	}
}

func TestCompare_DifferentMinute(t *testing.T) {
	r, err := diff.Compare("0 * * * *", "30 * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	changed := r.Changed()
	if len(changed) != 1 {
		t.Fatalf("expected 1 changed field, got %d", len(changed))
	}
	if changed[0].Name != "minute" {
		t.Errorf("expected minute to differ, got %s", changed[0].Name)
	}
	if changed[0].Left != "0" || changed[0].Right != "30" {
		t.Errorf("unexpected values: %s vs %s", changed[0].Left, changed[0].Right)
	}
}

func TestCompare_MultipleFieldsDiffer(t *testing.T) {
	r, err := diff.Compare("0 9 * * 1", "30 18 * * 5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(r.Changed()) != 3 {
		t.Errorf("expected 3 changed fields, got %d", len(r.Changed()))
	}
}

func TestCompare_InvalidLeft(t *testing.T) {
	_, err := diff.Compare("bad expr", "* * * * *")
	if err == nil {
		t.Error("expected error for invalid left expression")
	}
}

func TestCompare_InvalidRight(t *testing.T) {
	_, err := diff.Compare("* * * * *", "99 * * * *")
	if err == nil {
		t.Error("expected error for invalid right expression")
	}
}

func TestResult_Fields_AllPresent(t *testing.T) {
	r, err := diff.Compare("* * * * *", "* * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(r.Fields) != 5 {
		t.Errorf("expected 5 fields, got %d", len(r.Fields))
	}
}
