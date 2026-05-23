package parser

import (
	"fmt"
	"strings"
)

// Describe returns a human-readable summary of each cron field.
func Describe(result *ParseResult) string {
	if result == nil || len(result.Fields) != 5 {
		return "invalid expression"
	}

	names := FieldNames()
	var parts []string
	for i, field := range result.Fields {
		parts = append(parts, fmt.Sprintf("%s=%s", names[i], describeField(field)))
	}
	return strings.Join(parts, ", ")
}

// describeField returns a short description for a single cron field value.
func describeField(field string) string {
	switch field {
	case "*":
		return "every"
	case "0":
		return "0"
	}

	if strings.HasPrefix(field, "*/") {
		step := strings.TrimPrefix(field, "*/")
		return fmt.Sprintf("every %s", step)
	}

	if strings.Contains(field, "-") {
		parts := strings.SplitN(field, "-", 2)
		return fmt.Sprintf("%s through %s", parts[0], parts[1])
	}

	if strings.Contains(field, ",") {
		return fmt.Sprintf("[%s]", field)
	}

	return field
}
