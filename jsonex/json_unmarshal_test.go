package jsonex

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FieldStruct[T any] struct {
	Field T `json:"field"`
}

// This example show that we cannot tell that json field is exist after unmarshal when field is non pointer type
func TestUmarshalIntVal(t *testing.T) {
	want := FieldStruct[int]{}
	t.Run("unmarshal not exist field", func(t *testing.T) {
		inputWithField := []byte(`{}`)
		var got FieldStruct[int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})

	t.Run("unmarshal null field", func(t *testing.T) {
		inputWithField := []byte(`{"field": null}`)
		var got FieldStruct[int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})

	t.Run("unmarshal zero value field", func(t *testing.T) {
		inputWithField := []byte(`{"field": 0}`)
		var got FieldStruct[int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})
}

func TestUmarshalIntPointerVal(t *testing.T) {
	t.Run("unmarshal not exist field", func(t *testing.T) {
		inputWithField := []byte(`{}`)
		var got FieldStruct[*int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		want := FieldStruct[*int]{
			nil,
		}
		assert.Equal(t, got, want)
	})

	t.Run("unmarshal zero value field", func(t *testing.T) {
		inputWithField := []byte(`{"field": 0}`)
		var got FieldStruct[*int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		want := FieldStruct[*int]{
			new(int),
		}
		assert.Equal(t, got, want)
	})

	t.Run("unmarshal null field", func(t *testing.T) {
		inputWithField := []byte(`{"field": null}`)
		var got FieldStruct[*int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		want := FieldStruct[*int]{
			nil,
		}
		assert.Equal(t, got, want)
	})
}
func TestUmarshalIntArr(t *testing.T) {
	t.Run("unmarshal not exist field", func(t *testing.T) {
		inputWithField := []byte(`{}`)
		var got FieldStruct[[]int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		want := FieldStruct[[]int]{}
		assert.Equal(t, got, want)
	})

	t.Run("unmarshal null field to int array", func(t *testing.T) {
		inputWithField := []byte(`{"field": null}`)
		var got FieldStruct[[]int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		want := FieldStruct[[]int]{}
		assert.Equal(t, got, want)
	})

	t.Run("unmarshal zero length array", func(t *testing.T) {
		inputWithField := []byte(`{"field":[]}`)
		var got FieldStruct[[]int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		want := FieldStruct[[]int]{Field: []int{}}
		assert.Equal(t, got, want)
	})
}
