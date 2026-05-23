package formatter

import (
	"fmt"
	"strings"
	"time"
)

// TableRow represents a single row in the next-run table.
type TableRow struct {
	Index    int
	UTC      time.Time
	Local    time.Time
	Relative string
}

// BuildTable constructs a slice of TableRows from a list of run times and a target timezone.
func BuildTable(runs []time.Time, tz *time.Location) []TableRow {
	if tz == nil {
		tz = time.UTC
	}
	now := time.Now().UTC()
	rows := make([]TableRow, 0, len(runs))
	for i, t := range runs {
		rows = append(rows, TableRow{
			Index:    i + 1,
			UTC:      t.UTC(),
			Local:    t.In(tz),
			Relative: relativeTime(now, t),
		})
	}
	return rows
}

// RenderTable returns a formatted ASCII table string from the given rows.
func RenderTable(rows []TableRow, tzName string) string {
	if len(rows) == 0 {
		return "No upcoming runs."
	}

	var sb strings.Builder
	header := fmt.Sprintf("%-4s  %-20s  %-20s  %s\n", "#", "UTC", tzName, "Relative")
	sep := strings.Repeat("-", len(strings.TrimRight(header, "\n"))) + "\n"

	sb.WriteString(sep)
	sb.WriteString(header)
	sb.WriteString(sep)

	for _, r := range rows {
		sb.WriteString(fmt.Sprintf("%-4d  %-20s  %-20s  %s\n",
			r.Index,
			r.UTC.Format("2006-01-02 15:04:05"),
			r.Local.Format("2006-01-02 15:04:05"),
			r.Relative,
		))
	}
	sb.WriteString(sep)
	return sb.String()
}

// relativeTime returns a human-readable duration string from now until t.
func relativeTime(now, t time.Time) string {
	d := t.Sub(now)
	if d < 0 {
		return "in the past"
	}
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	if h > 0 {
		return fmt.Sprintf("in %dh %dm", h, m)
	}
	if m > 0 {
		return fmt.Sprintf("in %dm %ds", m, s)
	}
	return fmt.Sprintf("in %ds", s)
}
