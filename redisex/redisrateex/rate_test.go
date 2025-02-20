package redisrateex

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRateLimiter(t *testing.T) {
	s := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	err := rdb.FlushDB(t.Context()).Err()
	require.NoError(t, err)

	rateLimiter := redis_rate.NewLimiter(rdb)
	limit := redis_rate.PerHour(1)

	res, err := rateLimiter.Allow(t.Context(), "test_key", limit)
	require.NoError(t, err)
	assert.Equal(t, 1, res.Allowed)
	assert.Equal(t, 0, res.Remaining)
	assert.Equal(t, time.Duration(-1), res.RetryAfter)

	res, err = rateLimiter.Allow(t.Context(), "test_key", limit)
	require.NoError(t, err)
	assert.Equal(t, 0, res.Allowed)
	assert.Equal(t, 0, res.Remaining)
	assert.InDelta(t, res.RetryAfter, time.Hour, float64(time.Second))
}
