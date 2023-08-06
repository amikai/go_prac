package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/amikai/go_prac/httpex/db"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	ID := vars["id"]
	resp := Resp[*db.Product]{
		Data: db.GetProductByID(ID),
	}

	enc := json.NewEncoder(w)
	err := enc.Encode(resp)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func BooksCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	category := vars["category"]
	resp := Resp[[]*db.Book]{
		Data: db.GetBooksByCategory(category),
	}
	enc := json.NewEncoder(w)
	err := enc.Encode(resp)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	ID := vars["id"]
	resp := Resp[*db.Book]{
		Data: db.GetBooksByID(ID),
	}

	enc := json.NewEncoder(w)
	err := enc.Encode(resp)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}
