package rillex

import (
	"testing"
	"time"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func genWithTick(n int, ticker *time.Ticker) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 1; i <= n; i++ {
			<-ticker.C
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func TestBatch(t *testing.T) {
	rch := rill.FromChan(genWithTick(1000, time.NewTicker(1*time.Millisecond)), nil)
	brch := rill.Batch(rch, concLimit, 20*time.Millisecond)
	mrch := rill.Map(brch, concLimit, func(arr []int) (int, error) {
		sum := 0
		for _, v := range arr {
			sum += v
		}
		return sum, nil
	})

	ret, err := rill.ToSlice(mrch)
	require.NoError(t, err)

	sum := 0
	for _, v := range ret {
		sum += v
	}

	assert.Equal(t, 500500, sum)
}
