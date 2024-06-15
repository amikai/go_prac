// The contents of this file are entirely experimental, unverified, and not robust.
package iterex

import (
	"fmt"
	"iter"
)

func gen() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	return ch
}

// TODO: Handle the case when for-range All break.
func All[Ch chan E, E any](ch Ch) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range ch {
			if !yield(v) {
				return
			}
		}
	}
}

func ExampleChAll() {
	ch := gen()
	for v := range All(ch) {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}
