package scheduler

import (
	"time"

	"github.com/user/croncheck/internal/parser"
)

// OverlapResult holds the result of comparing two cron expressions for overlapping run times.
type OverlapResult struct {
	ExprA      string
	ExprB      string
	Overlaps   []time.Time
	TotalA     int
	TotalB     int
}

// FindOverlaps returns times within the given window where both expressions fire at the same minute.
func FindOverlaps(exprA, exprB *parser.Expression, from, to time.Time) (*OverlapResult, error) {
	if exprA == nil || exprB == nil {
		return nil, nil
	}

	runsA := runsInRange(exprA, from, to)
	runsB := runsInRange(exprB, from, to)

	// Index runs from B by truncated minute for O(n) lookup.
	setB := make(map[time.Time]struct{}, len(runsB))
	for _, t := range runsB {
		setB[t.Truncate(time.Minute)] = struct{}{}
	}

	var overlaps []time.Time
	for _, t := range runsA {
		key := t.Truncate(time.Minute)
		if _, ok := setB[key]; ok {
			overlaps = append(overlaps, key)
		}
	}

	return &OverlapResult{
		ExprA:    exprA.Raw,
		ExprB:    exprB.Raw,
		Overlaps: overlaps,
		TotalA:   len(runsA),
		TotalB:   len(runsB),
	}, nil
}

// runsInRange collects all matching times in [from, to).
func runsInRange(expr *parser.Expression, from, to time.Time) []time.Time {
	var results []time.Time
	current := from.Truncate(time.Minute)
	for !current.After(to) {
		if matches(expr, current) {
			results = append(results, current)
		}
		current = current.Add(time.Minute)
	}
	return results
}
