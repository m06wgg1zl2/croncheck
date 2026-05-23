package explain

import (
	"fmt"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
)

// Render returns a formatted terminal string for an Explanation.
func Render(e *Explanation, useColor bool) string {
	var sb strings.Builder

	header := fmt.Sprintf("Expression: %s", e.Expression)
	if useColor {
		header = fmt.Sprintf("%s%sExpression:%s %s%s%s",
			colorBold, colorCyan, colorReset,
			colorYellow, e.Expression, colorReset)
	}
	sb.WriteString(header + "\n")
	sb.WriteString(strings.Repeat("-", 40) + "\n")

	for _, f := range e.Fields {
		line := fmt.Sprintf("  %-10s %-6s %s", f.Name, f.Raw, f.Human)
		if useColor {
			line = fmt.Sprintf("  %s%-10s%s %-6s %s%s%s",
				colorBold, f.Name, colorReset,
				f.Raw,
				colorGreen, f.Human, colorReset)
		}
		sb.WriteString(line + "\n")
	}

	sb.WriteString(strings.Repeat("-", 40) + "\n")

	freqLine := fmt.Sprintf("Frequency : %s", e.Frequency)
	summaryLine := fmt.Sprintf("Summary   : %s", e.Summary)
	if useColor {
		freqLine = fmt.Sprintf("%sFrequency :%s %s", colorBold, colorReset, e.Frequency)
		summaryLine = fmt.Sprintf("%sSummary   :%s %s", colorBold, colorReset, e.Summary)
	}
	sb.WriteString(freqLine + "\n")
	sb.WriteString(summaryLine + "\n")

	return sb.String()
}
