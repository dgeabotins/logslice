package highlight_test

import (
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/highlight"
)

func TestLevel_Error(t *testing.T) {
	for _, lvl := range []string{"error", "ERROR", "fatal", "panic"} {
		out := highlight.Level(lvl)
		if highlight.Strip(out) != lvl {
			t.Errorf("Strip(Level(%q)) = %q, want %q", lvl, highlight.Strip(out), lvl)
		}
		if !strings.Contains(out, "\033[") {
			t.Errorf("Level(%q) contains no ANSI codes", lvl)
		}
	}
}

func TestLevel_Warn(t *testing.T) {
	out := highlight.Level("warn")
	stripped := highlight.Strip(out)
	if stripped != "warn" {
		t.Errorf("got %q, want %q", stripped, "warn")
	}
}

func TestLevel_Info(t *testing.T) {
	out := highlight.Level("info")
	if highlight.Strip(out) != "info" {
		t.Errorf("unexpected stripped value: %q", highlight.Strip(out))
	}
}

func TestLevel_Debug(t *testing.T) {
	out := highlight.Level("debug")
	if highlight.Strip(out) != "debug" {
		t.Errorf("unexpected stripped value: %q", highlight.Strip(out))
	}
}

func TestLevel_Unknown(t *testing.T) {
	out := highlight.Level("custom")
	if highlight.Strip(out) != "custom" {
		t.Errorf("unexpected stripped value: %q", highlight.Strip(out))
	}
}

func TestTimestamp(t *testing.T) {
	ts := "2024-01-02T15:04:05Z"
	out := highlight.Timestamp(ts)
	if highlight.Strip(out) != ts {
		t.Errorf("got %q, want %q", highlight.Strip(out), ts)
	}
}

func TestMessage(t *testing.T) {
	msg := "something happened"
	out := highlight.Message(msg)
	if highlight.Strip(out) != msg {
		t.Errorf("got %q, want %q", highlight.Strip(out), msg)
	}
}

func TestKey(t *testing.T) {
	out := highlight.Key("traceID")
	if highlight.Strip(out) != "traceID" {
		t.Errorf("got %q", highlight.Strip(out))
	}
}

func TestValue(t *testing.T) {
	out := highlight.Value("abc123")
	if highlight.Strip(out) != "abc123" {
		t.Errorf("got %q", highlight.Strip(out))
	}
}

func TestStrip_NoEscapes(t *testing.T) {
	plain := "hello world"
	if highlight.Strip(plain) != plain {
		t.Errorf("Strip modified plain string")
	}
}
