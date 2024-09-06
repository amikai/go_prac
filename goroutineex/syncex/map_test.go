package syncex

import (
	"fmt"
	"sync"
)

func Example_syncMapRange() {
	var m sync.Map

	m.Store("alice", 11)
	m.Store("bob", 12)
	m.Store("cindy", 13)

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	// Unordered Output:
	// alice 11
	// bob 12
	// cindy 13
}

func Example_syncMapRangeFor() {
	var m sync.Map

	m.Store("alice", 11)
	m.Store("bob", 12)
	m.Store("cindy", 13)

	for key, value := range m.Range {
		fmt.Println(key, value)
	}

	// Unordered Output:
	// alice 11
	// bob 12
	// cindy 13
}
