package main

import (
	"encoding/json"
	"net/http"

	"github.com/amikai/go_prac/httpex/db"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

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
	category := r.PathValue("category")
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
	id := r.PathValue("id")
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
