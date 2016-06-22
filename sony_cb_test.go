package bench

import (
	"testing"

	circuit "github.com/sony/gobreaker"
)

func BenchmarkSonyCbClosedNoOp(b *testing.B) {
	cb := &circuit.CircuitBreaker{}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cb.Execute(func() (interface{}, error) {
				return nil, nil
			})
		}
	})
}
