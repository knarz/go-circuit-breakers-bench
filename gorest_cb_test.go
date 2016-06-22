package bench

import (
	"testing"
	"time"

	circuit "github.com/eapache/go-resiliency/breaker"
)

func BenchmarkGoResCbClosedNoOp(b *testing.B) {
	cb := circuit.New(5, 1, 5*time.Second)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cb.Run(func() error {
				return nil
			})
		}
	})
}
