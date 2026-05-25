package scheduler

import (
	"testing"
	"time"

	"github.com/user/croncheck/internal/parser"
)

func mustParseStats(t *testing.T, expr string) *parser.Expression {
	t.Helper()
	e, err := parser.Parse(expr)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	return e
}

func TestComputeStats_EveryMinute(t *testing.T) {
	expr := mustParseStats(t, "* * * * *")
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(60 * time.Minute)

	s := ComputeStats(expr, start, end)

	if s.RunCount != 60 {
		t.Errorf("expected 60 runs, got %d", s.RunCount)
	}
	if s.AvgInterval != time.Minute {
		t.Errorf("expected avg interval 1m, got %v", s.AvgInterval)
	}
	if s.MinInterval != time.Minute {
		t.Errorf("expected min interval 1m, got %v", s.MinInterval)
	}
	if s.MaxInterval != time.Minute {
		t.Errorf("expected max interval 1m, got %v", s.MaxInterval)
	}
}

func TestComputeStats_Hourly(t *testing.T) {
	expr := mustParseStats(t, "0 * * * *")
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(24 * time.Hour)

	s := ComputeStats(expr, start, end)

	if s.RunCount != 24 {
		t.Errorf("expected 24 runs, got %d", s.RunCount)
	}
	if s.AvgInterval != time.Hour {
		t.Errorf("expected avg interval 1h, got %v", s.AvgInterval)
	}
}

func TestComputeStats_NilExpression(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	s := ComputeStats(nil, start, start.Add(time.Hour))

	if s.RunCount != 0 {
		t.Errorf("expected 0 runs for nil expression, got %d", s.RunCount)
	}
}

func TestComputeStats_EndBeforeStart(t *testing.T) {
	expr := mustParseStats(t, "* * * * *")
	start := time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC)
	end := start.Add(-time.Hour)

	s := ComputeStats(expr, start, end)

	if s.RunCount != 0 {
		t.Errorf("expected 0 runs, got %d", s.RunCount)
	}
}

func TestComputeStats_SingleRun(t *testing.T) {
	expr := mustParseStats(t, "30 12 * * *")
	start := time.Date(2024, 1, 1, 12, 29, 0, 0, time.UTC)
	end := time.Date(2024, 1, 1, 12, 31, 0, 0, time.UTC)

	s := ComputeStats(expr, start, end)

	if s.RunCount != 1 {
		t.Errorf("expected 1 run, got %d", s.RunCount)
	}
	if s.AvgInterval != 0 {
		t.Errorf("expected zero avg interval for single run, got %v", s.AvgInterval)
	}
}

func TestComputeStats_ExpressionString(t *testing.T) {
	expr := mustParseStats(t, "5 4 * * *")
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(24 * time.Hour)

	s := ComputeStats(expr, start, end)

	if s.Expression == "" {
		t.Error("expected non-empty expression string")
	}
}
