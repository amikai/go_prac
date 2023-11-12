package jsonex

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonDecoder(t *testing.T) {
	data := `
{
	"employees": [{
			"id": "001",
			"name": "John"
		},
		{
			"id": "002",
			"name": "Mary"
		},
		{
			"id": "003",
			"name": "Bob"
		}
	],
	"name": "ABCCompany"
}`

	want := Company{
		Employees: []*Employee{
			{ID: "001", Name: "John"},
			{ID: "002", Name: "Mary"},
			{ID: "003", Name: "Bob"},
		},
		Name: "ABCCompany",
	}

	var got Company
	dec := json.NewDecoder(strings.NewReader(data))
	err := dec.Decode(&got)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestJsonDecodeDisallowUnknownFields(t *testing.T) {
	data := `
{
	"employees": [{
			"id": "001",
			"name": "John"
		}
	],
	"name": "ABCCompany",
	"foo": "bar"
}`
	var got Company
	dec := json.NewDecoder(strings.NewReader(data))
	dec.DisallowUnknownFields()
	err := dec.Decode(&got)
	assert.Error(t, err)
}

func TestJsonEncoder(t *testing.T) {
	company := Company{
		Employees: []*Employee{
			{ID: "001", Name: "John"},
			{ID: "002", Name: "Mary"},
			{ID: "003", Name: "Bob"},
		},
		Name: "ABCCompany",
	}

	want := `{
	"name": "ABCCompany",
	"employees": [{
			"id": "001",
			"name": "John"
		},
		{
			"id": "002",
			"name": "Mary"
		},
		{
			"id": "003",
			"name": "Bob"
		}
	]
}`
	sb := &strings.Builder{}
	enc := json.NewEncoder(sb)
	err := enc.Encode(&company)
	assert.NoError(t, err)
	assert.JSONEq(t, want, sb.String())

}

func TestParseAsJsonNumber(t *testing.T) {
	data := `{"foo": 123}`
	m := map[string]interface{}{}

	dec := json.NewDecoder(strings.NewReader(data))
	dec.UseNumber()
	err := dec.Decode(&m)

	assert.NoError(t, err)
	num := m["foo"].(json.Number)

	assert.Equal(t, num.String(), "123")

	nf, err := num.Float64()
	assert.NoError(t, err)
	assert.Equal(t, nf, float64(123))

	ni, err := num.Int64()
	assert.NoError(t, err)
	assert.Equal(t, ni, int64(123))
}
