package testingex

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeElapse(t *testing.T) {
	synctest.Run(func() {
		start := time.Now()
		time.Sleep(time.Nanosecond)
		assert.Equal(t, time.Nanosecond, time.Since(start))
	})
}
