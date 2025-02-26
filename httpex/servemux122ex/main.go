package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Resp[T any] struct {
	Data  T      `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

const (
	ProductAPI       string = "GET /products/{id}"
	BooksCategoryAPI string = "GET /books/{category}/"
	BooksAPI         string = "GET /books/{id}"
)

func NotFoundHandler(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug(
			r.URL.Path,
			slog.Int("code", 404),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("query", r.URL.RawQuery),
			slog.String("ip", r.RemoteAddr),
			slog.String("user-agent", r.UserAgent()),
		)
		http.NotFound(w, r)
	})
}

func main() {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(h)

	r := http.NewServeMux()
	r.HandleFunc(ProductAPI, ProductHandler)
	r.HandleFunc(BooksCategoryAPI, BooksCategoryHandler)
	r.HandleFunc(BooksAPI, BooksHandler)

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
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("gracefully quit gorilla server")
	}
}
