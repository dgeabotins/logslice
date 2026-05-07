package sampler_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/parser"
	"github.com/yourorg/logslice/internal/sampler"
)

func BenchmarkSampler_Keep_BurstOnly(b *testing.B) {
	e := parser.Entry{Message: "benchmark log line"}
	s := sampler.New(uint64(b.N)+1, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Keep(e)
	}
}

func BenchmarkSampler_Keep_Every10(b *testing.B) {
	e := parser.Entry{Message: "benchmark log line"}
	s := sampler.New(0, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Keep(e)
	}
}

func BenchmarkSampler_Keep_Parallel(b *testing.B) {
	e := parser.Entry{Message: "benchmark log line"}
	s := sampler.New(100, 5)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Keep(e)
		}
	})
}
