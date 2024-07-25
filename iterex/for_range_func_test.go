package iterex

import (
	"iter"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func EvenNumsInt8(yield func(int) bool) {
	for i := 0; i <= math.MaxInt8-1; i += 2 {
		if !yield(i) {
			break
		}
	}
}

func BiggerThan100(in iter.Seq[int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		for v := range in {
			if v >= 100 {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func TestEvenNums(t *testing.T) {
	x := 0
	EvenNumsInt8(func(n int) bool {
		assert.Equal(t, x, n)
		x = x + 2
		return true
	})
}

func TestEvenNumsForRange(t *testing.T) {
	x := 0
	for n := range EvenNumsInt8 {
		assert.Equal(t, x, n)
		x = x + 2
	}
}

func TestEvenNumsBiggerThan100(t *testing.T) {
	x := 100
	for n := range BiggerThan100(EvenNumsInt8) {
		assert.Equal(t, x, n)
		x = x + 2
	}
}
