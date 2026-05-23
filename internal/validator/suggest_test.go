package validator

import (
	"strings"
	"testing"
)

func TestSuggest_Length(t *testing.T) {
	suggestions := Suggest()
	if len(suggestions) == 0 {
		t.Fatal("expected at least one suggestion")
	}
}

func TestSuggest_AllValid(t *testing.T) {
	for _, s := range Suggest() {
		errs := Validate(s.Expr)
		if len(errs) != 0 {
			t.Errorf("suggestion %q has invalid expr %q: %v", s.Label, s.Expr, errs)
		}
	}
}

func TestSuggest_NoDuplicateLabels(t *testing.T) {
	seen := map[string]bool{}
	for _, s := range Suggest() {
		if seen[s.Label] {
			t.Errorf("duplicate label: %q", s.Label)
		}
		seen[s.Label] = true
	}
}

func TestSuggest_EveryMinuteFirst(t *testing.T) {
	suggestions := Suggest()
	if suggestions[0].Label != "every minute" {
		t.Errorf("expected first suggestion to be 'every minute', got %q", suggestions[0].Label)
	}
}

func TestFormatSuggestions_NotEmpty(t *testing.T) {
	out := FormatSuggestions()
	if strings.TrimSpace(out) == "" {
		t.Error("expected non-empty formatted suggestions")
	}
}

func TestFormatSuggestions_ContainsExpr(t *testing.T) {
	out := FormatSuggestions()
	for _, s := range Suggest() {
		if !strings.Contains(out, s.Expr) {
			t.Errorf("formatted output missing expression %q", s.Expr)
		}
	}
}

func TestFormatSuggestions_ContainsLabel(t *testing.T) {
	out := FormatSuggestions()
	if !strings.Contains(out, "every minute") {
		t.Error("formatted output missing label 'every minute'")
	}
}
