package middleware

import (
	"gin-app-start/pkg/response"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	rate       int
	lastAccess map[string]time.Time
	tokens     map[string]int
	mu         sync.Mutex
}

func newRateLimiter(rate int) *rateLimiter {
	limiter := &rateLimiter{
		rate:       rate,
		lastAccess: make(map[string]time.Time),
		tokens:     make(map[string]int),
	}

	go limiter.cleanup()

	return limiter
}

func (rl *rateLimiter) allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	lastTime, exists := rl.lastAccess[key]

	if !exists {
		rl.lastAccess[key] = now
		rl.tokens[key] = rl.rate - 1
		return true
	}

	elapsed := now.Sub(lastTime).Seconds()
	tokensToAdd := int(elapsed * float64(rl.rate))

	if tokensToAdd > 0 {
		rl.tokens[key] += tokensToAdd
		if rl.tokens[key] > rl.rate {
			rl.tokens[key] = rl.rate
		}
		rl.lastAccess[key] = now
	}

	if rl.tokens[key] > 0 {
		rl.tokens[key]--
		return true
	}

	return false
}

func (rl *rateLimiter) cleanup() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for key, lastTime := range rl.lastAccess {
			if now.Sub(lastTime) > 5*time.Minute {
				delete(rl.lastAccess, key)
				delete(rl.tokens, key)
			}
		}
		rl.mu.Unlock()
	}
}

var globalLimiter *rateLimiter

func RateLimit(rate int) gin.HandlerFunc {
	if globalLimiter == nil {
		globalLimiter = newRateLimiter(rate)
	}

	return func(c *gin.Context) {
		key := c.ClientIP()

		if !globalLimiter.allow(key) {
			response.Error(c, 42900, "Too many requests, please try again later")
			c.Abort()
			return
		}

		c.Next()
	}
}
