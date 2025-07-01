package slidingWindow

import (
	"sync"
	"time"
)

// type limiter interface {
// 	SetRate(rate float64)
// 	SetWindow(window time.Duration)
// 	Allow() bool
// }

// slidingWindow implements a sliding window rate limiting algorithm.
type SlidingWindow struct {
	mu      sync.Mutex    // Ensures thread-safe access
	rate    int           // Max allowed requests in the window
	window  time.Duration // Window size
	history []time.Time   // Slice of timestamps for requests
}

// increment records a new request timestamp.
func (sw *SlidingWindow) increment(now time.Time) {
	sw.history = append(sw.history, now)
}

// removeExpired removes request timestamps that fall outside the current window.
func (sw *SlidingWindow) removeExpired(now time.Time) {
	cutoff := now.Add(-sw.window)
	idx := 0
	for i, t := range sw.history {
		if t.After(cutoff) {
			idx = i
			break
		}
	}
	sw.history = sw.history[idx:]
}

// countRequests returns the number of requests within the current sliding window.
func (sw *SlidingWindow) countRequests(now time.Time) int {
	sw.removeExpired(now)
	return len(sw.history)
}

// Allow determines whether a new request can be allowed.
func (sw *SlidingWindow) allow() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now()
	if sw.countRequests(now) >= sw.rate {
		return false
	}
	sw.increment(now)
	return true
}

// SetRate updates the maximum allowed requests per window.
func (sw *SlidingWindow) setRate(rate float64) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	sw.rate = int(rate * float64(sw.window) / float64(time.Second))
}

// SetWindow updates the window duration.
func (sw *SlidingWindow) setWindow(window time.Duration) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	sw.window = window
	sw.setRate(float64(sw.rate)) // Recalculate rate using new window
}

// ConfigurableSlidingWindowRateLimiter wraps slidingWindow and implements RateLimiter interface.
type ConfigurableSlidingWindowRateLimiter struct {
	sw *SlidingWindow
}

// NewConfigurableSlidingWindowRateLimiter creates a new rate limiter with provided rate and window.
func NewConfigurableSlidingWindowRateLimiter(rate float64, window time.Duration) *ConfigurableSlidingWindowRateLimiter {
	sw := &SlidingWindow{
		window:  window,
		history: make([]time.Time, 0),
	}
	sw.setRate(rate)

	return &ConfigurableSlidingWindowRateLimiter{sw: sw}
}

// SetRate adjusts the rate dynamically.
func (rl *ConfigurableSlidingWindowRateLimiter) SetRate(rate float64) {
	rl.sw.setRate(rate)
}

// SetWindow adjusts the window size dynamically.
func (rl *ConfigurableSlidingWindowRateLimiter) SetWindow(window time.Duration) {
	rl.sw.setWindow(window)
}

// Allow checks if a new request is allowed under current limits.
func (rl *ConfigurableSlidingWindowRateLimiter) Allow() bool {
	return rl.sw.allow()
}
