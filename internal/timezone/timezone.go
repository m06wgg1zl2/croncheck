// Package timezone provides helpers for resolving and listing
// IANA timezone locations used throughout croncheck.
package timezone

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// Common is a curated list of commonly used IANA timezone names
// shown as suggestions when an invalid timezone is provided.
var Common = []string{
	"UTC",
	"Local",
	"America/New_York",
	"America/Chicago",
	"America/Denver",
	"America/Los_Angeles",
	"Europe/London",
	"Europe/Paris",
	"Europe/Berlin",
	"Asia/Tokyo",
	"Asia/Shanghai",
	"Asia/Kolkata",
	"Australia/Sydney",
	"Pacific/Auckland",
}

// Resolve parses the given timezone string and returns a *time.Location.
// It accepts IANA names (e.g. "America/New_York") as well as the special
// values "UTC" and "Local". Returns an error if the name is unrecognised.
func Resolve(tz string) (*time.Location, error) {
	if tz == "" || strings.EqualFold(tz, "UTC") {
		return time.UTC, nil
	}
	if strings.EqualFold(tz, "Local") {
		return time.Local, nil
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, fmt.Errorf("unknown timezone %q: %w", tz, err)
	}
	return loc, nil
}

// Suggest returns Common timezone names that contain the given substring
// (case-insensitive). Results are sorted alphabetically.
func Suggest(fragment string) []string {
	frag := strings.ToLower(fragment)
	var matches []string
	for _, tz := range Common {
		if strings.Contains(strings.ToLower(tz), frag) {
			matches = append(matches, tz)
		}
	}
	sort.Strings(matches)
	return matches
}
