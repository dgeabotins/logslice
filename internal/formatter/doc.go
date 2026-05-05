// Package formatter converts parsed log entries into human-readable or
// machine-readable output.
//
// Three output formats are supported:
//
//   - FormatPretty — a single colourised line suitable for interactive
//     terminal sessions. Each line begins with an RFC 3339 timestamp,
//     followed by the log level in brackets, the message, and any
//     additional structured fields rendered as key=value pairs.
//
//   - FormatJSON — the entry is re-encoded as compact JSON. This is
//     useful when piping logslice output into another JSON-aware tool
//     such as jq.
//
//   - FormatRaw — the original unparsed line is passed through
//     unchanged. This preserves byte-for-byte fidelity with the source
//     and is the fastest option.
//
// Usage:
//
//	f := formatter.New(os.Stdout, formatter.FormatPretty, true, time.RFC3339)
//	if err := f.Write(entry); err != nil {
//		log.Fatal(err)
//	}
package formatter
