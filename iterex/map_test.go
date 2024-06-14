package iterex

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMapsAll() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i, v := range maps.All(m) {
		fmt.Printf("%s, %d\n", i, v)
	}
}

func TestMapsCollectAll(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	seq := maps.All(m)
	m2 := maps.Collect(seq)
	assert.Equal(t, m, m2)
}

func TestMapsKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	seq := maps.Keys(m)
	assert.ElementsMatch(t, []string{"a", "b", "c"}, slices.Collect(seq))
}

func TestMapsValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	seq := maps.Values(m)
	assert.ElementsMatch(t, []int{1, 2, 3}, slices.Collect(seq))
}
