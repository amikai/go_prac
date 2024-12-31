package jsonex

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Company struct {
	Employees []*Employee `json:"employees"`
	Name      string      `json:"name"`
}

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestUnmarshal(t *testing.T) {
	data := []byte(`
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
}`)

	want := Company{
		Employees: []*Employee{
			{ID: "001", Name: "John"},
			{ID: "002", Name: "Mary"},
			{ID: "003", Name: "Bob"},
		},
		Name: "ABCCompany",
	}

	var got Company
	err := json.Unmarshal(data, &got)
	assert.NoError(t, err)
	assert.Equal(t, got, want)
}

func TestUnmarshalFieldNotMatch(t *testing.T) {
	// By default, object keys which don't have a corresponding struct field are ignored
	data := []byte(`
{
	"foo": "bar"
}`)
	var want Company
	var got Company
	err := json.Unmarshal(data, &got)
	assert.NoError(t, err)
	assert.Equal(t, got, want)
}

func TestMarshal(t *testing.T) {
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

	got, err := json.Marshal(&company)
	assert.NoError(t, err)
	assert.JSONEq(t, want, string(got))

}

type B64Data struct {
	Data []byte `json:"data"`
}

func TestUnmarshalB64Data(t *testing.T) {
	// echo -n 'Hello World' | base64
	data := []byte(`
{
	"data": "SGVsbG8gV29ybGQ="
}`)

	want := B64Data{
		Data: []byte("Hello World"),
	}

	var got B64Data
	err := json.Unmarshal(data, &got)
	require.NoError(t, err)
	assert.Equal(t, got, want)

}
