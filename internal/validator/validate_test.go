package validator

import (
	"testing"
)

func TestValidate_Valid(t *testing.T) {
	cases := []string{
		"* * * * *",
		"0 * * * *",
		"30 6 * * 1",
		"*/5 * * * *",
		"0 0 1 1 *",
		"0-30 8-17 * * 1-5",
		"1,15,30 * * * *",
		"0 12 */2 * *",
	}
	for _, expr := range cases {
		t.Run(expr, func(t *testing.T) {
			errs := Validate(expr)
			if len(errs) != 0 {
				t.Errorf("expected no errors for %q, got %v", expr, errs)
			}
		})
	}
}

func TestValidate_WrongFieldCount(t *testing.T) {
	cases := []string{"* * *", "* * * *", "* * * * * *", ""}
	for _, expr := range cases {
		errs := Validate(expr)
		if len(errs) == 0 {
			t.Errorf("expected error for %q", expr)
		}
		if errs[0].Field != "expression" {
			t.Errorf("expected field=expression, got %q", errs[0].Field)
		}
	}
}

func TestValidate_OutOfRange(t *testing.T) {
	cases := []struct {
		expr  string
		field string
	}{
		{"60 * * * *", "minute"},
		{"* 24 * * *", "hour"},
		{"* * 32 * *", "day-of-month"},
		{"* * * 13 *", "month"},
		{"* * * * 7", "day-of-week"},
	}
	for _, tc := range cases {
		errs := Validate(tc.expr)
		if len(errs) == 0 {
			t.Errorf("expected error for %q", tc.expr)
			continue
		}
		if errs[0].Field != tc.field {
			t.Errorf("expr %q: expected field=%q, got %q", tc.expr, tc.field, errs[0].Field)
		}
	}
}

func TestValidate_InvalidStep(t *testing.T) {
	errs := Validate("*/0 * * * *")
	if len(errs) == 0 {
		t.Error("expected error for step=0")
	}
}

func TestValidate_InvalidRange(t *testing.T) {
	errs := Validate("10-5 * * * *")
	if len(errs) == 0 {
		t.Error("expected error for inverted range 10-5")
	}
}

func TestValidationError_Error(t *testing.T) {
	e := &ValidationError{Field: "minute", Reason: "value 99 out of bounds [0-59]"}
	got := e.Error()
	if got != "invalid minute: value 99 out of bounds [0-59]" {
		t.Errorf("unexpected error string: %q", got)
	}
}
