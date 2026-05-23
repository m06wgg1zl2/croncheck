package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/user/croncheck/internal/formatter"
	"github.com/user/croncheck/internal/parser"
	"github.com/user/croncheck/internal/scheduler"
)

const defaultCount = 5

func main() {
	tz := flag.String("tz", "UTC", "timezone for output (e.g. America/New_York)")
	count := flag.Int("n", defaultCount, "number of next runs to display")
	describe := flag.Bool("describe", false, "print a human-readable description of the expression")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: croncheck [flags] <cron expression>\n\n")
		fmt.Fprintf(os.Stderr, "Examples:\n")
		fmt.Fprintf(os.Stderr, "  croncheck \"*/5 * * * *\"\n")
		fmt.Fprintf(os.Stderr, "  croncheck -tz America/New_York -n 3 \"0 9 * * 1-5\"\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "error: exactly one cron expression required")
		flag.Usage()
		os.Exit(1)
	}

	expr, err := parser.Parse(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: invalid cron expression: %v\n", err)
		os.Exit(1)
	}

	loc, err := time.LoadLocation(*tz)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: unknown timezone %q: %v\n", *tz, err)
		os.Exit(1)
	}

	if *describe {
		fmt.Println(parser.Describe(expr))
		fmt.Println()
	}

	runs := scheduler.NextRuns(expr, time.Now().In(loc), *count)
	fmt.Print(formatter.Format(expr, runs, loc))
}
