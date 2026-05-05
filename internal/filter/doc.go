// Package filter implements log entry filtering for logslice.
//
// A Filter is constructed from Options and exposes a single Match method
// that determines whether a parsed log Entry satisfies all configured
// criteria:
//
//   - Level filtering (case-insensitive, multiple values allowed)
//   - Message regexp matching
//   - Time-range filtering via Since / Until
//   - Arbitrary field key=value matching
//
// Example:
//
//	f, err := filter.New(filter.Options{
//		Levels:     []string{"error", "warn"},
//		MsgPattern: `timeout`,
//		Fields:     map[string]string{"namespace": "prod"},
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	if f.Match(entry) {
//		// process entry
//	}
package filter
