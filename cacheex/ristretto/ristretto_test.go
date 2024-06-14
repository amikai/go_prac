package ristretto

import (
	"testing"

	"github.com/dgraph-io/ristretto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetGetDel(t *testing.T) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e4,     // number of keys to track frequency of (10K).
		MaxCost:     1 << 20, // maximum cost of cache (1M).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	require.NoError(t, err)

	cache.Set("key", "value", 1)
	// wait for value to pass through buffers
	cache.Wait()

	// get value from cache
	value, found := cache.Get("key")
	assert.True(t, found)
	assert.Equal(t, "value", value)

	cache.Del("key")
	_, found = cache.Get("key")
	assert.False(t, found)
}
