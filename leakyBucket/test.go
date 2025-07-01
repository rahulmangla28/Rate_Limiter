package leakyBucket

import (
	"fmt"
	"time"
)

func TestLeakyBucketLimiter() {
	// Create a leaky bucket with capacity 3 and leak interval of 1 second
	leakyBucket := NewLeakyBucket(3, 1*time.Second)

	fmt.Println("----- Sending requests every 200ms -----")
	for i := 1; i <= 10; i++ {
		leakyBucket.AddRequest(Request{ID: i})
		time.Sleep(200 * time.Millisecond)
	}

	// Allow time for processing
	time.Sleep(2 * time.Second)

	fmt.Println("----- Sending requests every 400ms -----")
	for i := 11; i <= 20; i++ {
		leakyBucket.AddRequest(Request{ID: i})
		time.Sleep(400 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)
	leakyBucket.Stop()
}
