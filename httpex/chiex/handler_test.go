package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amikai/go_prac/httpex/db"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
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

func TestProductHandler(t *testing.T) {
	req := httptest.NewRequest("GET", fmt.Sprintf("/products/%s", fakeProduct.ID), nil)
	rr := httptest.NewRecorder()
	db.GetProductByID = spyGetProductByID

	router := chi.NewRouter()
	router.Get("/products/{id}", ProductHandler)
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

func TestProductNotFound(t *testing.T) {
	dummyID := "dummy"
	req := httptest.NewRequest("GET", fmt.Sprintf("/products/%s", dummyID), nil)
	rr := httptest.NewRecorder()
	db.GetProductByID = spyGetProductByID

	router := chi.NewRouter()
	router.Get("/products/{id}", ProductHandler)
	router.ServeHTTP(rr, req)

	resp := rr.Result()
	gotBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	wantBody := `{}`
	assert.JSONEq(t, string(wantBody), string(gotBody))
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}

var fakeBooks = map[string]*db.Book{
	"BOOK-999": {
		ID:       "BOOK-999",
		Name:     "Fake name",
		Category: "Fake category",
	},
}

var booksByCategory = map[string][]*db.Book{
	"Fake category": {
		fakeBooks["BOOK-999"],
	},
}

func spyBooksByCategory(category string) []*db.Book {
	return booksByCategory[category]
}

func TestBooksCategoryHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/books/Fake%20category/", nil)
	rr := httptest.NewRecorder()
	db.GetBooksByCategory = spyBooksByCategory

	router := chi.NewRouter()
	router.Get("/books/{category}/", BooksCategoryHandler)
	router.ServeHTTP(rr, req)

	resp := rr.Result()
	gotBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	wantBody := `
{
  "data": [
    {
      "id": "BOOK-999",
      "name": "Fake name",
      "category": "Fake category"
    }
  ]
}`
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.JSONEq(t, string(wantBody), string(gotBody))
}

func TestBooksHandler(t *testing.T) {
	t.Skip("TODO: implement this test")
}

func testHTTPGetAPIBodyJSONEquality(t *testing.T, c *http.Client, url, want string) {
	resp, err := c.Get(url)
	assert.NoError(t, err)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, want, string(b))
}

func TestRouterE2E(t *testing.T) {
	db.GetProductByID = spyGetProductByID
	db.GetBooksByCategory = spyBooksByCategory
	s := httptest.NewServer(newRouter(newLogger()))
	defer s.Close()

	// TODO: convert it to test table

	url := s.URL + "/products/" + fakeProduct.ID
	wantBody := `
{
  "data": {
    "id": "ID-FAKE",
    "name": "FAKE_NAME"
  }
}`
	testHTTPGetAPIBodyJSONEquality(t, s.Client(), url, wantBody)

	url = s.URL + "/products/dummy"
	wantBody = `{}`
	testHTTPGetAPIBodyJSONEquality(t, s.Client(), url, wantBody)

	url = s.URL + "/books/Fake%20category/"
	wantBody = `
{
  "data": [
    {
      "id": "BOOK-999",
      "name": "Fake name",
      "category": "Fake category"
    }
  ]
}`
	testHTTPGetAPIBodyJSONEquality(t, s.Client(), url, wantBody)

}
