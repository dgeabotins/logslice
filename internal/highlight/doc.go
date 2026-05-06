// Package highlight provides ANSI terminal color helpers for log output.
//
// Functions in this package wrap log field values (level, timestamp, message,
// key, value) with ANSI escape sequences so they render with color in
// compatible terminals.
//
// Color mapping:
//
//	error / fatal / panic  → bold red
//	warn  / warning        → yellow
//	info                   → green
//	debug / trace          → gray
//	timestamp              → cyan
//	message                → bold white
//	field key              → cyan
//	field value            → gray
//
// The Strip function removes all ANSI escape sequences from a string, which
// is primarily useful in tests to assert on the underlying text content.
package highlight
