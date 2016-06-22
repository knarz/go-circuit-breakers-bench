package bench

import (
	"testing"

	circuit "github.com/rubyist/circuitbreaker"
)

func BenchmarkRubyistCbClosedNoOp(b *testing.B) {
	cb := circuit.NewThresholdBreaker(10)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cb.Call(func() error {
				return nil
			}, 0)
		}
	})
}
