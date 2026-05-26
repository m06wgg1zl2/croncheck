package scheduler

import (
	"strings"
	"testing"
	"time"
)

func TestRenderStats_ContainsExpression(t *testing.T) {
	expr := mustParseStats("*/15 * * * *")
	now := time.Now().UTC()
	s, err := ComputeStats(expr, now, now.Add(24*time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := RenderStats(s)
	if !strings.Contains(out, "*/15 * * * *") {
		t.Errorf("expected expression in output, got:\n%s", out)
	}
}

func TestRenderStats_ContainsTotalRuns(t *testing.T) {
	expr := mustParseStats("* * * * *")
	now := time.Now().UTC()
	s, err := ComputeStats(expr, now, now.Add(time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := RenderStats(s)
	if !strings.Contains(out, "Total runs") {
		t.Errorf("expected 'Total runs' in output, got:\n%s", out)
	}
}

func TestRenderStats_ContainsGaps(t *testing.T) {
	expr := mustParseStats("0 * * * *")
	now := time.Now().UTC()
	s, err := ComputeStats(expr, now, now.Add(48*time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := RenderStats(s)
	for _, label := range []string{"Avg gap", "Min gap", "Max gap"} {
		if !strings.Contains(out, label) {
			t.Errorf("expected %q in output, got:\n%s", label, out)
		}
	}
}

func TestRenderStats_NoRuns(t *testing.T) {
	expr := mustParseStats("0 0 31 2 *") // Feb 31 never happens
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	s, err := ComputeStats(expr, now, now.Add(7*24*time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := RenderStats(s)
	if !strings.Contains(out, "No runs") {
		t.Errorf("expected 'No runs' message, got:\n%s", out)
	}
}

func TestRenderStats_EmptyStats(t *testing.T) {
	out := RenderStats(Stats{})
	if !strings.Contains(out, "no stats available") {
		t.Errorf("expected fallback message for empty stats, got: %s", out)
	}
}

func TestRenderStats_ContainsWindow(t *testing.T) {
	expr := mustParseStats("*/30 * * * *")
	now := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	s, err := ComputeStats(expr, now, now.Add(2*time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := RenderStats(s)
	if !strings.Contains(out, "Window") {
		t.Errorf("expected 'Window' in output, got:\n%s", out)
	}
}
