package tokenBucket

import (
	"fmt"
	"net/http"
)

func TestTokenBucketLimiter() {
	// Init Limiter
	ipRateLimiter := newIPRateLimiter()

	mux := http.NewServeMux()
	mux.HandleFunc("/", RateLimitMiddleware(ipRateLimiter, handleRequest))

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", mux)
}
