package rillex

import (
	"fmt"
	"testing"

	"github.com/destel/rill"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	rch := rill.FromChan(gen(5), nil)
	rch = rill.Map(rch, concLimit, func(x int) (int, error) {
		return x * 3, nil
	})
	ret, err := rill.ToSlice(rch)

	require.NoError(t, err)
	assert.ElementsMatch(t, []int{3, 6, 9, 12, 15}, ret)
}

func TestFlatMap(t *testing.T) {
	rch := rill.FromChan(gen(5), nil)

	rch2 := rill.FlatMap(rch, concLimit, func(x int) <-chan rill.Try[string] {
		ch := make(chan rill.Try[string])
		go func() {
			ch <- rill.Wrap(fmt.Sprintf("a%d", x), nil)
			ch <- rill.Wrap(fmt.Sprintf("b%d", x), nil)
			ch <- rill.Wrap(fmt.Sprintf("c%d", x), nil)
			close(ch)
		}()
		return ch
	})
	ret, err := rill.ToSlice(rch2)
	require.NoError(t, err)
	assert.ElementsMatch(
		t,
		[]string{
			"a1",
			"b1",
			"c1",
			"a2",
			"b2",
			"c2",
			"a3",
			"b3",
			"c3",
			"a4",
			"b4",
			"c4",
			"a5",
			"b5",
			"c5",
		},
		ret,
	)
}

func TestOrderMap(t *testing.T) {
	rch := rill.FromChan(gen(5), nil)
	rch = rill.OrderedMap(rch, concLimit, func(x int) (int, error) {
		return x * 3, nil
	})
	ret, err := rill.ToSlice(rch)

	require.NoError(t, err)
	assert.Equal(t, []int{3, 6, 9, 12, 15}, ret)
}
