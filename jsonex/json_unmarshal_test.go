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
func TestUmarshalIntSameZeroVal(t *testing.T) {
	want := FieldStruct[int]{}
	t.Run("unmarshal with field exist", func(t *testing.T) {
		inputWithField := []byte(`{"field": 0}`)
		var got FieldStruct[int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})

	t.Run("unmarshal with not exist", func(t *testing.T) {
		inputWithField := []byte(`{}`)
		var got FieldStruct[int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})

}

func TestUmarshalArrIntSameZeroVal(t *testing.T) {
	want := FieldStruct[int]{}
	t.Run("unmarshal with field exist", func(t *testing.T) {
		inputWithField := []byte(`{"field": 0}`)
		var got FieldStruct[int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})

	t.Run("unmarshal with not exist", func(t *testing.T) {
		inputWithField := []byte(`{}`)
		var got FieldStruct[int] // zero value
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})

}

type IntPointerField struct {
	PointerField *int `json:"field"`
}

// This example show that we tell that json field is exist after unmarshal when field is pointer type
func TestUmarshalPointerVal(t *testing.T) {
	t.Run("unmarshal with field exist", func(t *testing.T) {
		inputWithField := []byte(`{"field": 0}`)
		var got IntPointerField // zero value
		want := IntPointerField{PointerField: new(int)}
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})

	t.Run("unmarshal with not exist", func(t *testing.T) {
		inputWithField := []byte(`{}`)
		var got IntPointerField // zero value
		want := IntPointerField{PointerField: nil}
		err := json.Unmarshal(inputWithField, &got)
		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})
}
