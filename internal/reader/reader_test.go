package reader_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/reader"
)

func TestReader_BasicLines(t *testing.T) {
	input := "line1\nline2\nline3\n"
	rd := reader.New(reader.Options{})
	ctx := context.Background()
	ch := rd.Lines(ctx, strings.NewReader(input))

	var got []string
	for line := range ch {
		got = append(got, line)
	}
	if len(got) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(got))
	}
	for i, want := range []string{"line1", "line2", "line3"} {
		if got[i] != want {
			t.Errorf("line %d: got %q, want %q", i, got[i], want)
		}
	}
}

func TestReader_EmptyInput(t *testing.T) {
	rd := reader.New(reader.Options{})
	ctx := context.Background()
	ch := rd.Lines(ctx, strings.NewReader(""))
	var got []string
	for line := range ch {
		got = append(got, line)
	}
	if len(got) != 0 {
		t.Fatalf("expected 0 lines, got %d", len(got))
	}
}

func TestReader_ContextCancellation(t *testing.T) {
	// Build a large input so the goroutine blocks on channel send.
	var sb strings.Builder
	for i := 0; i < 10000; i++ {
		sb.WriteString("log line\n")
	}
	rd := reader.New(reader.Options{})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	ch := rd.Lines(ctx, strings.NewReader(sb.String()))
	// Drain until closed; must not hang.
	for range ch {
	}
}

func TestReader_CustomBufSize(t *testing.T) {
	input := strings.Repeat("x", 1024) + "\n"
	rd := reader.New(reader.Options{BufSize: 4096})
	ctx := context.Background()
	ch := rd.Lines(ctx, strings.NewReader(input))
	var count int
	for range ch {
		count++
	}
	if count != 1 {
		t.Fatalf("expected 1 line, got %d", count)
	}
}
