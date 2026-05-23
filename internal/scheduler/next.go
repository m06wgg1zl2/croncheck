package scheduler

import (
	"time"

	"github.com/croncheck/internal/parser"
)

// NextRuns returns the next n scheduled run times after the given start time,
// computed from the parsed cron expression in the given timezone.
func NextRuns(expr *parser.Expression, start time.Time, n int, loc *time.Location) ([]time.Time, error) {
	if expr == nil {
		return nil, ErrNilExpression
	}
	if loc == nil {
		loc = time.UTC
	}
	if n <= 0 {
		return nil, nil
	}

	results := make([]time.Time, 0, n)
	// Start from the next minute boundary
	current := start.In(loc).Truncate(time.Minute).Add(time.Minute)

	// Safety limit to avoid infinite loops on degenerate expressions
	const maxIterations = 527040 // ~1 year of minutes
	iterations := 0

	for len(results) < n {
		if iterations > maxIterations {
			break
		}
		iterations++

		if matches(expr, current) {
			results = append(results, current)
		}
		current = current.Add(time.Minute)
	}

	return results, nil
}

// matches returns true if the given time satisfies the cron expression.
func matches(expr *parser.Expression, t time.Time) bool {
	return inField(expr.Minute, t.Minute()) &&
		inField(expr.Hour, t.Hour()) &&
		inField(expr.DayOfMonth, t.Day()) &&
		inField(expr.Month, int(t.Month())) &&
		inField(expr.DayOfWeek, int(t.Weekday()))
}

// inField checks whether a value is present in a cron field's allowed values.
func inField(values []int, val int) bool {
	for _, v := range values {
		if v == val {
			return true
		}
	}
	return false
}
