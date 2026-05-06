package highlight_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/highlight"
)

func BenchmarkLevel_Error(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = highlight.Level("error")
	}
}

func BenchmarkLevel_Info(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = highlight.Level("info")
	}
}

func BenchmarkTimestamp(b *testing.B) {
	ts := "2024-01-02T15:04:05.000Z"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = highlight.Timestamp(ts)
	}
}

func BenchmarkMessage(b *testing.B) {
	msg := "request completed successfully"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = highlight.Message(msg)
	}
}

func BenchmarkStrip(b *testing.B) {
	colored := highlight.Level("error") + " " + highlight.Message("something bad")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = highlight.Strip(colored)
	}
}
