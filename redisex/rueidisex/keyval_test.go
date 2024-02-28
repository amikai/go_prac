package rueidisex

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/rueidis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSet(t *testing.T) {
	s := miniredis.RunT(t)
	// Disable client cache because mini-redis doest not support now
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{s.Addr()}, DisableCache: true})
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()
	err = client.Do(ctx, client.B().Set().Key("key").Value("val").Build()).Error()
	require.NoError(t, err)
	val, err := client.Do(ctx, client.B().Get().Key("key").Build()).ToString()
	require.NoError(t, err)
	assert.Equal(t, "val", val)
}
