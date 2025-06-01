package syncex

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for range 1024 {
		wg.Add(1)
		go func() {
			instance := GetInstance()
			assert.Equal(t, 1, instance.Count)
			wg.Done()
		}()
	}
	wg.Wait()
}
