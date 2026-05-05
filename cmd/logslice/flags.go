package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/yourorg/logslice/internal/config"
)

func parseFlags() (*config.Config, error) {
	cfg := config.Default()

	flag.StringVar(&cfg.Level, "level", cfg.Level, "minimum log level to show (debug|info|warn|error)")
	flag.StringVar(&cfg.MsgFilter, "msg", cfg.MsgFilter, "regex filter applied to the message field")
	flag.StringVar(&cfg.Format, "format", cfg.Format, "output format: pretty|json|raw")
	flag.BoolVar(&cfg.Color, "color", cfg.Color, "enable ANSI color in pretty output")
	flag.BoolVar(&cfg.NoTimestamp, "no-ts", cfg.NoTimestamp, "omit timestamp from output")
	flag.IntVar(&cfg.BufSize, "bufsize", cfg.BufSize, "scanner buffer size in bytes")

	var since, until string
	flag.StringVar(&since, "since", "", "show entries at or after this time (RFC3339)")
	flag.StringVar(&until, "until", "", "show entries at or before this time (RFC3339)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: logslice [flags]\n\nFlags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n  kubectl logs my-pod | logslice -level=error -format=pretty\n")
	}

	flag.Parse()

	if since != ""	{
		t, err := time.Parse(time.RFC3339, since)
		if err != nil {
			return nil, fmt.Errorf("invalid -since value %q: %w", since, err)
		}
		cfg.Since = &t
	}
	if until != "" {
		t, err := time.Parse(time.RFC3339, until)
		if err != nil {
			return nil, fmt.Errorf("invalid -until value %q: %w", until, err)
		}
		cfg.Until = &t
	}

	// Allow extra args as a convenience: logslice error → -level=error
	if flag.NArg() == 1 {
		arg := strings.ToLower(flag.Arg(0))
		for _, lvl := range []string{"debug", "info", "warn", "error"} {
			if arg == lvl {
				cfg.Level = arg
				break
			}
		}
	}

	return cfg, nil
}
