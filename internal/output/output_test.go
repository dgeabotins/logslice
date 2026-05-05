package output_test

import (
	"bytes"
	"sync"
	"testing"

	"github.com/yourorg/logslice/internal/output"
)

func TestWriter_Write_AppendsNewline(t *testing.T) {
	var buf bytes.Buffer
	w := output.New(&buf, output.WithNewline())
	_, err := w.Write([]byte("hello"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := buf.String(); got != "hello\n" {
		t.Errorf("expected %q, got %q", "hello\n", got)
	}
}

func TestWriter_Write_NoDoubleNewline(t *testing.T) {
	var buf bytes.Buffer
	w := output.New(&buf, output.WithNewline())
	_, err := w.Write([]byte("hello\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := buf.String(); got != "hello\n" {
		t.Errorf("expected %q, got %q", "hello\n", got)
	}
}

func TestWriter_WriteLine(t *testing.T) {
	var buf bytes.Buffer
	w := output.New(&buf)
	if err := w.WriteLine([]byte("world")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := buf.String(); got != "world\n" {
		t.Errorf("expected %q, got %q", "world\n", got)
	}
}

func TestWriter_WriteString(t *testing.T) {
	var buf bytes.Buffer
	w := output.New(&buf)
	if err := w.WriteString("line"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := buf.String(); got != "line\n" {
		t.Errorf("expected %q, got %q", "line\n", got)
	}
}

func TestWriter_NilUsesStdout(t *testing.T) {
	// Just ensure no panic when nil is passed.
	w := output.New(nil)
	if w == nil {
		t.Fatal("expected non-nil writer")
	}
}

func TestWriter_ConcurrentWrites(t *testing.T) {
	var buf bytes.Buffer
	w := output.New(&buf)
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = w.WriteString("concurrent line")
		}()
	}
	wg.Wait()
}
