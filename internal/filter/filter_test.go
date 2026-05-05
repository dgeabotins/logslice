package filter_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/filter"
)

func entry(level, msg string, fields map[string]string) filter.Entry {
	if fields == nil {
		fields = map[string]string{}
	}
	return filter.Entry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   msg,
		Fields:    fields,
	}
}

func TestFilter_LevelMatch(t *testing.T) {
	f, _ := filter.New(filter.Options{Levels: []string{"error"}})
	if !f.Match(entry("error", "boom", nil)) {
		t.Error("expected error level to match")
	}
	if f.Match(entry("info", "ok", nil)) {
		t.Error("expected info level to not match")
	}
}

func TestFilter_LevelCaseInsensitive(t *testing.T) {
	f, _ := filter.New(filter.Options{Levels: []string{"WARN"}})
	if !f.Match(entry("warn", "watch out", nil)) {
		t.Error("expected case-insensitive level match")
	}
}

func TestFilter_MsgRegexp(t *testing.T) {
	f, err := filter.New(filter.Options{MsgPattern: `timeout`})
	if err != nil {
		t.Fatal(err)
	}
	if !f.Match(entry("error", "connection timeout", nil)) {
		t.Error("expected message to match")
	}
	if f.Match(entry("error", "all good", nil)) {
		t.Error("expected message not to match")
	}
}

func TestFilter_InvalidRegexp(t *testing.T) {
	_, err := filter.New(filter.Options{MsgPattern: `[invalid`})
	if err == nil {
		t.Error("expected error for invalid regexp")
	}
}

func TestFilter_TimeRange(t *testing.T) {
	now := time.Now()
	f, _ := filter.New(filter.Options{
		Since: now.Add(-time.Minute),
		Until: now.Add(time.Minute),
	})
	e := entry("info", "msg", nil)
	e.Timestamp = now
	if !f.Match(e) {
		t.Error("expected entry within time range to match")
	}
	e.Timestamp = now.Add(-2 * time.Minute)
	if f.Match(e) {
		t.Error("expected entry before Since to not match")
	}
}

func TestFilter_FieldMatch(t *testing.T) {
	f, _ := filter.New(filter.Options{Fields: map[string]string{"pod": "web-1"}})
	if !f.Match(entry("info", "msg", map[string]string{"pod": "web-1"})) {
		t.Error("expected field match")
	}
	if f.Match(entry("info", "msg", map[string]string{"pod": "web-2"})) {
		t.Error("expected field mismatch to fail")
	}
}

func TestFilter_NoOptions(t *testing.T) {
	f, _ := filter.New(filter.Options{})
	if !f.Match(entry("debug", "anything", nil)) {
		t.Error("empty filter should match everything")
	}
}
