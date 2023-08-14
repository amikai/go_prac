package syncex

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnce(t *testing.T) {
	a := 0
	var once sync.Once
	wg := sync.WaitGroup{}
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() { a++ })
			wg.Done()
		}()
	}
	assert.Equal(t, 1, a)
}

func TestOnceFunc(t *testing.T) {
	a := 0
	onceF := sync.OnceFunc(func() {
		a++
	})

	wg := sync.WaitGroup{}
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		go func() {
			onceF()
			wg.Done()
		}()
	}
	assert.Equal(t, 1, a)
}

func TestOnceValue(t *testing.T) {
	a := 0
	onceF := sync.OnceValue(func() int {
		a++
		return a
	})

	wg := sync.WaitGroup{}
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		go func() {
			assert.Equal(t, 1, onceF())
			wg.Done()
		}()
	}
	assert.Equal(t, 1, a)
}
