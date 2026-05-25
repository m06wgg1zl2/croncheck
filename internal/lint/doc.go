// Package lint provides heuristic analysis of cron expressions, surfacing
// common mistakes and suspicious patterns that are syntactically valid but
// likely unintended.
//
// # Warning Codes
//
//   - W001: Step value of 1 is redundant (e.g. */1 should be *).
//   - W002: Range with equal start and end (e.g. 5-5 should be 5).
//   - W003: Duplicate value in a comma-separated list.
//   - W004: Both day-of-month and day-of-week are restricted simultaneously,
//     which most cron implementations evaluate with OR semantics.
//
// # Usage
//
//	warnings, err := lint.Lint("*/1 * * * *")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Print(lint.Render("*/1 * * * *", warnings))
package lint
