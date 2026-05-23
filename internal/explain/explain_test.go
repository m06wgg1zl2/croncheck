package explain_test

import (
	"strings"
	"testing"

	"croncheck/internal/explain"
)

func TestExplain_EveryMinute(t *testing.T) {
	exp, err := explain.Explain("* * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exp.Frequency != "runs every minute" {
		t.Errorf("expected 'runs every minute', got %q", exp.Frequency)
	}
}

func TestExplain_FieldCount(t *testing.T) {
	exp, err := explain.Explain("0 12 * * 1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(exp.Fields) != 5 {
		t.Errorf("expected 5 fields, got %d", len(exp.Fields))
	}
}

func TestExplain_StarField(t *testing.T) {
	exp, err := explain.Explain("* * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, f := range exp.Fields {
		if !strings.HasPrefix(f.Human, "every ") {
			t.Errorf("field %q: expected 'every ...' got %q", f.Name, f.Human)
		}
	}
}

func TestExplain_StepField(t *testing.T) {
	exp, err := explain.Explain("*/5 * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(exp.Fields[0].Human, "every 5") {
		t.Errorf("expected step explanation, got %q", exp.Fields[0].Human)
	}
}

func TestExplain_InvalidExpression(t *testing.T) {
	_, err := explain.Explain("not a cron")
	if err == nil {
		t.Error("expected error for invalid expression")
	}
}

func TestExplain_SummaryNotEmpty(t *testing.T) {
	exp, err := explain.Explain("0 9 * * 1-5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exp.Summary == "" {
		t.Error("expected non-empty summary")
	}
}

func TestExplain_DailyFrequency(t *testing.T) {
	exp, err := explain.Explain("0 6 * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exp.Frequency != "runs once per day" {
		t.Errorf("expected 'runs once per day', got %q", exp.Frequency)
	}
}
