package fixedWindow

import (
	"fmt"
	"time"
)

func TestFixedWindowLimiter() {
	// Create a new rate limiter: 5 requests allowed every 10 seconds
	rl := NewRateLimiter(5, 10*time.Second)

	clientID := "client1"

	for i := 1; i <= 10; i++ {
		if rl.allowRequest(clientID) {
			fmt.Printf("Request %d allowed for %s\n", i, clientID)
		} else {
			fmt.Printf("Request %d denied for %s (Rate limit exceeded)\n", i, clientID)
		}
		time.Sleep(1 * time.Second)
	}

	// Wait and demonstrate that the window resets
	fmt.Println("\n Waiting 11 seconds to reset the window...\n")
	time.Sleep(11 * time.Second)

	for i := 1; i <= 5; i++ {
		if rl.allowRequest(clientID) {
			fmt.Printf("Request %d allowed for %s after reset\n", i, clientID)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
