package scheduler

import (
	"fmt"
	"time"
)

// HumanizeDuration returns a human-readable string describing the duration
// between now and a future or past time.
func HumanizeDuration(d time.Duration) string {
	if d < 0 {
		d = -d
	}

	seconds := int(d.Seconds())
	minutes := int(d.Minutes())
	hours := int(d.Hours())
	days := hours / 24

	switch {
	case seconds < 60:
		return fmt.Sprintf("%ds", seconds)
	case minutes < 60:
		secs := seconds % 60
		if secs == 0 {
			return fmt.Sprintf("%dm", minutes)
		}
		return fmt.Sprintf("%dm %ds", minutes, secs)
	case hours < 24:
		mins := minutes % 60
		if mins == 0 {
			return fmt.Sprintf("%dh", hours)
		}
		return fmt.Sprintf("%dh %dm", hours, mins)
	case days < 7:
		remHours := hours % 24
		if remHours == 0 {
			return fmt.Sprintf("%dd", days)
		}
		return fmt.Sprintf("%dd %dh", days, remHours)
	default:
		weeks := days / 7
		remDays := days % 7
		if remDays == 0 {
			return fmt.Sprintf("%dw", weeks)
		}
		return fmt.Sprintf("%dw %dd", weeks, remDays)
	}
}

// HumanizeNext returns a string like "in 5m 30s" or "5m 30s ago"
// relative to a reference time.
func HumanizeNext(ref, target time.Time) string {
	d := target.Sub(ref)
	if d >= 0 {
		return "in " + HumanizeDuration(d)
	}
	return HumanizeDuration(d) + " ago"
}
