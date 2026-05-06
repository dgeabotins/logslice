package highlight

import "os"

// Enabled reports whether ANSI color output should be used.
//
// Color is disabled when the NO_COLOR environment variable is set (non-empty)
// or when the TERM environment variable is set to "dumb".
func Enabled() bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	if os.Getenv("TERM") == "dumb" {
		return false
	}
	return true
}

// LevelAuto returns a colored level string when color is enabled, or the
// plain level string otherwise.
func LevelAuto(level string) string {
	if Enabled() {
		return Level(level)
	}
	return level
}

// TimestampAuto returns a colored timestamp when color is enabled, or the
// plain timestamp otherwise.
func TimestampAuto(ts string) string {
	if Enabled() {
		return Timestamp(ts)
	}
	return ts
}

// MessageAuto returns a colored message when color is enabled, or the plain
// message otherwise.
func MessageAuto(msg string) string {
	if Enabled() {
		return Message(msg)
	}
	return msg
}
