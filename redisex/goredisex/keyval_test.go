package goredisex

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestGetSet(t *testing.T) {
	s := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	ctx := t.Context()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	assert.NoError(t, err)

	val, err := rdb.Get(ctx, "key").Result()
	assert.NoError(t, err)
	assert.Equal(t, "value", val)

	_, err = rdb.Get(ctx, "key2").Result()
	assert.Equal(t, redis.Nil, err)
}
