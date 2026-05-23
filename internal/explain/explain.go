// Package explain provides human-readable breakdowns of cron expressions,
// including field-by-field analysis and schedule frequency estimation.
package explain

import (
	"fmt"
	"strings"

	"croncheck/internal/parser"
)

// FieldExplanation holds the explanation for a single cron field.
type FieldExplanation struct {
	Name  string
	Raw   string
	Human string
}

// Explanation holds the full breakdown of a cron expression.
type Explanation struct {
	Expression string
	Fields     []FieldExplanation
	Summary    string
	Frequency  string
}

// Explain parses and explains each field of a cron expression.
func Explain(expr string) (*Explanation, error) {
	exp, err := parser.Parse(expr)
	if err != nil {
		return nil, fmt.Errorf("explain: %w", err)
	}

	names := parser.FieldNames()
	raw := strings.Fields(expr)

	fields := make([]FieldExplanation, len(names))
	for i, name := range names {
		fields[i] = FieldExplanation{
			Name:  name,
			Raw:   raw[i],
			Human: explainField(name, raw[i]),
		}
	}

	return &Explanation{
		Expression: expr,
		Fields:     fields,
		Summary:    parser.Describe(exp),
		Frequency:  estimateFrequency(raw),
	}, nil
}

func explainField(name, value string) string {
	if value == "*" {
		return fmt.Sprintf("every %s", strings.ToLower(name))
	}
	if strings.HasPrefix(value, "*/") {
		step := strings.TrimPrefix(value, "*/")
		return fmt.Sprintf("every %s %s(s)", step, strings.ToLower(name))
	}
	if strings.Contains(value, "-") {
		parts := strings.SplitN(value, "-", 2)
		return fmt.Sprintf("%s from %s to %s", strings.ToLower(name), parts[0], parts[1])
	}
	if strings.Contains(value, ",") {
		return fmt.Sprintf("%s at %s", strings.ToLower(name), value)
	}
	return fmt.Sprintf("%s = %s", strings.ToLower(name), value)
}

func estimateFrequency(raw []string) string {
	if raw[0] == "*" && raw[1] == "*" {
		return "runs every minute"
	}
	if raw[0] != "*" && raw[1] == "*" {
		return "runs once per hour"
	}
	if raw[0] != "*" && raw[1] != "*" && raw[2] == "*" && raw[3] == "*" && raw[4] == "*" {
		return "runs once per day"
	}
	return "runs on a custom schedule"
}
