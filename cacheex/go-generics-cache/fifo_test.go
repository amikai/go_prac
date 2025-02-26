package cache

import (
	"testing"

	cache "github.com/Code-Hex/go-generics-cache"
	"github.com/Code-Hex/go-generics-cache/policy/fifo"
	"github.com/stretchr/testify/assert"
)

func TestFIFO(t *testing.T) {
	c := cache.NewContext(t.Context(), cache.AsFIFO[string, int](fifo.WithCapacity(3)))

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

	c.Set("d", 4)
	dv, dok := c.Get("d")
	assert.True(t, dok)
	assert.Equal(t, 4, dv)

	// a is clean
	av, aok = c.Get("a")
	assert.False(t, aok)
	assert.Equal(t, 0, av)
}
