package slidingWindow

import (
	"fmt"
	"time"
)

func TestSlidingWindowLimiter() {
	// Configure rate limiter: 10 requests per second, over a 1-minute window
	rate := 10.0
	window := 1 * time.Minute

	limiter := NewConfigurableSlidingWindowRateLimiter(rate, window)

	// Simulate 100 requests
	for i := 0; i < 100; i++ {
		if limiter.Allow() {
			fmt.Printf("[%02d] Request allowed\n", i+1)
		} else {
			fmt.Printf("[%02d] Request rejected\n", i+1)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
