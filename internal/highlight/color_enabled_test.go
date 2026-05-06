//go:build !noansi

package highlight_test

import (
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/highlight"
)

func TestEnabled_True(t *testing.T) {
	if !highlight.Enabled() {
		t.Fatal("expected Enabled() to return true in ANSI build")
	}
}

func TestLevelAuto_ContainsANSI(t *testing.T) {
	cases := []struct {
		level string
		want  string // partial ANSI prefix expected
	}{
		{"error", "\033["},
		{"warn", "\033["},
		{"info", "\033["},
		{"debug", "\033["},
		{"unknown", "unknown"}, // no ANSI for unknown levels
	}
	for _, tc := range cases {
		t.Run(tc.level, func(t *testing.T) {
			got := highlight.LevelAuto(tc.level)
			if !strings.Contains(got, tc.want) {
				t.Errorf("LevelAuto(%q) = %q, want it to contain %q", tc.level, got, tc.want)
			}
		})
	}
}

func TestTimestampAuto_ContainsANSI(t *testing.T) {
	ts := "2024-01-02T15:04:05Z"
	got := highlight.TimestampAuto(ts)
	if !strings.Contains(got, ts) {
		t.Errorf("TimestampAuto(%q) does not contain original string", ts)
	}
	if !strings.Contains(got, "\033[") {
		t.Errorf("TimestampAuto(%q) missing ANSI escape: %q", ts, got)
	}
}

func TestMessageAuto_ContainsANSI(t *testing.T) {
	msg := "something happened"
	got := highlight.MessageAuto(msg)
	if !strings.Contains(got, msg) {
		t.Errorf("MessageAuto(%q) does not contain original string", msg)
	}
	if !strings.Contains(got, "\033[") {
		t.Errorf("MessageAuto(%q) missing ANSI escape: %q", msg, got)
	}
}
