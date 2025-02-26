package assertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
}

// The implementation of assert.Equal utilizes reflect.DeepEqual.

func TestCompareStruct(t *testing.T) {
	p := Person{}
	e := Employee{}
	assert.NotEqual(t, p, e)
	assert.EqualValues(t, p, e)
	// can not compile this line
	// assert.True(t, p == e)
}

func TestCompareAnonymousStruct(t *testing.T) {
	p := Person{}
	a := struct {
		Name string
		Age  int
	}{}

	// Type of p and a are different
	assert.NotEqual(t, p, a)

	// Values of p and a are same
	assert.EqualValues(t, p, a)

	// Both structs have the same names, orders and types
	assert.True(t, p == a)
}

func TestCompareAnonymousStruct2(t *testing.T) {
	p := Person{}
	a := struct {
		Age  int
		Name string
	}{}

	assert.NotEqual(t, p, a)

	assert.NotEqualValues(t, p, a)
}

func TestMapEq(t *testing.T) {
	m1 := map[string]int{
		"a": 1, "b": 2, "c": 3,
	}
	m2 := map[string]int{
		"c": 3, "b": 2, "a": 1,
	}
	assert.Equal(t, m1, m2)
}

func TestSliceEq(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	assert.Equal(t, s1, s2)
}

// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func TestPointToSameValueEq(t *testing.T) {
	a := 1
	b := 1
	pa := &a
	pb := &b
	assert.Equal(t, pa, pb)
}

func TestSameTypeAlias(t *testing.T) {
	type (
		A = int
		B = int
	)

	var a A = 1
	var b B = 1
	assert.Equal(t, a, b)
}

func TestSameUnderlyingTypeEqual(t *testing.T) {
	type (
		A int
		B int
	)

	var a A = 1
	var b B = 1
	assert.NotEqual(t, a, b)
}

func TestSameUnderlyingTypeEqualValues(t *testing.T) {
	type (
		A int
		B int
	)

	var a A = 1
	var b B = 1
	assert.EqualValues(t, a, b)
}
