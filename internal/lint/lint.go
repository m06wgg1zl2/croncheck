// Package lint provides heuristic checks on cron expressions,
// warning about suspicious or potentially unintended patterns.
package lint

import (
	"fmt"
	"strconv"
	"strings"

	"croncheck/internal/parser"
)

// Warning represents a single lint advisory for a cron expression.
type Warning struct {
	Field   string
	Code    string
	Message string
}

// Lint analyses a raw cron expression and returns a slice of warnings.
// An empty slice means no issues were detected.
func Lint(expr string) ([]Warning, error) {
	e, err := parser.Parse(expr)
	if err != nil {
		return nil, fmt.Errorf("lint: parse error: %w", err)
	}

	var warnings []Warning
	fields := []struct {
		name  string
		value string
	}{
		{"minute", e.Minute},
		{"hour", e.Hour},
		{"day-of-month", e.DayOfMonth},
		{"month", e.Month},
		{"day-of-week", e.DayOfWeek},
	}

	for _, f := range fields {
		warnings = append(warnings, checkField(f.name, f.value)...)
	}

	// Warn when both day-of-month and day-of-week are restricted.
	if e.DayOfMonth != "*" && e.DayOfWeek != "*" {
		warnings = append(warnings, Warning{
			Field:   "day-of-month+day-of-week",
			Code:    "W004",
			Message: "both day-of-month and day-of-week are restricted; most implementations use OR semantics which may be unintended",
		})
	}

	return warnings, nil
}

func checkField(name, value string) []Warning {
	var w []Warning

	// Detect step of 1 (e.g. */1) which is equivalent to *.
	if strings.Contains(value, "/") {
		parts := strings.SplitN(value, "/", 2)
		if step, err := strconv.Atoi(parts[1]); err == nil && step == 1 {
			w = append(w, Warning{
				Field:   name,
				Code:    "W001",
				Message: fmt.Sprintf("step of 1 in %q is redundant; use \"*\" instead", value),
			})
		}
	}

	// Detect a range where start equals end (e.g. 5-5).
	if strings.Contains(value, "-") && !strings.Contains(value, "/") {
		parts := strings.SplitN(value, "-", 2)
		lo, errLo := strconv.Atoi(parts[0])
		hi, errHi := strconv.Atoi(parts[1])
		if errLo == nil && errHi == nil && lo == hi {
			w = append(w, Warning{
				Field:   name,
				Code:    "W002",
				Message: fmt.Sprintf("range %q has equal start and end; use a single value instead", value),
			})
		}
	}

	// Detect lists with duplicate values.
	if strings.Contains(value, ",") {
		seen := map[string]bool{}
		for _, item := range strings.Split(value, ",") {
			if seen[item] {
				w = append(w, Warning{
					Field:   name,
					Code:    "W003",
					Message: fmt.Sprintf("duplicate value %q in list for field %s", item, name),
				})
				break
			}
			seen[item] = true
		}
	}

	return w
}
