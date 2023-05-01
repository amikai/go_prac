package builtin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestMapClear(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	maps.Clear(m)
	assert.Empty(t, m)
}

func TestMapClone(t *testing.T) {
	src := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	dst := maps.Clone(src)
	assert.Equal(t, src, dst)
}

func TestMapCopy(t *testing.T) {
	src := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	dst := make(map[string]int, len(src))
	maps.Copy(dst, src)
	assert.Equal(t, src, dst)
}

func TestMapKeys(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	keys := maps.Keys(m)
	assert.ElementsMatch(t, keys, []string{"c", "b", "a"})
}

func TestMapValues(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 2,
	}
	keys := maps.Values(m)
	assert.ElementsMatch(t, keys, []int{2, 1, 2})
}

func TestMapEqual(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	assert.True(t, maps.Equal(m1, m2))
}

type People struct {
	Name string
	Age  int
}

func (p1 People) Equal(p2 People) bool {
	return p1.Name == p2.Name && p1.Age == p2.Age
}

func TestMapEqualFunc(t *testing.T) {
	m1 := map[string]People{
		"Mary": {Name: "Mary", Age: 20},
		"John": {Name: "John", Age: 30},
		"Bob":  {Name: "Bob", Age: 40},
	}
	m2 := map[string]People{
		"Mary": {Name: "Mary", Age: 20},
		"John": {Name: "John", Age: 30},
		"Bob":  {Name: "Bob", Age: 40},
	}
	assert.True(t, maps.EqualFunc(m1, m2, func(p1, p2 People) bool {
		return p1.Equal(p2)
	}))
}

func TestDeleteFunc(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
	}
	maps.DeleteFunc(m, func(k string, v int) bool {
		return v%2 == 0
	})

	exp := map[string]int{
		"a": 1,
		"c": 3,
		"e": 5,
		"g": 7,
	}
	assert.Equal(t, exp, m)
}
