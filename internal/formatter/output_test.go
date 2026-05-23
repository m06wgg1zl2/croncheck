package formatter_test

import (
	"strings"
	"testing"
	"time"

	"github.com/user/croncheck/internal/formatter"
	"github.com/user/croncheck/internal/parser"
)

func mustParse(t *testing.T, expr string) *parser.Expression {
	t.Helper()
	e, err := parser.Parse(expr)
	if err != nil {
		t.Fatalf("failed to parse %q: %v", expr, err)
	}
	return e
}

func sampleRuns() []time.Time {
	base := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	return []time.Time{
		base,
		base.Add(time.Minute),
		base.Add(2 * time.Minute),
	}
}

func TestFormat_ContainsExpression(t *testing.T) {
	expr := mustParse(t, "* * * * *")
	out := formatter.Format(expr, sampleRuns(), formatter.Options{NoColor: true})
	if !strings.Contains(out, "* * * * *") {
		t.Errorf("expected output to contain expression, got:\n%s", out)
	}
}

func TestFormat_ContainsTimezone(t *testing.T) {
	expr := mustParse(t, "0 * * * *")
	out := formatter.Format(expr, sampleRuns(), formatter.Options{NoColor: true, Timezone: "America/New_York"})
	if !strings.Contains(out, "America/New_York") {
		t.Errorf("expected timezone in output, got:\n%s", out)
	}
}

func TestFormat_DefaultTimezoneUTC(t *testing.T) {
	expr := mustParse(t, "0 * * * *")
	out := formatter.Format(expr, sampleRuns(), formatter.Options{NoColor: true})
	if !strings.Contains(out, "UTC") {
		t.Errorf("expected UTC timezone default, got:\n%s", out)
	}
}

func TestFormat_ShowDescription(t *testing.T) {
	expr := mustParse(t, "0 9 * * 1")
	out := formatter.Format(expr, sampleRuns(), formatter.Options{NoColor: true, ShowDesc: true})
	if !strings.Contains(out, "Description") {
		t.Errorf("expected description in output, got:\n%s", out)
	}
}

func TestFormat_NilExpression(t *testing.T) {
	out := formatter.Format(nil, sampleRuns(), formatter.Options{})
	if out != "" {
		t.Errorf("expected empty string for nil expression, got: %q", out)
	}
}

func TestFormat_RunsNumbered(t *testing.T) {
	expr := mustParse(t, "* * * * *")
	out := formatter.Format(expr, sampleRuns(), formatter.Options{NoColor: true})
	for _, marker := range []string{" 1.", " 2.", " 3."} {
		if !strings.Contains(out, marker) {
			t.Errorf("expected run marker %q in output, got:\n%s", marker, out)
		}
	}
}
