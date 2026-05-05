package output_test

import (
	"io"
	"testing"

	"github.com/yourorg/logslice/internal/output"
)

var line = []byte(`{"level":"info","msg":"request handled","ts":"2024-01-01T00:00:00Z"}`)

func BenchmarkWriter_WriteLine(b *testing.B) {
	w := output.New(io.Discard)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = w.WriteLine(line)
	}
}

func BenchmarkWriter_WriteString(b *testing.B) {
	w := output.New(io.Discard)
	s := string(line)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = w.WriteString(s)
	}
}

func BenchmarkWriter_Write_WithNewline(b *testing.B) {
	w := output.New(io.Discard, output.WithNewline())
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = w.Write(line)
	}
}
