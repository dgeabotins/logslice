// Package timerange implements time-window filtering for log entries.
//
// A Range defines an optional lower bound (Since) and upper bound (Until).
// Both bounds are inclusive. Either bound may be omitted (zero value) to
// leave that end open.
//
// Time strings are accepted in the following formats:
//
//	- RFC3339 / RFC3339Nano  (e.g. "2024-01-15T10:00:00Z")
//	- Relative Go durations  (e.g. "5m", "2h", "30s")
//	  A relative duration d is interpreted as time.Now().Add(-d), so
//	  "5m" means "five minutes ago".
//
// Example:
//
//	r, err := timerange.New("30m", "")
//	if err != nil {
//		log.Fatal(err)
//	}
//	if r.Contains(entry.Timestamp) {
//		// process entry
//	}
package timerange
