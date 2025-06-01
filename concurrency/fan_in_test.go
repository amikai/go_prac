package goroutineex

import (
	"math/rand/v2"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	MaxSliceLen   = 10000
	MaxChannelNum = 100
)

func buildSlice[T any](gen func(int) T) []T {
	randLen := rand.N(MaxSliceLen)
	ret := make([]T, randLen)
	for i := range randLen {
		ret[i] = gen(i)
	}
	return ret
}

func buildChannels[T any]() []chan T {
	ret := make([]chan T, rand.N(MaxChannelNum))
	for i := range ret {
		ret[i] = make(chan T)
	}
	return ret
}

func TestMerge(t *testing.T) {
	intSlice := buildSlice(func(i int) int {
		return rand.Int()
	})
	chs := buildChannels[int]()
	wg := sync.WaitGroup{}

	send := func(ch chan int, n int) {
		ch <- n
		wg.Done()
	}

	wg.Add(len(intSlice))
	for _, num := range intSlice {
		whichCh := rand.N(len(chs))
		go send(chs[whichCh], num)
	}

	// wait for finishing sending, then close all channels
	go func() {
		wg.Wait()
		for _, ch := range chs {
			close(ch)
		}
	}()

	out := Merge(chs...)

	var got []int
	for num := range out {
		got = append(got, num)
	}
	assert.ElementsMatch(t, intSlice, got)
}
