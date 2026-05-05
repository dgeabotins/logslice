package config

import (
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	cfg := Default()
	if cfg.OutputFormat != FormatPretty {
		t.Errorf("expected default format %q, got %q", FormatPretty, cfg.OutputFormat)
	}
	if cfg.BufSize <= 0 {
		t.Errorf("expected positive BufSize, got %d", cfg.BufSize)
	}
}

func TestValidate_ValidFormats(t *testing.T) {
	formats := []Format{FormatPretty, FormatJSON, FormatRaw}
	for _, f := range formats {
		cfg := Default()
		cfg.OutputFormat = f
		if err := cfg.Validate(); err != nil {
			t.Errorf("format %q should be valid, got error: %v", f, err)
		}
	}
}

func TestValidate_UnknownFormat(t *testing.T) {
	cfg := Default()
	cfg.OutputFormat = "xml"
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for unknown format, got nil")
	}
}

func TestValidate_EmptyFormatDefaultsToPretty(t *testing.T) {
	cfg := Default()
	cfg.OutputFormat = ""
	if err := cfg.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.OutputFormat != FormatPretty {
		t.Errorf("expected format to default to %q, got %q", FormatPretty, cfg.OutputFormat)
	}
}

func TestValidate_NegativeBufSize(t *testing.T) {
	cfg := Default()
	cfg.BufSize = -1
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for negative BufSize, got nil")
	}
}

func TestValidate_UntilBeforeSince(t *testing.T) {
	cfg := Default()
	now := time.Now()
	cfg.Since = now
	cfg.Until = now.Add(-time.Hour)
	if err := cfg.Validate(); err == nil {
		t.Error("expected error when until is before since, got nil")
	}
}

func TestValidate_ValidTimeRange(t *testing.T) {
	cfg := Default()
	now := time.Now()
	cfg.Since = now.Add(-time.Hour)
	cfg.Until = now
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for valid time range: %v", err)
	}
}
