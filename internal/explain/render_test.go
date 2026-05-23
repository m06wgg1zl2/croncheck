package explain_test

import (
	"strings"
	"testing"

	"croncheck/internal/explain"
)

func mustExplain(t *testing.T, expr string) *explain.Explanation {
	t.Helper()
	e, err := explain.Explain(expr)
	if err != nil {
		t.Fatalf("mustExplain: %v", err)
	}
	return e
}

func TestRender_ContainsExpression(t *testing.T) {
	e := mustExplain(t, "0 12 * * *")
	out := explain.Render(e, false)
	if !strings.Contains(out, "0 12 * * *") {
		t.Errorf("expected expression in output, got:\n%s", out)
	}
}

func TestRender_ContainsFieldNames(t *testing.T) {
	e := mustExplain(t, "* * * * *")
	out := explain.Render(e, false)
	for _, name := range []string{"Minute", "Hour", "Day", "Month", "Weekday"} {
		if !strings.Contains(out, name) {
			t.Errorf("expected field name %q in output", name)
		}
	}
}

func TestRender_ContainsFrequency(t *testing.T) {
	e := mustExplain(t, "* * * * *")
	out := explain.Render(e, false)
	if !strings.Contains(out, "runs every minute") {
		t.Errorf("expected frequency in output, got:\n%s", out)
	}
}

func TestRender_ContainsSummary(t *testing.T) {
	e := mustExplain(t, "0 9 * * 1")
	out := explain.Render(e, false)
	if !strings.Contains(out, "Summary") {
		t.Errorf("expected Summary label in output, got:\n%s", out)
	}
}

func TestRender_ColorModeContainsEscapes(t *testing.T) {
	e := mustExplain(t, "*/5 * * * *")
	out := explain.Render(e, true)
	if !strings.Contains(out, "\033[") {
		t.Errorf("expected ANSI escape codes in color output")
	}
}

func TestRender_NoColorModeNoEscapes(t *testing.T) {
	e := mustExplain(t, "*/5 * * * *")
	out := explain.Render(e, false)
	if strings.Contains(out, "\033[") {
		t.Errorf("expected no ANSI escape codes in plain output")
	}
}
