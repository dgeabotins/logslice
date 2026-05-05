// Package config provides configuration loading and validation for logslice.
package config

import (
	"errors"
	"time"
)

// Format controls the output rendering style.
type Format string

const (
	FormatPretty Format = "pretty"
	FormatJSON   Format = "json"
	FormatRaw    Format = "raw"
)

// Config holds all runtime options for a logslice invocation.
type Config struct {
	// Filter options
	Level    string
	MsgRegex string
	Since    time.Time
	Until    time.Time

	// Output options
	OutputFormat Format
	NoColor      bool
	NoTimestamp  bool

	// Reader options
	BufSize int
}

// Default returns a Config populated with sensible defaults.
func Default() Config {
	return Config{
		OutputFormat: FormatPretty,
		BufSize:      64 * 1024,
	}
}

// Validate checks that the Config fields are consistent and valid.
func (c *Config) Validate() error {
	switch c.OutputFormat {
	case FormatPretty, FormatJSON, FormatRaw:
		// valid
	case "":
		c.OutputFormat = FormatPretty
	default:
		return errors.New("config: unknown output format " + string(c.OutputFormat))
	}

	if c.BufSize <= 0 {
		return errors.New("config: buf_size must be positive")
	}

	if !c.Since.IsZero() && !c.Until.IsZero() && c.Until.Before(c.Since) {
		return errors.New("config: until must not be before since")
	}

	return nil
}
