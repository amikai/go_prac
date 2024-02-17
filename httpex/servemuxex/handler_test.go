package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amikai/go_prac/httpex/db"
	"github.com/stretchr/testify/assert"
)

var fakeProduct = db.Product{
	ID:   "ID-FAKE",
	Name: "FAKE_NAME",
}

func spyGetProductByID(id string) *db.Product {
	if id == fakeProduct.ID {
		return &fakeProduct
	}
	return nil
}

func TestProductHandler(t *testing.T) {
	req := httptest.NewRequest(ProductAPI.Method, "/products/ID-FAKE", nil)
	rr := httptest.NewRecorder()

	db.GetProductByID = spyGetProductByID
	mux := http.NewServeMux()
	mux.HandleFunc(ProductAPI.ServeMuxPattern(), ProductHandler)
	mux.ServeHTTP(rr, req)

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
	req := httptest.NewRequest(ProductAPI.Method, "/products/dummy", nil)
	rr := httptest.NewRecorder()
	db.GetProductByID = spyGetProductByID

	mux := http.NewServeMux()
	mux.HandleFunc(ProductAPI.ServeMuxPattern(), ProductHandler)
	mux.ServeHTTP(rr, req)

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
	req := httptest.NewRequest(BooksCategoryAPI.Method, "/books/Fake%20category/", nil)
	rr := httptest.NewRecorder()
	db.GetBooksByCategory = spyBooksByCategory

	mux := http.NewServeMux()
	mux.HandleFunc(BooksCategoryAPI.ServeMuxPattern(), BooksCategoryHandler)
	mux.ServeHTTP(rr, req)

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
	s := httptest.NewServer(newRouter())
	defer s.Close()

	tests := map[string]struct {
		reqURL   string
		wantBody string
	}{
		"product found": {
			reqURL: s.URL + "/products/" + fakeProduct.ID,
			wantBody: `
{
  "data": {
    "id": "ID-FAKE",
    "name": "FAKE_NAME"
  }
}`,
		},
		"product not found": {
			reqURL:   s.URL + "/products/dummy",
			wantBody: `{}`,
		},
		"books category found": {
			reqURL: s.URL + "/books/Fake%20category/",
			wantBody: `
{
  "data": [
    {
      "id": "BOOK-999",
      "name": "Fake name",
      "category": "Fake category"
    }
  ]
}`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			testHTTPGetAPIBodyJSONEquality(t, s.Client(), tc.reqURL, tc.wantBody)

		})
	}
}
