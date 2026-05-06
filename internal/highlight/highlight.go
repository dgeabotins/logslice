package highlight

import "fmt"

// ANSI color codes.
const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	red    = "\033[31m"
	yellow = "\033[33m"
	green  = "\033[32m"
	cyan   = "\033[36m"
	gray   = "\033[90m"
	white  = "\033[97m"
)

// Level returns an ANSI-colored representation of the log level string.
func Level(level string) string {
	switch normalizeLevel(level) {
	case "error", "fatal", "panic":
		return fmt.Sprintf("%s%s%s%s", bold, red, level, reset)
	case "warn", "warning":
		return fmt.Sprintf("%s%s%s", yellow, level, reset)
	case "info":
		return fmt.Sprintf("%s%s%s", green, level, reset)
	case "debug", "trace":
		return fmt.Sprintf("%s%s%s", gray, level, reset)
	default:
		return fmt.Sprintf("%s%s%s", white, level, reset)
	}
}

// Timestamp returns an ANSI-colored representation of a timestamp string.
func Timestamp(ts string) string {
	return fmt.Sprintf("%s%s%s", cyan, ts, reset)
}

// Message returns an ANSI-colored representation of a log message.
func Message(msg string) string {
	return fmt.Sprintf("%s%s%s", bold, msg, reset)
}

// Key returns an ANSI-colored representation of a field key.
func Key(key string) string {
	return fmt.Sprintf("%s%s%s", cyan, key, reset)
}

// Value returns a plain (uncolored) representation of a field value.
func Value(val string) string {
	return fmt.Sprintf("%s%s%s", gray, val, reset)
}

// Strip removes ANSI escape sequences from s (naive, for testing).
func Strip(s string) string {
	out := make([]byte, 0, len(s))
	i := 0
	for i < len(s) {
		if s[i] == '\033' && i+1 < len(s) && s[i+1] == '[' {
			i += 2
			for i < len(s) && s[i] != 'm' {
				i++
			}
			i++ // skip 'm'
			continue
		}
		out = append(out, s[i])
		i++
	}
	return string(out)
}

func normalizeLevel(level string) string {
	if len(level) == 0 {
		return ""
	}
	b := make([]byte, len(level))
	for i := 0; i < len(level); i++ {
		c := level[i]
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		b[i] = c
	}
	return string(b)
}
