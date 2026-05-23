package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/user/croncheck/internal/parser"
)

const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
	colorBold   = "\033[1m"
)

// Options controls output formatting behavior.
type Options struct {
	Timezone   string
	Count      int
	NoColor    bool
	ShowDesc   bool
}

// Format renders a cron expression and its next run times as a terminal string.
func Format(expr *parser.Expression, runs []time.Time, opts Options) string {
	if expr == nil {
		return ""
	}

	var sb strings.Builder

	if opts.NoColor {
		sb.WriteString(fmt.Sprintf("Expression : %s\n", expr.Raw))
	} else {
		sb.WriteString(fmt.Sprintf("%s%sExpression%s : %s%s%s\n",
			colorBold, colorCyan, colorReset,
			colorGreen, expr.Raw, colorReset))
	}

	if opts.ShowDesc {
		desc := parser.Describe(expr)
		if opts.NoColor {
			sb.WriteString(fmt.Sprintf("Description: %s\n", desc))
		} else {
			sb.WriteString(fmt.Sprintf("%sDescription%s: %s\n", colorYellow, colorReset, desc))
		}
	}

	tz := opts.Timezone
	if tz == "" {
		tz = "UTC"
	}

	if opts.NoColor {
		sb.WriteString(fmt.Sprintf("Timezone   : %s\n", tz))
		sb.WriteString("Next runs  :\n")
	} else {
		sb.WriteString(fmt.Sprintf("%s%sTimezone%s   : %s\n", colorBold, colorCyan, colorReset, tz))
		sb.WriteString(fmt.Sprintf("%s%sNext runs%s  :\n", colorBold, colorCyan, colorReset))
	}

	for i, t := range runs {
		formatted := t.Format("2006-01-02 15:04:05 MST")
		if opts.NoColor {
			sb.WriteString(fmt.Sprintf("  %2d. %s\n", i+1, formatted))
		} else {
			sb.WriteString(fmt.Sprintf("  %s%2d.%s %s%s%s\n",
				colorBold, i+1, colorReset,
				colorGreen, formatted, colorReset))
		}
	}

	return sb.String()
}
