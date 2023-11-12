package main

import (
	"encoding/json"
	"net/http"

	"github.com/amikai/go_prac/httpex/db"
	"github.com/go-chi/chi/v5"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp := Resp[*db.Product]{
		Data: db.GetProductByID(id),
	}

	b, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func BooksCategoryHandler(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	resp := Resp[[]*db.Book]{
		Data: db.GetBooksByCategory(category),
	}

	b, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resp := Resp[*db.Book]{
		Data: db.GetBooksByID(id),
	}

	b, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
