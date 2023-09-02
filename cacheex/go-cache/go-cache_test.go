package gocache

import (
	"sync"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

// https://pkg.go.dev/github.com/patrickmn/go-cache

func testConcGet[T any](t *testing.T, c *cache.Cache, key string, wantVal T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 32; i++ {
		wg.Add(1)
		func() {
			got, ok := c.Get(key)
			assert.True(t, ok)
			assert.Equal(t, wantVal, got.(T))
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestCacheSimpleUsage(t *testing.T) {
	c := cache.New(cache.NoExpiration, cache.DefaultExpiration)
	c.Set("zero", 0, cache.DefaultExpiration)
	c.Set("one", 1, cache.DefaultExpiration)
	c.Set("two", 2, cache.DefaultExpiration)

	testConcGet(t, c, "zero", 0)
	testConcGet(t, c, "one", 1)
	testConcGet(t, c, "two", 2)

	assert.Equal(t, 3, c.ItemCount())
}

func TestCacheExpired(t *testing.T) {
	c := cache.New(100*time.Millisecond, cache.DefaultExpiration)
	c.Set("zero", 0, cache.DefaultExpiration)
	c.Set("one", 1, cache.DefaultExpiration)
	c.Set("two", 2, cache.DefaultExpiration)

	testConcGet(t, c, "zero", 0)
	testConcGet(t, c, "one", 1)
	testConcGet(t, c, "two", 2)

	time.Sleep(200 * time.Millisecond)
	// alread expired
	_, ok := c.Get("zero")
	assert.False(t, ok)
	_, ok = c.Get("one")
	assert.False(t, ok)
	_, ok = c.Get("two")
	assert.False(t, ok)
}
