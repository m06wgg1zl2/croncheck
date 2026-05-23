// Package explain provides detailed, human-readable explanations of cron
// expressions for display in the terminal.
//
// It breaks down each field (minute, hour, day, month, weekday) into plain
// English, estimates the run frequency, and produces a formatted summary
// suitable for both plain-text and ANSI-color terminal output.
//
// Example usage:
//
//	e, err := explain.Explain("0 9 * * 1-5")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Print(explain.Render(e, true))
//
// Output (approximate):
//
//	Expression: 0 9 * * 1-5
//	----------------------------------------
//	  Minute     0      minute = 0
//	  Hour       9      hour = 9
//	  Day        *      every day
//	  Month      *      every month
//	  Weekday    1-5    weekday from 1 to 5
//	----------------------------------------
//	Frequency : runs on a custom schedule
//	Summary   : At 09:00, Monday through Friday
package explain
