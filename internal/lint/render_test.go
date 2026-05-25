package lint_test

import (
	"strings"
	"testing"

	"croncheck/internal/lint"
)

func mustLint(t *testing.T, expr string) []lint.Warning {
	t.Helper()
	w, err := lint.Lint(expr)
	if err != nil {
		t.Fatalf("lint error: %v", err)
	}
	return w
}

func TestRender_NoWarnings(t *testing.T) {
	w := mustLint(t, "0 9 * * 1-5")
	out := lint.Render("0 9 * * 1-5", w)
	if !strings.Contains(out, "No lint warnings") {
		t.Errorf("expected no-warning message, got: %s", out)
	}
}

func TestRender_ContainsExpression(t *testing.T) {
	w := mustLint(t, "*/1 * * * *")
	out := lint.Render("*/1 * * * *", w)
	if !strings.Contains(out, "*/1 * * * *") {
		t.Errorf("expected expression in output, got: %s", out)
	}
}

func TestRender_ContainsCode(t *testing.T) {
	w := mustLint(t, "*/1 * * * *")
	out := lint.Render("*/1 * * * *", w)
	if !strings.Contains(out, "W001") {
		t.Errorf("expected warning code W001 in output, got: %s", out)
	}
}

func TestRender_ContainsCount(t *testing.T) {
	w := mustLint(t, "*/1 * * * *")
	out := lint.Render("*/1 * * * *", w)
	if !strings.Contains(out, "warning(s)") {
		t.Errorf("expected warning count in output, got: %s", out)
	}
}

func TestRender_ContainsField(t *testing.T) {
	w := mustLint(t, "*/1 * * * *")
	out := lint.Render("*/1 * * * *", w)
	if !strings.Contains(out, "minute") {
		t.Errorf("expected field name in output, got: %s", out)
	}
}

func TestRender_MultipleWarnings(t *testing.T) {
	w := mustLint(t, "*/1 * 15 * 1")
	out := lint.Render("*/1 * 15 * 1", w)
	if !strings.Contains(out, "W001") || !strings.Contains(out, "W004") {
		t.Errorf("expected multiple warning codes, got: %s", out)
	}
}
