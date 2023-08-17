package syncex

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		go func() {
			instance := GetInstance()
			assert.Equal(t, instance.Count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
}
