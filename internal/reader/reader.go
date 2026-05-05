package reader

import (
	"bufio"
	"context"
	"io"
)

// DefaultBufSize is the default buffer size for reading lines.
const DefaultBufSize = 64 * 1024

// Options configures the Reader behaviour.
type Options struct {
	// BufSize overrides the default scanner buffer size.
	BufSize int
	// Follow keeps reading after EOF (tail -f behaviour).
	Follow bool
}

// Reader streams lines from an io.Reader to a channel.
type Reader struct {
	opts Options
}

// New creates a new Reader with the given options.
func New(opts Options) *Reader {
	if opts.BufSize <= 0 {
		opts.BufSize = DefaultBufSize
	}
	return &Reader{opts: opts}
}

// Lines reads lines from r and sends them to the returned channel.
// The channel is closed when the context is cancelled, an error occurs,
// or EOF is reached (unless Follow is set).
func (rd *Reader) Lines(ctx context.Context, r io.Reader) <-chan string {
	ch := make(chan string, 128)
	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(r)
		buf := make([]byte, rd.opts.BufSize)
		scanner.Buffer(buf, rd.opts.BufSize)
		for {
			for scanner.Scan() {
				line := scanner.Text()
				select {
				case ch <- line:
				case <-ctx.Done():
					return
				}
			}
			if err := scanner.Err(); err != nil {
				return
			}
			// EOF reached
			if !rd.opts.Follow {
				return
			}
			// Follow mode: wait for context cancellation
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}()
	return ch
}
