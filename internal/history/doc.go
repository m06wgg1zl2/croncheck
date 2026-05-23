// Package history provides persistent storage and formatting of cron expression
// check history for the croncheck tool.
//
// History entries are stored as JSON in the user's home directory under
// ~/.croncheck/history.json by default. Each entry records the expression,
// timezone, timestamp, and whether validation succeeded.
//
// Usage:
//
//	store, err := history.NewStore(path)
//	if err != nil { ... }
//
//	store.Add(history.Entry{
//		Expression: "0 9 * * 1",
//		Timezone:   "America/New_York",
//		CheckedAt:  time.Now().UTC(),
//		Valid:      true,
//	})
//
//	entries, err := store.Load()
//	fmt.Print(history.FormatTable(entries, time.UTC))
package history
