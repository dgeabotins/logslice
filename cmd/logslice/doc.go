// Package main is the entry point for the logslice CLI tool.
//
// logslice reads a JSON log stream from stdin, filters entries by level,
// message content, and time range, then writes formatted output to stdout.
//
// Usage:
//
//	kubectl logs <pod> | logslice [flags]
//
// Flags:
//
//	-level string    minimum log level (debug|info|warn|error)
//	-msg   string    regex filter on the message field
//	-since string    RFC3339 lower-bound timestamp
//	-until string    RFC3339 upper-bound timestamp
//	-format string   output format: pretty|json|raw (default "pretty")
//	-color           enable ANSI colour in pretty mode
//	-no-ts           omit the timestamp column
//	-bufsize int     scanner buffer size in bytes
//
// logslice exits 0 on success, 1 on any error.
package main
