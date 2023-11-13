package jsonex

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The "omitempty" option specifies that the field should be omitted from the encoding if the field has an empty value,
// defined as false, 0, a nil pointer, a nil interface value, and any empty array, slice, map, or string.
type BasicTypeOmitEmpty[T any] struct {
	Int       int          `json:"int,omitempty"`
	Float     float64      `json:"float,omitempty"`
	String    string       `json:"string,omitempty"`
	Bool      bool         `json:"bool,omitempty"`
	Array     []T          `json:"array,omitempty"`
	Map       map[string]T `json:"map,omitempty"`
	Pointer   *T           `json:"pointer,omitempty"`
	Interface interface{}  `json:"interface,omitempty"`
}

type BasicType[T any] struct {
	Int       int          `json:"int"`
	Float     float64      `json:"float"`
	String    string       `json:"string"`
	Bool      bool         `json:"bool"`
	Array     []T          `json:"array"`
	Map       map[string]T `json:"map"`
	Pointer   *T           `json:"pointer"`
	Interface interface{}  `json:"interface"`
}

func TestMarshalZeroValue(t *testing.T) {
	t.Run("basic type with omitempty", func(t *testing.T) {
		bo := BasicTypeOmitEmpty[int]{}
		got, err := json.Marshal(&bo)
		assert.NoError(t, err)
		want := `{}`
		assert.JSONEq(t, want, string(got))
	})

	t.Run("basic type", func(t *testing.T) {
		bo := BasicType[int]{}
		got, err := json.Marshal(&bo)
		assert.NoError(t, err)
		want := `{
"int": 0,
"float": 0.0,
"string": "",
"bool": false,
"array":null,
"map": null,
"pointer": null,
"interface": null
		}`
		assert.JSONEq(t, want, string(got))
	})
}
