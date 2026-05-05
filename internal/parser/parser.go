// Package parser provides JSON log line parsing for Kubernetes pod log streams.
package parser

import (
	"encoding/json"
	"fmt"
	"time"
)

// LogEntry represents a single parsed structured log line.
type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
	Fields    map[string]any
	Raw       string
}

// knownLevelKeys are common field names used for log level across frameworks.
var knownLevelKeys = []string{"level", "severity", "lvl"}

// knownMessageKeys are common field names used for the log message.
var knownMessageKeys = []string{"msg", "message", "text"}

// knownTimeKeys are common field names used for the timestamp.
var knownTimeKeys = []string{"time", "timestamp", "ts", "@timestamp"}

// Parse attempts to parse a raw log line as a JSON structured log entry.
// Non-JSON lines are returned as an entry with only the Raw field set.
func Parse(line string) (*LogEntry, error) {
	if line == "" {
		return nil, fmt.Errorf("empty line")
	}

	entry := &LogEntry{Raw: line, Fields: make(map[string]any)}

	var raw map[string]any
	if err := json.Unmarshal([]byte(line), &raw); err != nil {
		// Not JSON — treat as plain text
		entry.Message = line
		return entry, nil
	}

	for _, key := range knownLevelKeys {
		if v, ok := raw[key]; ok {
			entry.Level = fmt.Sprintf("%v", v)
			delete(raw, key)
			break
		}
	}

	for _, key := range knownMessageKeys {
		if v, ok := raw[key]; ok {
			entry.Message = fmt.Sprintf("%v", v)
			delete(raw, key)
			break
		}
	}

	for _, key := range knownTimeKeys {
		if v, ok := raw[key]; ok {
			entry.Timestamp = parseTime(fmt.Sprintf("%v", v))
			delete(raw, key)
			break
		}
	}

	entry.Fields = raw
	return entry, nil
}

func parseTime(s string) time.Time {
	formats := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02T15:04:05.999999999Z",
	}
	for _, f := range formats {
		if t, err := time.Parse(f, s); err == nil {
			return t
		}
	}
	return time.Time{}
}
