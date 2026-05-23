package history

import (
	"fmt"
	"strings"
	"time"
)

const (
	colWidth = 22
	statusOK  = "✓ valid"
	statusErr = "✗ invalid"
)

// FormatTable renders a slice of history entries as a plain-text table.
func FormatTable(entries []Entry, tz *time.Location) string {
	if tz == nil {
		tz = time.UTC
	}
	if len(entries) == 0 {
		return "No history entries found.\n"
	}

	var sb strings.Builder
	header := fmt.Sprintf("%-5s %-20s %-15s %-10s %s\n",
		"#", "Checked At", "Timezone", "Status", "Expression")
	sep := strings.Repeat("-", 75) + "\n"
	sb.WriteString(header)
	sb.WriteString(sep)

	for i, e := range entries {
		ts := e.CheckedAt.In(tz).Format("2006-01-02 15:04:05")
		status := statusOK
		if !e.Valid {
			status = statusErr
		}
		zone := e.Timezone
		if zone == "" {
			zone = "UTC"
		}
		line := fmt.Sprintf("%-5d %-20s %-15s %-10s %s\n",
			i+1, ts, zone, status, e.Expression)
		sb.WriteString(line)
	}
	return sb.String()
}

// Summary returns a brief human-readable summary of the history.
func Summary(entries []Entry) string {
	if len(entries) == 0 {
		return "No history recorded."
	}
	valid := 0
	for _, e := range entries {
		if e.Valid {
			valid++
		}
	}
	return fmt.Sprintf("%d checks recorded: %d valid, %d invalid.",
		len(entries), valid, len(entries)-valid)
}
