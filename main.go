package bench

// Make glide import our packages, until glide/#461 gets in

import (
	res "github.com/eapache/go-resiliency/breaker"
	rubyist "github.com/rubyist/circuitbreaker"
	sony "github.com/sony/gobreaker"
)

var _ = rubyist.Breaker{}
var _ = sony.CircuitBreaker{}
var _ = res.Breaker{}
