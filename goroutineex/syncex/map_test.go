package syncex

import (
	"fmt"
	"sync"
	"testing"
)

func ExampleMapRange(t *testing.T) {
	var m sync.Map

	m.Store("alice", 11)
	m.Store("bob", 12)
	m.Store("cindy", 13)

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	// Output:
	// alice 11
	// bob 12
	// cindy 13
}

func ExampleMapRangeForRange(t *testing.T) {
	var m sync.Map

	m.Store("alice", 11)
	m.Store("bob", 12)
	m.Store("cindy", 13)

	for key, value := range m.Range {
		fmt.Println(key, value)
	}

	// Output:
	// alice 11
	// bob 12
	// cindy 13
}
