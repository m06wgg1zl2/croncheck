// Package diff compares two cron expressions and highlights their differences.
package diff

import (
	"fmt"
	"strings"

	"github.com/user/croncheck/internal/parser"
)

// FieldDiff represents the difference in a single cron field.
type FieldDiff struct {
	Name  string
	Left  string
	Right string
	Same  bool
}

// Result holds the full diff between two cron expressions.
type Result struct {
	Left   string
	Right  string
	Fields []FieldDiff
}

// Changed returns only the fields that differ between the two expressions.
func (r *Result) Changed() []FieldDiff {
	var out []FieldDiff
	for _, f := range r.Fields {
		if !f.Same {
			out = append(out, f)
		}
	}
	return out
}

// Compare parses both expressions and returns a field-by-field diff.
// Returns an error if either expression is invalid.
func Compare(left, right string) (*Result, error) {
	lExpr, err := parser.Parse(left)
	if err != nil {
		return nil, fmt.Errorf("left expression: %w", err)
	}
	rExpr, err := parser.Parse(right)
	if err != nil {
		return nil, fmt.Errorf("right expression: %w", err)
	}

	lParts := strings.Fields(left)
	rParts := strings.Fields(right)
	names := parser.FieldNames()

	_ = lExpr
	_ = rExpr

	fields := make([]FieldDiff, len(names))
	for i, name := range names {
		lVal := lParts[i]
		rVal := rParts[i]
		fields[i] = FieldDiff{
			Name:  name,
			Left:  lVal,
			Right: rVal,
			Same:  lVal == rVal,
		}
	}

	return &Result{
		Left:   left,
		Right:  right,
		Fields: fields,
	}, nil
}
