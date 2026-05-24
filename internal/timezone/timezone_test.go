package timezone_test

import (
	"testing"
	"time"

	"github.com/yourorg/croncheck/internal/timezone"
)

func TestResolve_UTC(t *testing.T) {
	for _, input := range []string{"", "UTC", "utc"} {
		loc, err := timezone.Resolve(input)
		if err != nil {
			t.Fatalf("Resolve(%q) unexpected error: %v", input, err)
		}
		if loc != time.UTC {
			t.Errorf("Resolve(%q) = %v, want UTC", input, loc)
		}
	}
}

func TestResolve_Local(t *testing.T) {
	loc, err := timezone.Resolve("Local")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc != time.Local {
		t.Errorf("got %v, want Local", loc)
	}
}

func TestResolve_ValidIANA(t *testing.T) {
	loc, err := timezone.Resolve("America/New_York")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc == nil {
		t.Fatal("expected non-nil location")
	}
	if loc.String() != "America/New_York" {
		t.Errorf("got %v, want America/New_York", loc)
	}
}

func TestResolve_Invalid(t *testing.T) {
	_, err := timezone.Resolve("Mars/Olympus")
	if err == nil {
		t.Fatal("expected error for unknown timezone, got nil")
	}
}

func TestSuggest_MatchesFragment(t *testing.T) {
	results := timezone.Suggest("america")
	if len(results) == 0 {
		t.Fatal("expected at least one result for 'america'")
	}
	for _, r := range results {
		if len(r) == 0 {
			t.Error("empty timezone name in results")
		}
	}
}

func TestSuggest_NoMatch(t *testing.T) {
	results := timezone.Suggest("zzznomatch")
	if len(results) != 0 {
		t.Errorf("expected no results, got %v", results)
	}
}

func TestSuggest_CaseInsensitive(t *testing.T) {
	lower := timezone.Suggest("europe")
	upper := timezone.Suggest("EUROPE")
	if len(lower) != len(upper) {
		t.Errorf("case sensitivity mismatch: %d vs %d", len(lower), len(upper))
	}
}

func TestCommon_NotEmpty(t *testing.T) {
	if len(timezone.Common) == 0 {
		t.Error("Common timezone list must not be empty")
	}
}
