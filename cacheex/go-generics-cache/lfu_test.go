package cache

import (
	"testing"

	cache "github.com/Code-Hex/go-generics-cache"
	"github.com/Code-Hex/go-generics-cache/policy/lfu"
	"github.com/stretchr/testify/assert"
)

func TestLFU(t *testing.T) {
	c := cache.NewContext(t.Context(), cache.AsLFU[string, int](lfu.WithCapacity(3)))

	c.Set("a", 1)
	av, aok := c.Get("a")
	assert.True(t, aok)
	assert.Equal(t, 1, av)

	c.Set("b", 2)
	bv, bok := c.Get("b")
	assert.True(t, bok)
	assert.Equal(t, 2, bv)

	c.Set("c", 3)
	cv, cok := c.Get("c")
	assert.True(t, cok)
	assert.Equal(t, 3, cv)

	// use a three times
	_, _ = c.Get("a")
	_, _ = c.Get("a")
	_, _ = c.Get("a")
	// use b two times
	_, _ = c.Get("b")
	_, _ = c.Get("b")
	// use c one time
	_, _ = c.Get("c")

	// insert d, then one of the key will be evicted
	c.Set("d", 4)
	dv, dok := c.Get("d")
	assert.True(t, dok)
	assert.Equal(t, 4, dv)

	// c is clean
	cv, cok = c.Get("c")
	assert.False(t, cok)
	assert.Equal(t, 0, cv)
}
