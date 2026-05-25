package lint_test

import (
	"testing"

	"croncheck/internal/lint"
)

func TestLint_Clean(t *testing.T) {
	warnings, err := lint.Lint("0 9 * * 1-5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(warnings) != 0 {
		t.Errorf("expected no warnings, got %d: %+v", len(warnings), warnings)
	}
}

func TestLint_StepOne(t *testing.T) {
	warnings, err := lint.Lint("*/1 * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	hasW001 := false
	for _, w := range warnings {
		if w.Code == "W001" {
			hasW001 = true
		}
	}
	if !hasW001 {
		t.Error("expected W001 warning for */1")
	}
}

func TestLint_RedundantRange(t *testing.T) {
	warnings, err := lint.Lint("5-5 * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	hasW002 := false
	for _, w := range warnings {
		if w.Code == "W002" {
			hasW002 = true
		}
	}
	if !hasW002 {
		t.Error("expected W002 warning for 5-5")
	}
}

func TestLint_DuplicateListValue(t *testing.T) {
	warnings, err := lint.Lint("1,2,1 * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	hasW003 := false
	for _, w := range warnings {
		if w.Code == "W003" {
			hasW003 = true
		}
	}
	if !hasW003 {
		t.Error("expected W003 warning for duplicate list value")
	}
}

func TestLint_DualDayRestriction(t *testing.T) {
	warnings, err := lint.Lint("0 12 15 * 1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	hasW004 := false
	for _, w := range warnings {
		if w.Code == "W004" {
			hasW004 = true
		}
	}
	if !hasW004 {
		t.Error("expected W004 warning for dual day restriction")
	}
}

func TestLint_InvalidExpression(t *testing.T) {
	_, err := lint.Lint("not a cron")
	if err == nil {
		t.Error("expected error for invalid expression")
	}
}

func TestLint_WarningHasField(t *testing.T) {
	warnings, err := lint.Lint("*/1 * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(warnings) == 0 {
		t.Fatal("expected at least one warning")
	}
	if warnings[0].Field == "" {
		t.Error("warning Field should not be empty")
	}
}
