package golanglru

import (
	"testing"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/stretchr/testify/assert"
)

func TestExpirableLru(t *testing.T) {
	c := expirable.NewLRU[string, string](16, nil, 0)

	evicited := c.Add("key", "value")
	assert.False(t, evicited)

	got, ok := c.Get("key")
	assert.True(t, ok)
	assert.Equal(t, got, "value")
}

func TestExpirableLruExpire(t *testing.T) {
	c := expirable.NewLRU[string, string](16, nil, time.Microsecond*10)

	evicited := c.Add("key", "value")
	assert.False(t, evicited)

	// before expiration
	got, ok := c.Get("key")
	assert.True(t, ok)
	assert.Equal(t, got, "value")

	time.Sleep(time.Microsecond * 15)

	// after expiration
	_, ok = c.Get("key")
	assert.False(t, ok)
}
