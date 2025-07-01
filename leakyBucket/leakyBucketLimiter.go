package leakyBucket

import (
	"fmt"
	"sync"
	"time"
)

// Request represents an incoming task with a unique ID
type Request struct {
	ID int
}

// leakyBucket implements a simple leaky bucket rate limiter
type LeakyBucket struct {
	queue        []Request     // The bucket holding requests
	capacity     int           // Max capacity of the bucket
	emptyRate    time.Duration // Rate at which requests are processed (leak)
	stopRefiller chan struct{} // Channel to stop the leak goroutine
	mu           sync.Mutex    // Mutex for safe concurrent access
}

// NewLeakyBucket creates and starts a new LeakyBucket
func NewLeakyBucket(capacity int, emptyRate time.Duration) *LeakyBucket {
	lb := &LeakyBucket{
		queue:        make([]Request, 0, capacity),
		capacity:     capacity,
		emptyRate:    emptyRate,
		stopRefiller: make(chan struct{}),
	}
	go lb.removeRequestsFromQueue()
	return lb
}

// AddRequest adds a request to the bucket or throttles it if full
func (lb *LeakyBucket) AddRequest(req Request) bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if len(lb.queue) < lb.capacity {
		lb.queue = append(lb.queue, req)
		fmt.Printf("Added Request ID: %d\n", req.ID)
		return true
	}

	fmt.Printf("Throttled Request ID: %d (Bucket Full)\n", req.ID)
	return false
}

// removeRequestsFromQueue leaks requests from the bucket at a fixed rate
func (lb *LeakyBucket) removeRequestsFromQueue() {
	ticker := time.NewTicker(lb.emptyRate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			lb.mu.Lock()
			if len(lb.queue) > 0 {
				for _, req := range lb.queue {
					fmt.Printf("Processing Request ID: %d\n", req.ID)
				}
				lb.queue = lb.queue[:0]
			}
			lb.mu.Unlock()

		case <-lb.stopRefiller:
			fmt.Println("Stopping request processing.")
			return
		}
	}
}

// Stop gracefully shuts down the request processing goroutine
func (lb *LeakyBucket) Stop() {
	lb.stopRefiller <- struct{}{}
}
