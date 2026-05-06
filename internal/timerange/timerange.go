// Package timerange provides utilities for filtering log entries by time range.
package timerange

import (
	"fmt"
	"time"
)

// Range represents an inclusive time window [Since, Until].
// A zero value for either bound means that bound is unbounded.
type Range struct {
	Since time.Time
	Until time.Time
}

// New creates a Range from optional since/until strings.
// Accepted formats: RFC3339, RFC3339Nano, and relative durations (e.g. "5m", "2h").
func New(since, until string) (Range, error) {
	var r Range
	var err error

	if since != "" {
		r.Since, err = parse(since)
		if err != nil {
			return Range{}, fmt.Errorf("invalid --since: %w", err)
		}
	}

	if until != "" {
		r.Until, err = parse(until)
		if err != nil {
			return Range{}, fmt.Errorf("invalid --until: %w", err)
		}
	}

	if !r.Since.IsZero() && !r.Until.IsZero() && r.Until.Before(r.Since) {
		return Range{}, fmt.Errorf("--until must not be before --since")
	}

	return r, nil
}

// Contains reports whether t falls within the range.
// A zero bound is treated as open (unbounded).
func (r Range) Contains(t time.Time) bool {
	if t.IsZero() {
		return true
	}
	if !r.Since.IsZero() && t.Before(r.Since) {
		return false
	}
	if !r.Until.IsZero() && t.After(r.Until) {
		return false
	}
	return true
}

// IsZero reports whether the range has no bounds set.
func (r Range) IsZero() bool {
	return r.Since.IsZero() && r.Until.IsZero()
}

// parse tries RFC3339Nano, RFC3339, then relative duration from now.
func parse(s string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return t, nil
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, nil
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return time.Time{}, fmt.Errorf("cannot parse %q as RFC3339 or duration", s)
	}
	return time.Now().Add(-d), nil
}
