package ratelimiter

import (
	"time"
)

type RateLimiter struct {
	requests chan int
	ticker   *time.Ticker
}

func NewRateLimiter(rate int, per time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(chan int, rate),
		ticker:   time.NewTicker(per),
	}

	go rl.run()

	return rl
}

func (rl *RateLimiter) run() {
	for range rl.ticker.C {
		for i := 0; i < cap(rl.requests); i++ {
			select {
			case rl.requests <- 1:
			default:
			}
		}
	}
}

func (rl *RateLimiter) Wait() {
	<-rl.requests
}

func (rl *RateLimiter) Stop() {
	rl.ticker.Stop()
}
