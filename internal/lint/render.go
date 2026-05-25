package lint

import (
	"fmt"
	"strings"
)

// Render formats a slice of lint warnings into a human-readable string
// suitable for terminal output. Returns an empty string when there are
// no warnings.
func Render(expr string, warnings []Warning) string {
	if len(warnings) == 0 {
		return fmt.Sprintf("✔  No lint warnings for: %s\n", expr)
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("⚠  Lint warnings for: %s\n", expr))
	sb.WriteString(strings.Repeat("─", 55) + "\n")

	for i, w := range warnings {
		sb.WriteString(fmt.Sprintf(" %d. [%s] (%s)\n", i+1, w.Code, w.Field))
		sb.WriteString(fmt.Sprintf("    %s\n", w.Message))
	}

	sb.WriteString(fmt.Sprintf("\n%d warning(s) found.\n", len(warnings)))
	return sb.String()
}
