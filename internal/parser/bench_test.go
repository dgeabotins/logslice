package parser_test

import (
	"fmt"
	"testing"

	"github.com/logslice/logslice/internal/parser"
)

var sink *parser.LogEntry

func BenchmarkParse_JSON(b *testing.B) {
	line := `{"level":"info","msg":"request handled","time":"2024-03-01T12:00:00Z","method":"GET","path":"/api/v1/pods","status":200,"latency_ms":12}`
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e, _ := parser.Parse(line)
		sink = e
	}
}

func BenchmarkParse_PlainText(b *testing.B) {
	line := "E0301 12:00:00.000000       1 reflector.go:147] k8s.io/client-go: failed to list"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e, _ := parser.Parse(line)
		sink = e
	}
}

func BenchmarkParse_ManyFields(b *testing.B) {
	fields := `"f1":"v1","f2":"v2","f3":"v3","f4":"v4","f5":"v5"`
	line := fmt.Sprintf(`{"level":"warn","msg":"many fields",%s}`, fields)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e, _ := parser.Parse(line)
		sink = e
	}
}
