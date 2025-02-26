package cache

import (
	"testing"

	cache "github.com/Code-Hex/go-generics-cache"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	c := cache.NewContext[string, int](t.Context())
	c.Set("a", 1)
	c.Set("b", 2)

	av, aok := c.Get("a")
	assert.True(t, aok)
	assert.Equal(t, 1, av)

	bv, bok := c.Get("b")
	assert.True(t, bok)
	assert.Equal(t, 2, bv)

	cv, cok := c.Get("c")
	assert.False(t, cok)
	assert.Equal(t, 0, cv)

	c.Delete("a")
	_, aok2 := c.Get("a")
	assert.False(t, aok2)

	c.Set("b", 3)
	newbv, newbok := c.Get("b")
	assert.True(t, newbok)
	assert.Equal(t, 3, newbv)
}
