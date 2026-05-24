// Package diff provides utilities for comparing two cron expressions
// field by field and rendering the differences in a human-readable format.
//
// # Usage
//
// Use [Compare] to produce a [Result] containing per-field diffs:
//
//	result, err := diff.Compare("0 9 * * 1", "30 18 * * 5")
//	if err != nil {
//		log.Fatal(err)
//	}
//
// Use [Render] to format the result for terminal output:
//
//	fmt.Print(diff.Render(result, false)) // color enabled
//
// # Field Order
//
// Fields are compared in standard cron order: minute, hour, day, month, weekday.
// Only fields with differing values are returned by [Result.Changed].
package diff
