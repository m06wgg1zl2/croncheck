package scheduler

import (
	"time"

	"github.com/user/croncheck/internal/parser"
)

// Window represents a time range with a start and end.
type Window struct {
	Start time.Time
	End   time.Time
}

// RunsInWindow returns all cron trigger times within the given time window
// for the provided parsed expression. The start is inclusive, end is exclusive.
func RunsInWindow(expr *parser.Expression, w Window) ([]time.Time, error) {
	if expr == nil {
		return nil, ErrNilExpression
	}
	if !w.End.After(w.Start) {
		return []time.Time{}, nil
	}

	var results []time.Time

	// Align to the next whole minute at or after Start.
	current := w.Start.Truncate(time.Minute)
	if current.Before(w.Start) {
		current = current.Add(time.Minute)
	}

	for !current.After(w.End) && !current.Equal(w.End) {
		if matches(expr, current) {
			results = append(results, current)
		}
		current = current.Add(time.Minute)
	}

	return results, nil
}

// WindowDuration is a convenience helper that builds a Window relative to a
// reference time and a duration, then delegates to RunsInWindow.
func WindowDuration(expr *parser.Expression, ref time.Time, d time.Duration) ([]time.Time, error) {
	return RunsInWindow(expr, Window{
		Start: ref,
		End:   ref.Add(d),
	})
}
