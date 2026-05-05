// Command logslice is a fast structured log filter and formatter
// for JSON log streams from Kubernetes pods.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/yourorg/logslice/internal/config"
	"github.com/yourorg/logslice/internal/filter"
	"github.com/yourorg/logslice/internal/formatter"
	"github.com/yourorg/logslice/internal/output"
	"github.com/yourorg/logslice/internal/parser"
	"github.com/yourorg/logslice/internal/reader"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "logslice: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := parseFlags()
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	r := reader.New(os.Stdin, cfg.BufSize)
	f, err := filter.New(cfg)
	if err != nil {
		return fmt.Errorf("filter: %w", err)
	}
	fmt := formatter.New(cfg)
	out := output.New(os.Stdout)

	for line := range r.Lines(ctx) {
		entry := parser.Parse(line)
		if !f.Match(entry) {
			continue
		}
		formatted, fErr := fmt.Format(entry)
		if fErr != nil {
			return fmt.Errorf("format: %w", fErr)
		}
		if wErr := out.WriteLine(formatted); wErr != nil {
			return fmt.Errorf("write: %w", wErr)
		}
	}

	return r.Err()
}
