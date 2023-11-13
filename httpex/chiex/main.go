package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/amikai/go_prac/httpex"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
)

type Resp[T any] struct {
	Data  T      `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

const (
	ProductAPI       string = "/products/{id}"
	BooksCategoryAPI string = "/books/{category}/"
	BooksAPI         string = "/books/{id}"
)

func newRouter(logger *slog.Logger) *chi.Mux {
	r := chi.NewRouter()
	// chi choose func(http.Handler) http.Handler as middleware type,
	// this is the common type for http middlware in community
	r.Use(httpex.RequestLogMiddleware(logger))
	r.Get(ProductAPI, ProductHandler)
	r.Get(BooksCategoryAPI, BooksCategoryHandler)
	r.Get(BooksAPI, BooksHandler)
	return r
}

func newLogger() *slog.Logger {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	return slog.New(h)

}

func main() {
	logger := newLogger()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: newRouter(logger),
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
		logger.Info("gracefully quit chi server")
	} else if err != nil {
		logger.Error(err.Error())
	}
}
