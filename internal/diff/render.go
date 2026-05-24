package diff

import (
	"fmt"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
)

// Render formats a diff.Result as a human-readable terminal string.
// If noColor is true, ANSI escape codes are omitted.
func Render(r *Result, noColor bool) string {
	var sb strings.Builder

	header := fmt.Sprintf("Comparing cron expressions:\n  A: %s\n  B: %s\n", r.Left, r.Right)
	sb.WriteString(header)
	sb.WriteString(strings.Repeat("-", 40) + "\n")
	sb.WriteString(fmt.Sprintf("%-10s  %-15s  %-15s  %s\n", "FIELD", "A", "B", "STATUS"))
	sb.WriteString(strings.Repeat("-", 40) + "\n")

	for _, f := range r.Fields {
		status := "same"
		statusColor := colorGreen
		if !f.Same {
			status = "changed"
			statusColor = colorYellow
		}

		if noColor {
			sb.WriteString(fmt.Sprintf("%-10s  %-15s  %-15s  %s\n",
				f.Name, f.Left, f.Right, status))
		} else {
			lineColor := colorCyan
			if !f.Same {
				lineColor = colorRed
			}
			sb.WriteString(fmt.Sprintf("%s%-10s%s  %-15s  %-15s  %s%s%s\n",
				lineColor, f.Name, colorReset,
				f.Left, f.Right,
				statusColor, status, colorReset))
		}
	}

	changed := r.Changed()
	sb.WriteString(strings.Repeat("-", 40) + "\n")
	if len(changed) == 0 {
		sb.WriteString("Expressions are identical.\n")
	} else {
		sb.WriteString(fmt.Sprintf("%d field(s) differ.\n", len(changed)))
	}
	return sb.String()
}
