package filter_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/filter"
)

var sink bool

func BenchmarkFilter_Match_AllCriteria(b *testing.B) {
	f, _ := filter.New(filter.Options{
		Levels:     []string{"error", "warn"},
		MsgPattern: `timeout|refused`,
		Since:      time.Now().Add(-time.Hour),
		Until:      time.Now().Add(time.Hour),
		Fields:     map[string]string{"namespace": "prod", "pod": "api-server-1"},
	})
	e := filter.Entry{
		Timestamp: time.Now(),
		Level:     "error",
		Message:   "connection timeout to database",
		Fields:    map[string]string{"namespace": "prod", "pod": "api-server-1"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = f.Match(e)
	}
}

func BenchmarkFilter_Match_LevelOnly(b *testing.B) {
	f, _ := filter.New(filter.Options{Levels: []string{"error"}})
	e := filter.Entry{Level: "error", Message: "boom", Fields: map[string]string{}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = f.Match(e)
	}
}

func BenchmarkFilter_Match_NoMatch(b *testing.B) {
	f, _ := filter.New(filter.Options{
		Levels:     []string{"error"},
		MsgPattern: `critical`,
	})
	e := filter.Entry{Level: "info", Message: "all systems nominal", Fields: map[string]string{}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = f.Match(e)
	}
}
