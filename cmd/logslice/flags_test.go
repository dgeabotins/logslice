package main

import (
	"flag"
	"os"
	"testing"
)

// resetFlags resets the flag.CommandLine between subtests.
func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}

func TestParseFlags_Defaults(t *testing.T) {
	resetFlags()
	os.Args = []string{"logslice"}

	cfg, err := parseFlags()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Format != "pretty" {
		t.Errorf("expected default format pretty, got %q", cfg.Format)
	}
	if cfg.BufSize <= 0 {
		t.Errorf("expected positive BufSize, got %d", cfg.BufSize)
	}
}

func TestParseFlags_LevelFlag(t *testing.T) {
	resetFlags()
	os.Args = []string{"logslice", "-level=error"}

	cfg, err := parseFlags()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Level != "error" {
		t.Errorf("expected level error, got %q", cfg.Level)
	}
}

func TestParseFlags_SinceUntil(t *testing.T) {
	resetFlags()
	os.Args = []string{"logslice", "-since=2024-01-01T00:00:00Z", "-until=2024-12-31T23:59:59Z"}

	cfg, err := parseFlags()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Since == nil || cfg.Until == nil {
		t.Fatal("expected Since and Until to be set")
	}
	if cfg.Since.Year() != 2024 {
		t.Errorf("unexpected Since year: %d", cfg.Since.Year())
	}
}

func TestParseFlags_InvalidSince(t *testing.T) {
	resetFlags()
	os.Args = []string{"logslice", "-since=not-a-time"}

	_, err := parseFlags()
	if err == nil {
		t.Fatal("expected error for invalid -since value")
	}
}

func TestParseFlags_PositionalLevel(t *testing.T) {
	resetFlags()
	os.Args = []string{"logslice", "warn"}

	cfg, err := parseFlags()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Level != "warn" {
		t.Errorf("expected level warn from positional arg, got %q", cfg.Level)
	}
}
