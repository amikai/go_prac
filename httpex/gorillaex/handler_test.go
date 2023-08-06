package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/amikai/go_prac/httpex/db"
)

var fakeProduct = db.Product{
	ID:   "ID-FAKE",
	Name: "FAKE_NAME",
}

func spyGetProductByID(ID string) *db.Product {
	if ID == fakeProduct.ID {
		return &fakeProduct
	}
	return nil
}

// TODO: teet the other product ID that doesn't get the result
func TestProductHandler(t *testing.T) {
	req := httptest.NewRequest("GET", fmt.Sprintf("/products/%s", fakeProduct.ID), nil)
	rr := httptest.NewRecorder()
	db.GetProductByID = spyGetProductByID

	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", ProductHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	resp := rr.Result()
	gotBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	wantBody := `
		{
		  "data": {
			"id": "ID-FAKE",
			"name": "FAKE_NAME"
		  }
		}`
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.JSONEq(t, string(wantBody), string(gotBody))
}

func TestBooksCategoryHandler(t *testing.T) {
	t.Skip("TODO: implement this test")
}

func TestBooksHandler(t *testing.T) {
	t.Skip("TODO: implement this test")
}
