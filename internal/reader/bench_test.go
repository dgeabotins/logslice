package reader_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/yourorg/logslice/internal/reader"
)

func BenchmarkReader_Lines(b *testing.B) {
	const lineCount = 1000
	var sb strings.Builder
	for i := 0; i < lineCount; i++ {
		fmt.Fprintf(&sb, `{"level":"info","msg":"request handled","ts":"2024-01-15T10:00:00Z","id":%d}`+"\n", i)
	}
	input := sb.String()

	rd := reader.New(reader.Options{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ch := rd.Lines(ctx, strings.NewReader(input))
		for range ch {
		}
	}
	b.SetBytes(int64(len(input)))
}

func BenchmarkReader_LargeLines(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < 100; i++ {
		fmt.Fprintf(&sb, `{"level":"debug","msg":"%s","ts":"2024-01-15T10:00:00Z"}`+"\n",
			strings.Repeat("x", 512))
	}
	input := sb.String()

	rd := reader.New(reader.Options{BufSize: 128 * 1024})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ch := rd.Lines(ctx, strings.NewReader(input))
		for range ch {
		}
	}
	b.SetBytes(int64(len(input)))
}
