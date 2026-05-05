package reader

import (
	"context"
	"io"
	"sync"
)

// MultiLines merges lines from multiple io.Readers into a single channel.
// Each reader is consumed concurrently; order of lines across readers is
// non-deterministic. The returned channel is closed once all readers are
// exhausted or the context is cancelled.
func (rd *Reader) MultiLines(ctx context.Context, readers ...io.Reader) <-chan string {
	out := make(chan string, 128*len(readers))
	var wg sync.WaitGroup
	for _, r := range readers {
		wg.Add(1)
		go func(r io.Reader) {
			defer wg.Done()
			for line := range rd.Lines(ctx, r) {
				select {
				case out <- line:
				case <-ctx.Done():
					return
				}
			}
		}(r)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
