package output

import "io"

// MultiWriter fans out writes to multiple Writers in order.
// All targets receive every write; the first error encountered is returned.
type MultiWriter struct {
	writers []*Writer
}

// NewMultiWriter creates a MultiWriter that writes to all provided Writers.
func NewMultiWriter(writers ...*Writer) *MultiWriter {
	return &MultiWriter{writers: writers}
}

// Write writes p to all underlying Writers.
func (m *MultiWriter) Write(p []byte) (int, error) {
	for _, w := range m.writers {
		if _, err := w.Write(p); err != nil {
			return 0, err
		}
	}
	return len(p), nil
}

// WriteLine writes a line to all underlying Writers.
func (m *MultiWriter) WriteLine(line []byte) error {
	for _, w := range m.writers {
		if err := w.WriteLine(line); err != nil {
			return err
		}
	}
	return nil
}

// Unwrap returns the underlying io.Writer slice for compatibility.
func (m *MultiWriter) Unwrap() []io.Writer {
	out := make([]io.Writer, len(m.writers))
	for i, w := range m.writers {
		out[i] = w
	}
	return out
}
