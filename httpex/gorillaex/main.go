package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"

	"github.com/amikai/go_prac/httpex"
)

type Resp[T any] struct {
	Data  T      `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
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
		logger.Info("gracefully quit gorilla server")
	} else if err != nil {
		logger.Error(err.Error())
	}
}
