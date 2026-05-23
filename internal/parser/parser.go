package parser

import (
	"fmt"
	"strings"

	"github.com/robfig/cron/v3"
)

// ParseResult holds the result of parsing a cron expression.
type ParseResult struct {
	Expression string
	Schedule   cron.Schedule
	Fields     []string
}

// Parse validates and parses a cron expression.
// Supports standard 5-field expressions (minute hour dom month dow).
func Parse(expr string) (*ParseResult, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, fmt.Errorf("cron expression must not be empty")
	}

	fields := strings.Fields(expr)
	if len(fields) != 5 {
		return nil, fmt.Errorf("expected 5 fields (minute hour dom month dow), got %d", len(fields))
	}

	parser := cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	schedule, err := parser.Parse(expr)
	if err != nil {
		return nil, fmt.Errorf("invalid cron expression %q: %w", expr, err)
	}

	return &ParseResult{
		Expression: expr,
		Schedule:   schedule,
		Fields:     fields,
	}, nil
}

// FieldNames returns human-readable names for the 5 cron fields.
func FieldNames() []string {
	return []string{"minute", "hour", "day-of-month", "month", "day-of-week"}
}
