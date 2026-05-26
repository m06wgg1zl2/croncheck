package scheduler

import (
	"fmt"
	"strings"
)

// RenderStats formats a Stats result as a human-readable terminal block.
func RenderStats(s Stats) string {
	if s.Expression == "" {
		return "no stats available\n"
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Expression : %s\n", s.Expression))
	sb.WriteString(fmt.Sprintf("Window     : %s → %s\n",
		s.WindowStart.Format("2006-01-02 15:04:05 MST"),
		s.WindowEnd.Format("2006-01-02 15:04:05 MST"),
	))
	sb.WriteString(fmt.Sprintf("Total runs : %d\n", s.TotalRuns))

	if s.TotalRuns == 0 {
		sb.WriteString("No runs in the given window.\n")
		return sb.String()
	}

	sb.WriteString(fmt.Sprintf("Avg gap    : %s\n", HumanizeDuration(s.AvgGap)))
	sb.WriteString(fmt.Sprintf("Min gap    : %s\n", HumanizeDuration(s.MinGap)))
	sb.WriteString(fmt.Sprintf("Max gap    : %s\n", HumanizeDuration(s.MaxGap)))

	if !s.FirstRun.IsZero() {
		sb.WriteString(fmt.Sprintf("First run  : %s\n", s.FirstRun.Format("2006-01-02 15:04:05 MST")))
	}
	if !s.LastRun.IsZero() {
		sb.WriteString(fmt.Sprintf("Last run   : %s\n", s.LastRun.Format("2006-01-02 15:04:05 MST")))
	}

	return sb.String()
}
