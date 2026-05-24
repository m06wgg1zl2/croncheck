package scheduler

import (
	"time"

	"github.com/croncheck/internal/parser"
)

// PrevRuns returns the n most recent past run times for the given cron expression,
// working backwards from the reference time t.
func PrevRuns(expr *parser.Expression, t time.Time, n int) []time.Time {
	if expr == nil || n <= 0 {
		return nil
	}

	// Truncate to the previous minute and step back one minute to start searching.
	current := t.Truncate(time.Minute).Add(-time.Minute)

	results := make([]time.Time, 0, n)

	// Guard against infinite loops: search at most 4 years back.
	limit := current.Add(-4 * 365 * 24 * time.Hour)

	for len(results) < n && current.After(limit) {
		if matches(expr, current) {
			results = append(results, current)
		}
		current = current.Add(-time.Minute)
	}

	return results
}
