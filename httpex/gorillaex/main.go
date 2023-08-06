package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"

	"github.com/amikai/go_prac/httpex"
	"github.com/amikai/go_prac/httpex/db"
)

type Resp[T any] struct {
	Data  T      `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

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

func NotFoundHandler(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug(r.URL.Path, slog.Int("code", 404), slog.String("method", r.Method), slog.String("path", r.URL.Path), slog.String("query", r.URL.RawQuery),
			slog.String("ip", r.RemoteAddr), slog.String("user-agent", r.UserAgent()))
		http.NotFound(w, r)
	})
}

func main() {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(h)

	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", ProductHandler).Methods("GET")
	r.HandleFunc("/books/{category}/", BooksCategoryHandler).Methods("GET")
	r.HandleFunc("/books/{id:BOOK-[0-9]+}", BooksHandler).Methods("GET")
	r.Use(mux.MiddlewareFunc(httpex.RequestLogMiddleware(logger)))
	r.NotFoundHandler = NotFoundHandler(logger)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	g, errCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	g.Go(func() error {
		<-errCtx.Done()
		return srv.Shutdown(context.Background())
	})

	err := g.Wait()
	if errors.Is(err, context.Canceled) || err == nil {
		log.Println("gracefully quit gorilla server")
	} else if err != nil {
		log.Println(err)
	}
}
