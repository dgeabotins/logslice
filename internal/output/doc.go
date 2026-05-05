// Package output provides a thread-safe writer for logslice output streams.
//
// It wraps any io.Writer with mutex-guarded writes and optional automatic
// newline appending, making it safe to use from multiple goroutines — for
// example when tailing logs from multiple Kubernetes pods concurrently.
//
// Basic usage:
//
//	w := output.New(os.Stdout, output.WithNewline())
//	w.WriteString("formatted log line")
//
// The Writer satisfies io.Writer so it can be passed directly to formatters
// and other components that accept an io.Writer.
package output
