package parser_test

import (
	"testing"

	"github.com/croncheck/croncheck/internal/parser"
)

func TestParse_Valid(t *testing.T) {
	cases := []struct {
		name string
		expr string
	}{
		{"every minute", "* * * * *"},
		{"every hour", "0 * * * *"},
		{"daily at midnight", "0 0 * * *"},
		{"weekdays at noon", "0 12 * * 1-5"},
		{"first of month", "0 9 1 * *"},
		{"every 15 minutes", "*/15 * * * *"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parser.Parse(tc.expr)
			if err != nil {
				t.Fatalf("expected no error for %q, got: %v", tc.expr, err)
			}
			if result.Expression != tc.expr {
				t.Errorf("expected expression %q, got %q", tc.expr, result.Expression)
			}
			if len(result.Fields) != 5 {
				t.Errorf("expected 5 fields, got %d", len(result.Fields))
			}
		})
	}
}

func TestParse_Invalid(t *testing.T) {
	cases := []struct {
		name string
		expr string
	}{
		{"empty string", ""},
		{"too few fields", "* * *"},
		{"too many fields", "* * * * * *"},
		{"bad minute range", "60 * * * *"},
		{"bad hour range", "* 25 * * *"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parser.Parse(tc.expr)
			if err == nil {
				t.Fatalf("expected error for %q, got nil", tc.expr)
			}
		})
	}
}

func TestFieldNames(t *testing.T) {
	names := parser.FieldNames()
	if len(names) != 5 {
		t.Errorf("expected 5 field names, got %d", len(names))
	}
}
