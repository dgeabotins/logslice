package sampler_test

import (
	"testing"

	"github.com/yourorg/logslice/internal/parser"
	"github.com/yourorg/logslice/internal/sampler"
)

func entry() parser.Entry { return parser.Entry{Message: "test"} }

func TestNew_ZeroEveryCoercedToOne(t *testing.T) {
	s := sampler.New(0, 0)
	if s.Every != 1 {
		t.Fatalf("expected Every=1, got %d", s.Every)
	}
}

func TestKeep_BurstPassesAll(t *testing.T) {
	s := sampler.New(5, 2)
	for i := 0; i < 5; i++ {
		if !s.Keep(entry()) {
			t.Fatalf("entry %d should pass during burst", i)
		}
	}
}

func TestKeep_EveryNthAfterBurst(t *testing.T) {
	s := sampler.New(2, 3)
	// consume burst
	s.Keep(entry())
	s.Keep(entry())

	// after burst: keep=true at offset 1,4,7,...
	want := []bool{true, false, false, true, false, false, true}
	for i, w := range want {
		got := s.Keep(entry())
		if got != w {
			t.Errorf("post-burst entry %d: want %v got %v", i, w, got)
		}
	}
}

func TestKeep_Every1KeepsAll(t *testing.T) {
	s := sampler.New(0, 1)
	for i := 0; i < 20; i++ {
		if !s.Keep(entry()) {
			t.Fatalf("entry %d should always pass with Every=1", i)
		}
	}
}

func TestReset(t *testing.T) {
	s := sampler.New(2, 10)
	for i := 0; i < 5; i++ {
		s.Keep(entry())
	}
	s.Reset()
	// after reset first two entries should pass (burst)
	if !s.Keep(entry()) {
		t.Fatal("first entry after reset should pass")
	}
	if !s.Keep(entry()) {
		t.Fatal("second entry after reset should pass")
	}
}
