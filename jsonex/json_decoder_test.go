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
