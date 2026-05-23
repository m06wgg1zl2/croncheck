package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/user/croncheck/internal/parser"
)

// Format produces a human-readable output for a cron expression,
// including its description, next run times, and a timezone-aware table.
func Format(expr *parser.Expression, runs []time.Time, tz *time.Location) string {
	if expr == nil {
		return "Error: nil expression provided."
	}
	if tz == nil {
		tz = time.UTC
	}

	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("Expression : %s\n", expr.Raw))
	sb.WriteString(fmt.Sprintf("Description: %s\n", parser.Describe(expr)))
	sb.WriteString(fmt.Sprintf("Timezone   : %s\n", tz.String()))
	sb.WriteString("\n")

	// Next-run table
	rows := BuildTable(runs, tz)
	sb.WriteString(RenderTable(rows, tz.String()))

	return sb.String()
}
