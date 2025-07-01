package tokenBucket

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

// Define Rate Limiter
type RateLimiter struct {
	tokens         float64   // Current number of tokens
	maxTokens      float64   // Maximum tokens allowed
	refillRate     float64   // Tokens added per second
	lastRefillTime time.Time // Last time tokens were refilled
	mutex          sync.Mutex
}

// Constructor for Rate Limiter
func newRateLimiter(maxTokens, refillRate float64) *RateLimiter {
	return &RateLimiter{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

// Implements Token Refill
func (r *RateLimiter) refillTokens() {
	now := time.Now()
	duration := now.Sub(r.lastRefillTime).Seconds()
	tokensToAdd := duration * r.refillRate

	r.tokens += tokensToAdd
	if r.tokens > r.maxTokens {
		r.tokens = r.maxTokens
	}
	r.lastRefillTime = now
}

// Checks if request is allowed to process
func (r *RateLimiter) isReqAllowed() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.refillTokens()

	if r.tokens >= 1 {
		r.tokens -= 1
		return true
	}
	return false
}

// Define Rate Limiter for per IP
type IPRateLimiter struct {
	limiters map[string]*RateLimiter
	mutex    sync.Mutex
}

// Constructor for IP Rate Limiter
func newIPRateLimiter() *IPRateLimiter {
	return &IPRateLimiter{
		limiters: make(map[string]*RateLimiter),
	}
}

// Get rate kimiter for IP ip
func (i *IPRateLimiter) getRateLimiter(ip string) *RateLimiter {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	limiter, exists := i.limiters[ip]

	if !exists {
		limiter = newRateLimiter(3, 0.05)
		i.limiters[ip] = limiter
	}

	return limiter
}

// Middleware to process IP
func RateLimitMiddleware(ipRateLimiter *IPRateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Invalid IP", http.StatusInternalServerError)
			return
		}

		limiter := ipRateLimiter.getRateLimiter(ip)
		if limiter.isReqAllowed() {
			next(w, r)
		} else {
			http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
		}
	}
}

// Log request
func handleRequest(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Request processed successfully at %v\n", time.Now())
}
