package parser_test

import (
	"strings"
	"testing"

	"github.com/croncheck/croncheck/internal/parser"
)

func TestDescribe(t *testing.T) {
	cases := []struct {
		expr     string
		contains string
	}{
		{"* * * * *", "every"},
		{"0 12 * * *", "minute=0"},
		{"*/15 * * * *", "every 15"},
		{"0 9-17 * * 1-5", "through"},
		{"0 0 1,15 * *", "[1,15]"},
	}

	for _, tc := range cases {
		t.Run(tc.expr, func(t *testing.T) {
			result, err := parser.Parse(tc.expr)
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}
			desc := parser.Describe(result)
			if !strings.Contains(desc, tc.contains) {
				t.Errorf("expected description to contain %q, got: %q", tc.contains, desc)
			}
		})
	}
}

func TestDescribe_Nil(t *testing.T) {
	result := parser.Describe(nil)
	if result != "invalid expression" {
		t.Errorf("expected 'invalid expression' for nil input, got %q", result)
	}
}
