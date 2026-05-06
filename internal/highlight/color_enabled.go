//go:build !noansi

package highlight

import "strings"

const (
	reset  = "\033[0m"
	bold   = "\033[1m"

	red     = "\033[31m"
	yellow  = "\033[33m"
	green   = "\033[32m"
	cyan    = "\033[36m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	gray    = "\033[90m"
)

// Enabled reports whether ANSI color output is available.
func Enabled() bool { return true }

// LevelAuto colorizes a log level string using ANSI escape codes.
func LevelAuto(level string) string {
	switch strings.ToLower(level) {
	case "error", "err", "fatal", "panic":
		return red + bold + level + reset
	case "warn", "warning":
		return yellow + level + reset
	case "info":
		return green + level + reset
	case "debug", "trace":
		return gray + level + reset
	default:
		return level
	}
}

// TimestampAuto colorizes a timestamp string.
func TimestampAuto(ts string) string {
	return cyan + ts + reset
}

// MessageAuto colorizes a log message string.
func MessageAuto(msg string) string {
	return bold + msg + reset
}
