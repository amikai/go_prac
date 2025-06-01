package syncex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/singleflight"
)

func TestSingleflightDo(t *testing.T) {
	sg := new(singleflight.Group)
	x := 0
	for i := 0; i < 1024; i++ {
		go func() {
			got, err, _ := sg.Do("key", func() (interface{}, error) {
				x++
				// long time process
				time.Sleep(1 * time.Second)
				return x, nil
			})
			assert.NoError(t, err)
			assert.Equal(t, got.(int), 1)
		}()
	}
}
