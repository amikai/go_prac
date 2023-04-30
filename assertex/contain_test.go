package assertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapContain(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	assert.Contains(t, m, "a")
	assert.Contains(t, m, "c")
	assert.Contains(t, m, "c")
	assert.NotContains(t, m, "z")
}

func TestSliceContain(t *testing.T) {
	s := []string{"a", "b", "xyz"}
	assert.Contains(t, s, "a")
	assert.Contains(t, s, "xyz")
	assert.NotContains(t, s, "z")
}

func TestStringContain(t *testing.T) {
	// "cde" is a substring of s
	assert.Contains(t, "abcdefghi", "cde")
	// "xyz" is not a substring of s
	assert.NotContains(t, "abcdefghi", "xyz")
}

func TestSliceContainStruct(t *testing.T) {
	p1 := Person{
		Name: "Amikai",
		Age:  29,
	}
	p2 := Person{
		Name: "John",
		Age:  30,
	}
	p3 := Person{
		Name: "Mary",
		Age:  31,
	}

	persons := []Person{p1, p2}
	assert.Contains(t, persons, p1)
	assert.Contains(t, persons, p2)
	assert.NotContains(t, persons, p3)
}

func TestSliceContainPointerStruct(t *testing.T) {
	p1 := &Person{
		Name: "Amikai",
		Age:  29,
	}
	p2 := &Person{
		Name: "John",
		Age:  30,
	}
	p3 := &Person{
		Name: "Mary",
		Age:  31,
	}

	persons := []*Person{p1, p2}
	assert.Contains(t, persons, p1)
	assert.Contains(t, persons, p2)
	assert.NotContains(t, persons, p3)
}
