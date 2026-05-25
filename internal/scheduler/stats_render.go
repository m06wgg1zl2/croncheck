package scheduler

import (
	"fmt"
	"strings"
)

// RenderStats formats a Stats value as a human-readable terminal string.
func RenderStats(s Stats) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Expression : %s\n", s.Expression)
	fmt.Fprintf(&b, "Window     : %s → %s\n",
		s.WindowStart.Format("2006-01-02 15:04 MST"),
		s.WindowEnd.Format("2006-01-02 15:04 MST"),
	)
	fmt.Fprintf(&b, "Run count  : %d\n", s.RunCount)

	if s.RunCount < 2 {
		b.WriteString("Interval   : n/a (fewer than 2 runs)\n")
		return b.String()
	}

	fmt.Fprintf(&b, "Avg interval: %s\n", HumanizeDuration(s.AvgInterval))
	fmt.Fprintf(&b, "Min interval: %s\n", HumanizeDuration(s.MinInterval))
	fmt.Fprintf(&b, "Max interval: %s\n", HumanizeDuration(s.MaxInterval))

	return b.String()
}
