package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/amikai/go_prac/httpex"
	"golang.org/x/sync/errgroup"
)

type Resp[T any] struct {
	Data  T      `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

type HTTPAPI struct {
	Path   string
	Method string
}

func (h HTTPAPI) ServeMuxPattern() string {
	return h.Method + " " + h.Path
}

var (
	ProductAPI       = HTTPAPI{Path: "/products/{id}", Method: "GET"}
	BooksCategoryAPI = HTTPAPI{Path: "/books/{category}/", Method: "GET"}
	BooksAPI         = HTTPAPI{Path: "/books/{id}", Method: "GET"}
)

func newRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(ProductAPI.ServeMuxPattern(), ProductHandler)
	r.HandleFunc(BooksCategoryAPI.ServeMuxPattern(), BooksCategoryHandler)
	r.HandleFunc(BooksAPI.ServeMuxPattern(), BooksHandler)
	return r
}

func newLogger() *slog.Logger {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	return slog.New(h)
}

func main() {
	logger := newLogger()
	r := newRouter()
	h := httpex.RequestLogMiddleware(logger)(r)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: h,
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
