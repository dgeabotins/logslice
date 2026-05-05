package output

import (
	"io"
	"os"
	"sync"
)

// Writer wraps an io.Writer with optional buffering and thread-safe writes.
type Writer struct {
	mu  sync.Mutex
	w   io.Writer
	nl  bool // append newline if missing
}

// Option configures a Writer.
type Option func(*Writer)

// WithNewline ensures each written line ends with a newline character.
func WithNewline() Option {
	return func(w *Writer) { w.nl = true }
}

// New creates a Writer wrapping the given io.Writer.
// If w is nil, os.Stdout is used.
func New(w io.Writer, opts ...Option) *Writer {
	if w == nil {
		w = os.Stdout
	}
	out := &Writer{w: w}
	for _, o := range opts {
		o(out)
	}
	return out
}

// Write writes p to the underlying writer in a thread-safe manner.
func (w *Writer) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.nl && len(p) > 0 && p[len(p)-1] != '\n' {
		buf := make([]byte, len(p)+1)
		copy(buf, p)
		buf[len(p)] = '\n'
		return w.w.Write(buf)
	}
	return w.w.Write(p)
}

// WriteLine writes a single line, always appending a newline.
func (w *Writer) WriteLine(line []byte) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	var err error
	if len(line) > 0 && line[len(line)-1] == '\n' {
		_, err = w.w.Write(line)
	} else {
		buf := make([]byte, len(line)+1)
		copy(buf, line)
		buf[len(line)] = '\n'
		_, err = w.w.Write(buf)
	}
	return err
}

// WriteString writes a string line.
func (w *Writer) WriteString(s string) error {
	return w.WriteLine([]byte(s))
}
