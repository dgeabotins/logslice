package timerange_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/timerange"
)

func mustNew(t *testing.T, since, until string) timerange.Range {
	t.Helper()
	r, err := timerange.New(since, until)
	if err != nil {
		t.Fatalf("New(%q, %q) unexpected error: %v", since, until, err)
	}
	return r
}

func TestNew_Empty(t *testing.T) {
	r := mustNew(t, "", "")
	if !r.IsZero() {
		t.Error("expected zero range")
	}
}

func TestNew_RFC3339(t *testing.T) {
	const s = "2024-01-15T10:00:00Z"
	r := mustNew(t, s, "")
	if r.Since.IsZero() {
		t.Error("Since should not be zero")
	}
}

func TestNew_RelativeDuration(t *testing.T) {
	before := time.Now()
	r := mustNew(t, "5m", "")
	after := time.Now()

	if r.Since.Before(before.Add(-5*time.Minute-time.Second)) ||
		r.Since.After(after.Add(-5*time.Minute+time.Second)) {
		t.Errorf("Since out of expected range: %v", r.Since)
	}
}

func TestNew_UntilBeforeSince(t *testing.T) {
	_, err := timerange.New("2024-06-01T00:00:00Z", "2024-01-01T00:00:00Z")
	if err == nil {
		t.Error("expected error when until < since")
	}
}

func TestNew_InvalidSince(t *testing.T) {
	_, err := timerange.New("not-a-time", "")
	if err == nil {
		t.Error("expected error for invalid since")
	}
}

func TestContains_Unbounded(t *testing.T) {
	r := mustNew(t, "", "")
	if !r.Contains(time.Now()) {
		t.Error("unbounded range should contain any time")
	}
}

func TestContains_WithinBounds(t *testing.T) {
	sinceStr := "2024-01-01T00:00:00Z"
	untilStr := "2024-12-31T23:59:59Z"
	r := mustNew(t, sinceStr, untilStr)

	mid, _ := time.Parse(time.RFC3339, "2024-06-15T12:00:00Z")
	if !r.Contains(mid) {
		t.Error("mid-year time should be within bounds")
	}
}

func TestContains_OutsideBounds(t *testing.T) {
	r := mustNew(t, "2024-06-01T00:00:00Z", "2024-06-30T23:59:59Z")

	early, _ := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
	if r.Contains(early) {
		t.Error("early time should be outside bounds")
	}

	late, _ := time.Parse(time.RFC3339, "2024-12-01T00:00:00Z")
	if r.Contains(late) {
		t.Error("late time should be outside bounds")
	}
}

func TestContains_ZeroTime(t *testing.T) {
	r := mustNew(t, "2024-01-01T00:00:00Z", "")
	if !r.Contains(time.Time{}) {
		t.Error("zero time should always pass through")
	}
}
