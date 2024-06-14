package bigcache

import (
	"context"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(1*time.Minute))
	require.NoError(t, err)
	defer cache.Close()

	err = cache.Set("key", []byte("value"))
	require.NoError(t, err)

	entry, err := cache.Get("key")
	require.NoError(t, err)
	assert.Equal(t, []byte("value"), entry)
}

func TestDelete(t *testing.T) {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(1*time.Minute))
	require.NoError(t, err)
	defer cache.Close()

	err = cache.Set("key", []byte("value"))
	require.NoError(t, err)

	err = cache.Delete("key")
	require.NoError(t, err)

	_, err = cache.Get("key")
	assert.ErrorIs(t, err, bigcache.ErrEntryNotFound)
}
