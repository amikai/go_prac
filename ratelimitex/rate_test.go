package ratelimitex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

// Note: golang rate limiter is token bucket rate limiter

func TestLimiterAllow(t *testing.T) {
	refillRate := rate.Every(1 * time.Second)
	burst := 1
	limiter := rate.NewLimiter(refillRate, burst)

	assert.Equal(t, refillRate, limiter.Limit())
	assert.Equal(t, burst, limiter.Burst())

	assert.True(t, limiter.Allow())
	assert.False(t, limiter.Allow())
	assert.False(t, limiter.Allow())
	assert.False(t, limiter.Allow())
}

func TestLimiterWait(t *testing.T) {
	refillRate := rate.Every(250 * time.Millisecond)
	burst := 1
	limiter := rate.NewLimiter(refillRate, burst)

	assert.True(t, limiter.Allow())

	ctx := t.Context()
	startTime := time.Now()
	for i := 0; i < 4; i++ {
		err := limiter.Wait(ctx)
		assert.NoError(t, err)
	}
	assert.WithinDuration(t, startTime.Add(4*250*time.Millisecond), time.Now(), 30*time.Millisecond)
}

func TestLimiterReserve(t *testing.T) {
	t.Skip("TODO")
}
