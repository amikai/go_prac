package goroutineex

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var MaxSliceLen = 10000
var MaxChannelNum = 100

func buildSlice[T any](gen func(int) T) []T {
	randLen := rand.Intn(MaxSliceLen)
	ret := make([]T, randLen)
	for i := 0; i < randLen; i++ {
		ret[i] = gen(i)
	}
	return ret
}

func buildChannels[T any]() []chan T {
	ret := make([]chan T, rand.Intn(MaxChannelNum))
	for i := 0; i < len(ret); i++ {
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
		whichCh := rand.Intn(len(chs))
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
