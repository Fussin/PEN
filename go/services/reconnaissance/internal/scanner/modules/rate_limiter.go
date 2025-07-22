package modules

import (
	"sync"
	"time"
)

type RateLimiter struct {
	tokens        int
	maxTokens     int
	refillRate    int
	refillTicker  *time.Ticker
	mtx           sync.Mutex
	refillStopped chan bool
}

func NewRateLimiter(rate int, interval time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		tokens:        rate,
		maxTokens:     rate,
		refillRate:    rate,
		refillTicker:  time.NewTicker(interval),
		refillStopped: make(chan bool),
	}

	go limiter.refill()
	return limiter
}

func (r *RateLimiter) refill() {
	for {
		select {
		case <-r.refillTicker.C:
			r.mtx.Lock()
			r.tokens = r.maxTokens
			r.mtx.Unlock()
		case <-r.refillStopped:
			return
		}
	}
}

func (r *RateLimiter) Wait() {
	for {
		r.mtx.Lock()
		if r.tokens > 0 {
			r.tokens--
			r.mtx.Unlock()
			return
		}
		r.mtx.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func (r *RateLimiter) Stop() {
	r.refillTicker.Stop()
	close(r.refillStopped)
}
