package loop

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func cStyleLoopSum[T constraints.Unsigned](n T) T {
	var sum T
	for i := T(1); i <= n; i++ {
		sum += i
	}
	return sum
}

func whileLoopSum[T constraints.Unsigned](n T) T {
	var i, sum T
	for i <= n {
		sum += i
		i++
	}
	return sum
}

func infiniteLoopSum[T constraints.Unsigned](n T) T {
	var i, sum T
	for {
		if i > n {
			break
		}
		sum += i
		i++
	}
	return sum
}

func forEachArraySum[T constraints.Unsigned](arr []T) T {
	var sum T
	for _, element := range arr {
		sum += element
	}
	return sum
}
