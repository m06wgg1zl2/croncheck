package validator

import (
	"fmt"
	"strconv"
	"strings"
)

// FieldRange defines the allowed min/max for a cron field.
type FieldRange struct {
	Name string
	Min  int
	Max  int
}

// fieldRanges defines valid ranges for each cron field (minute, hour, dom, month, dow).
var fieldRanges = []FieldRange{
	{"minute", 0, 59},
	{"hour", 0, 23},
	{"day-of-month", 1, 31},
	{"month", 1, 12},
	{"day-of-week", 0, 6},
}

// ValidationError holds a field name and a human-readable reason.
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %s", e.Field, e.Reason)
}

// Validate checks each field of a 5-part cron expression string.
// It returns a slice of ValidationErrors (one per offending field).
func Validate(expr string) []ValidationError {
	parts := strings.Fields(expr)
	if len(parts) != 5 {
		return []ValidationError{{
			Field:  "expression",
			Reason: fmt.Sprintf("expected 5 fields, got %d", len(parts)),
		}}
	}

	var errs []ValidationError
	for i, part := range parts {
		if err := validateField(part, fieldRanges[i]); err != nil {
			errs = append(errs, *err)
		}
	}
	return errs
}

func validateField(field string, r FieldRange) *ValidationError {
	if field == "*" {
		return nil
	}
	// Handle step values e.g. */5 or 1-5/2
	base := field
	if idx := strings.Index(field, "/"); idx != -1 {
		step := field[idx+1:]
		if s, err := strconv.Atoi(step); err != nil || s < 1 {
			return &ValidationError{r.Name, fmt.Sprintf("invalid step value %q", step)}
		}
		base = field[:idx]
		if base == "*" {
			return nil
		}
	}
	// Handle ranges e.g. 1-5
	if strings.Contains(base, "-") {
		parts := strings.SplitN(base, "-", 2)
		lo, err1 := strconv.Atoi(parts[0])
		hi, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return &ValidationError{r.Name, fmt.Sprintf("non-numeric range %q", base)}
		}
		if lo < r.Min || hi > r.Max || lo > hi {
			return &ValidationError{r.Name, fmt.Sprintf("range %d-%d out of bounds [%d-%d]", lo, hi, r.Min, r.Max)}
		}
		return nil
	}
	// Handle lists e.g. 1,3,5
	for _, val := range strings.Split(base, ",") {
		n, err := strconv.Atoi(val)
		if err != nil {
			return &ValidationError{r.Name, fmt.Sprintf("non-numeric value %q", val)}
		}
		if n < r.Min || n > r.Max {
			return &ValidationError{r.Name, fmt.Sprintf("value %d out of bounds [%d-%d]", n, r.Min, r.Max)}
		}
	}
	return nil
}
