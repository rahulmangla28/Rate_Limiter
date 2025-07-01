package fixedWindow

import (
	"sync"
	"time"
)

// RateLimiter is the core structure for tracking requests per client.
type rateLimiter struct {
	requests map[string]int // Track request counts per client
	limit    int            // Max requests allowed per window
	window   time.Duration  // Time window duration
	mutex    sync.Mutex     // Mutex to ensure thread-safe operations
}

// NewRateLimiter creates and returns a new RateLimiter instance.
func NewRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		requests: make(map[string]int),
		limit:    limit,
		window:   window,
	}
}

// AllowRequest checks if a request from the given clientID is allowed.
// Returns true if allowed, false if rate limit is exceeded.
func (rl *rateLimiter) allowRequest(clientID string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	count, exists := rl.requests[clientID]

	if !exists || count < rl.limit {
		// Start a reset timer only once when the client first hits the map
		if !exists {
			go rl.resetCount(clientID)
		}
		rl.requests[clientID]++
		return true
	}
	return false
}

// resetCount clears the request count for a client after the time window.
func (rl *rateLimiter) resetCount(clientID string) {
	time.Sleep(rl.window)
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	delete(rl.requests, clientID)
}
