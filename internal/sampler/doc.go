// Package sampler provides a token-bucket-style log entry sampler that limits
// the volume of entries forwarded to downstream consumers.
//
// # Overview
//
// When processing high-throughput Kubernetes pod logs it is useful to reduce
// noise by sampling repetitive entries.  The [Sampler] type lets a configurable
// number of entries (Burst) pass through unconditionally, then keeps every Nth
// subsequent entry (Every).
//
// # Usage
//
//	s := sampler.New(10, 5) // burst=10, then keep every 5th
//	for entry := range entries {
//		if s.Keep(entry) {
//			output.Write(entry)
//		}
//	}
//
// Sampler is safe for concurrent use.
package sampler
