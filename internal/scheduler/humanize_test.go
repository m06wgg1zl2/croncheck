package scheduler

import (
	"testing"
	"time"
)

func TestHumanizeDuration_Seconds(t *testing.T) {
	d := 45 * time.Second
	got := HumanizeDuration(d)
	if got != "45s" {
		t.Errorf("expected '45s', got %q", got)
	}
}

func TestHumanizeDuration_Minutes(t *testing.T) {
	d := 5*time.Minute + 30*time.Second
	got := HumanizeDuration(d)
	if got != "5m 30s" {
		t.Errorf("expected '5m 30s', got %q", got)
	}
}

func TestHumanizeDuration_ExactMinutes(t *testing.T) {
	d := 10 * time.Minute
	got := HumanizeDuration(d)
	if got != "10m" {
		t.Errorf("expected '10m', got %q", got)
	}
}

func TestHumanizeDuration_Hours(t *testing.T) {
	d := 2*time.Hour + 15*time.Minute
	got := HumanizeDuration(d)
	if got != "2h 15m" {
		t.Errorf("expected '2h 15m', got %q", got)
	}
}

func TestHumanizeDuration_Days(t *testing.T) {
	d := 3 * 24 * time.Hour
	got := HumanizeDuration(d)
	if got != "3d" {
		t.Errorf("expected '3d', got %q", got)
	}
}

func TestHumanizeDuration_Weeks(t *testing.T) {
	d := 14 * 24 * time.Hour
	got := HumanizeDuration(d)
	if got != "2w" {
		t.Errorf("expected '2w', got %q", got)
	}
}

func TestHumanizeDuration_Negative(t *testing.T) {
	d := -90 * time.Second
	got := HumanizeDuration(d)
	if got != "1m 30s" {
		t.Errorf("expected '1m 30s' for negative duration, got %q", got)
	}
}

func TestHumanizeNext_Future(t *testing.T) {
	ref := time.Now()
	target := ref.Add(5 * time.Minute)
	got := HumanizeNext(ref, target)
	if got != "in 5m" {
		t.Errorf("expected 'in 5m', got %q", got)
	}
}

func TestHumanizeNext_Past(t *testing.T) {
	ref := time.Now()
	target := ref.Add(-2 * time.Minute)
	got := HumanizeNext(ref, target)
	if got != "2m ago" {
		t.Errorf("expected '2m ago', got %q", got)
	}
}
