package golanglru

import (
	"testing"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/stretchr/testify/assert"
)

func TestSimpleLru(t *testing.T) {
	c, err := lru.New[string, string](256)
	assert.NoError(t, err)

	evicited := c.Add("key", "value")
	assert.False(t, evicited)

	got, ok := c.Get("key")
	assert.True(t, ok)
	assert.Equal(t, got, "value")
}
