package validator

import (
	"fmt"
	"strings"
)

// Report summarises the result of validating a cron expression.
type Report struct {
	Expression string
	Valid       bool
	Errors      []ValidationError
}

// NewReport validates expr and returns a populated Report.
func NewReport(expr string) Report {
	errs := Validate(expr)
	return Report{
		Expression: expr,
		Valid:       len(errs) == 0,
		Errors:      errs,
	}
}

// String returns a human-readable summary of the report.
func (r Report) String() string {
	var sb strings.Builder
	if r.Valid {
		sb.WriteString(fmt.Sprintf("✓ Expression %q is valid.\n", r.Expression))
	} else {
		sb.WriteString(fmt.Sprintf("✗ Expression %q is invalid:\n", r.Expression))
		for _, e := range r.Errors {
			sb.WriteString(fmt.Sprintf("  - %s\n", e.Error()))
		}
	}
	return sb.String()
}

// HasError returns true if the report contains at least one ValidationError
// whose Field matches the given field name.
func (r Report) HasError(field string) bool {
	for _, e := range r.Errors {
		if e.Field == field {
			return true
		}
	}
	return false
}
