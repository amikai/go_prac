package assertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntZero(t *testing.T) {
	var i int
	assert.Zero(t, i)
	// zero value of int is 0
	assert.Equal(t, 0, i)
}

func TestFloatZero(t *testing.T) {
	var f float64
	assert.Zero(t, f)
	// zero value of float64 is 0.0
	assert.Equal(t, 0, 0, f)
}

func TestStringZero(t *testing.T) {
	var s string
	assert.Zero(t, s)
	// zero value of string is ""
	assert.Equal(t, "", s)
}

func PointerZero(t *testing.T) {
	var p *int
	assert.Zero(t, p)
	// zero value of pointer is nil
	assert.Nil(t, p)
}

func TestMapZero(t *testing.T) {
	var m map[string]interface{}
	assert.Zero(t, m)
	// zero value of map is nil
	assert.Nil(t, m)
}

func TestSliceZero(t *testing.T) {
	var sli []interface{}
	assert.Zero(t, sli)
	// zero value of slice is nil
	assert.Nil(t, sli)
}

func TestZeroStruct(t *testing.T) {
	type all struct {
		i   int
		f   float64
		s   string
		m   map[string]interface{}
		sli []interface{}
	}

	a := all{}
	assert.Zero(t, a)
	assert.Equal(t, all{
		i:   0,
		f:   0.0,
		s:   "",
		m:   nil,
		sli: nil,
	}, a)
}
