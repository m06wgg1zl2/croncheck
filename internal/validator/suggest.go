package validator

import "fmt"

// CommonExpressions maps a friendly label to a standard cron expression.
var CommonExpressions = map[string]string{
	"every minute":        "* * * * *",
	"every hour":          "0 * * * *",
	"every day at midnight": "0 0 * * *",
	"every day at noon":   "0 12 * * *",
	"every Monday":        "0 0 * * 1",
	"every weekday":       "0 0 * * 1-5",
	"every weekend":       "0 0 * * 0,6",
	"every 5 minutes":     "*/5 * * * *",
	"every 15 minutes":    "*/15 * * * *",
	"every 30 minutes":    "*/30 * * * *",
	"first day of month":  "0 0 1 * *",
	"every 6 hours":       "0 */6 * * *",
}

// Suggestion holds a label and its corresponding cron expression.
type Suggestion struct {
	Label string
	Expr  string
}

// Suggest returns a list of common cron expressions as Suggestion values.
func Suggest() []Suggestion {
	// Return in a deterministic, human-friendly order.
	ordered := []string{
		"every minute",
		"every 5 minutes",
		"every 15 minutes",
		"every 30 minutes",
		"every hour",
		"every 6 hours",
		"every day at midnight",
		"every day at noon",
		"every weekday",
		"every weekend",
		"every Monday",
		"first day of month",
	}
	out := make([]Suggestion, 0, len(ordered))
	for _, label := range ordered {
		out = append(out, Suggestion{Label: label, Expr: CommonExpressions[label]})
	}
	return out
}

// FormatSuggestions returns a human-readable string listing all suggestions.
func FormatSuggestions() string {
	suggestions := Suggest()
	var sb string
	for _, s := range suggestions {
		sb += fmt.Sprintf("  %-28s %s\n", s.Label, s.Expr)
	}
	return sb
}
