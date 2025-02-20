package goredisex

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func newPool(address string) *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			c, err := redis.DialContext(ctx, "tcp", address)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func TestGetSet(t *testing.T) {
	s := miniredis.RunT(t)
	pool := newPool(s.Addr())
	conn := pool.Get()
	defer conn.Close()

	resp, err := redis.DoContext(conn, t.Context(), "PING")
	resp, err = redis.String(resp, err)
	assert.NoError(t, err)
	assert.Equal(t, resp, "PONG")

	resp, err = redis.DoContext(conn, t.Context(), "SET", "key", "value")
	resp, err = redis.String(resp, err)
	assert.NoError(t, err)
	assert.Equal(t, resp, "OK")

	resp, err = redis.DoContext(conn, t.Context(), "GET", "key")
	resp, err = redis.String(resp, err)
	assert.NoError(t, err)
	assert.Equal(t, resp, "value")
}
