package diff_test

import (
	"strings"
	"testing"

	"github.com/user/croncheck/internal/diff"
)

func mustCompare(t *testing.T, left, right string) *diff.Result {
	t.Helper()
	r, err := diff.Compare(left, right)
	if err != nil {
		t.Fatalf("Compare(%q, %q) failed: %v", left, right, err)
	}
	return r
}

func TestRender_ContainsBothExpressions(t *testing.T) {
	r := mustCompare(t, "0 9 * * 1", "30 18 * * 5")
	out := diff.Render(r, true)
	if !strings.Contains(out, "0 9 * * 1") {
		t.Error("expected left expression in output")
	}
	if !strings.Contains(out, "30 18 * * 5") {
		t.Error("expected right expression in output")
	}
}

func TestRender_ContainsFieldNames(t *testing.T) {
	r := mustCompare(t, "* * * * *", "* * * * *")
	out := diff.Render(r, true)
	for _, name := range []string{"minute", "hour", "day", "month", "weekday"} {
		if !strings.Contains(out, name) {
			t.Errorf("expected field name %q in output", name)
		}
	}
}

func TestRender_SameShowsIdentical(t *testing.T) {
	r := mustCompare(t, "* * * * *", "* * * * *")
	out := diff.Render(r, true)
	if !strings.Contains(out, "identical") {
		t.Error("expected 'identical' in output for same expressions")
	}
}

func TestRender_ChangedShowsCount(t *testing.T) {
	r := mustCompare(t, "0 9 * * 1", "30 18 * * 5")
	out := diff.Render(r, true)
	if !strings.Contains(out, "3 field(s) differ") {
		t.Errorf("expected diff count in output, got:\n%s", out)
	}
}

func TestRender_NoColorOmitsEscapes(t *testing.T) {
	r := mustCompare(t, "0 * * * *", "30 * * * *")
	out := diff.Render(r, true)
	if strings.Contains(out, "\033[") {
		t.Error("expected no ANSI escape codes in no-color mode")
	}
}

func TestRender_ColorIncludesEscapes(t *testing.T) {
	r := mustCompare(t, "0 * * * *", "30 * * * *")
	out := diff.Render(r, false)
	if !strings.Contains(out, "\033[") {
		t.Error("expected ANSI escape codes in color mode")
	}
}
