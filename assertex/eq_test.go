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
