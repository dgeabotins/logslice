package formatter_test

import (
	"io"
	"testing"
	"time"

	"github.com/logslice/logslice/internal/formatter"
	"github.com/logslice/logslice/internal/parser"
)

var benchEntry = parser.Entry{
	Timestamp: time.Now(),
	Level:     "info",
	Message:   "request completed",
	Fields: map[string]any{
		"method":  "GET",
		"path":    "/api/v1/users",
		"status":  200,
		"latency": "12ms",
	},
	Raw: `{"level":"info","msg":"request completed","method":"GET"}`,
}

func BenchmarkFormatter_Pretty(b *testing.B) {
	f := formatter.New(io.Discard, formatter.FormatPretty, false, time.RFC3339)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Write(benchEntry)
	}
}

func BenchmarkFormatter_JSON(b *testing.B) {
	f := formatter.New(io.Discard, formatter.FormatJSON, false, time.RFC3339)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Write(benchEntry)
	}
}

func BenchmarkFormatter_Raw(b *testing.B) {
	f := formatter.New(io.Discard, formatter.FormatRaw, false, "")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Write(benchEntry)
	}
}

func BenchmarkFormatter_PrettyColor(b *testing.B) {
	f := formatter.New(io.Discard, formatter.FormatPretty, true, time.RFC3339)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Write(benchEntry)
	}
}
