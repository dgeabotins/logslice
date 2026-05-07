package sampler

import (
	"sync/atomic"

	"github.com/yourorg/logslice/internal/parser"
)

// Sampler drops log entries that exceed a given rate, keeping every Nth
// matching entry after the burst limit is reached.
type Sampler struct {
	// Burst is the number of entries to pass through before sampling begins.
	Burst uint64
	// Every controls which entries are kept once burst is exhausted:
	// 1 means keep all, 2 means keep every 2nd, etc.
	Every uint64

	counter atomic.Uint64
}

// New returns a Sampler with the provided burst and every values.
// If every is 0 it is coerced to 1 (keep all after burst).
func New(burst, every uint64) *Sampler {
	if every == 0 {
		every = 1
	}
	return &Sampler{Burst: burst, Every: every}
}

// Keep reports whether the given entry should be forwarded downstream.
// It is safe to call from multiple goroutines.
func (s *Sampler) Keep(_ parser.Entry) bool {
	n := s.counter.Add(1) // 1-based
	if n <= s.Burst {
		return true
	}
	offset := n - s.Burst
	return offset%s.Every == 1
}

// Reset sets the internal counter back to zero.
func (s *Sampler) Reset() {
	s.counter.Store(0)
}
