/*
Package parser implements structured log line parsing for logslice.

It supports JSON-formatted log lines produced by common Go logging libraries
such as zap, zerolog, logrus, and slog. The parser normalises varying field
names for level, message, and timestamp into a unified LogEntry struct while
preserving all remaining fields for downstream filtering and formatting.

Non-JSON lines (e.g. plain-text stderr output from a container) are returned
as a LogEntry with only the Raw and Message fields populated so that they can
still be displayed by the formatter without being dropped.

Usage:

	entry, err := parser.Parse(line)
	if err != nil {
	    // line was empty
	}
	fmt.Println(entry.Level, entry.Message)
*/
package parser
