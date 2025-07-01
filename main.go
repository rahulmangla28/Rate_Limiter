package main

import (
	"RATE_LIMITER/fixedWindow"
	"RATE_LIMITER/leakyBucket"
	"RATE_LIMITER/slidingWindow"
	"RATE_LIMITER/tokenBucket"
)

func main() {
	// =============== Implements Token Bucket ====================
	tokenBucket.TestTokenBucketLimiter()

	// =============== Implements Leaky Bucket ====================
	leakyBucket.TestLeakyBucketLimiter()

	// =============== Implements Fixed Window ====================
	fixedWindow.TestFixedWindowLimiter()

	// =============== Implements Sliding Window ====================
	slidingWindow.TestSlidingWindowLimiter()
}
