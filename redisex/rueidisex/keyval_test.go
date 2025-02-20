package rueidisex

import (
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

	err = client.Do(t.Context(), client.B().Set().Key("key").Value("val").Build()).Error()
	require.NoError(t, err)
	val, err := client.Do(t.Context(), client.B().Get().Key("key").Build()).ToString()
	require.NoError(t, err)
	assert.Equal(t, "val", val)
}
