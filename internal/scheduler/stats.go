package scheduler

import (
	"time"

	"github.com/user/croncheck/internal/parser"
)

// Stats holds frequency statistics for a cron expression over a given window.
type Stats struct {
	Expression  string
	WindowStart time.Time
	WindowEnd   time.Time
	RunCount    int
	AvgInterval time.Duration
	MinInterval time.Duration
	MaxInterval time.Duration
}

// ComputeStats calculates run frequency statistics for the given expression
// over the window [start, end).
func ComputeStats(expr *parser.Expression, start, end time.Time) Stats {
	if expr == nil || !end.After(start) {
		return Stats{WindowStart: start, WindowEnd: end}
	}

	runs := RunsInWindow(expr, start, end)
	s := Stats{
		Expression:  formatExpression(expr),
		WindowStart: start,
		WindowEnd:   end,
		RunCount:    len(runs),
	}

	if len(runs) < 2 {
		return s
	}

	var total time.Duration
	min := runs[1].Sub(runs[0])
	max := runs[1].Sub(runs[0])

	for i := 1; i < len(runs); i++ {
		d := runs[i].Sub(runs[i-1])
		total += d
		if d < min {
			min = d
		}
		if d > max {
			max = d
		}
	}

	s.AvgInterval = total / time.Duration(len(runs)-1)
	s.MinInterval = min
	s.MaxInterval = max

	return s
}

// WindowDuration returns the total duration of the stats window.
func (s Stats) WindowDuration() time.Duration {
	return s.WindowEnd.Sub(s.WindowStart)
}

// RunsPerHour returns the average number of runs per hour over the window.
// Returns 0 if the window duration is less than one hour.
func (s Stats) RunsPerHour() float64 {
	hours := s.WindowDuration().Hours()
	if hours <= 0 {
		return 0
	}
	return float64(s.RunCount) / hours
}

func formatExpression(expr *parser.Expression) string {
	if expr == nil {
		return ""
	}
	return expr.Minute + " " + expr.Hour + " " + expr.DayOfMonth + " " + expr.Month + " " + expr.DayOfWeek
}
