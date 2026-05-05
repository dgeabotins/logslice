// Package reader provides a line-oriented streaming reader for log sources.
//
// It wraps an io.Reader and emits lines over a channel, making it easy to
// integrate with pipeline stages such as the parser and filter packages.
//
// Basic usage:
//
//	rd := reader.New(reader.Options{})
//	for line := range rd.Lines(ctx, os.Stdin) {
//		entry, err := parser.Parse(line)
//		// ...
//	}
//
// Follow mode:
//
// When Options.Follow is true the reader does not stop at EOF, mimicking
// "tail -f" behaviour. The caller must cancel the context to terminate the
// stream.
//
// Buffer size:
//
// Lines longer than BufSize bytes are skipped by the underlying bufio.Scanner.
// Increase Options.BufSize when processing logs with very large payloads.
package reader
