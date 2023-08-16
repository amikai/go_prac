package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"golang.org/x/sync/errgroup"

	"github.com/amikai/go_prac/httpex"
)

func main() {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(h)
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	// TODO: add the middleware for recovery
	r.Use(adapter.Wrap(httpex.RequestLogMiddleware(logger)))
	r.GET("/products/:id", ProductHandler)
	r.GET("/books/:category/", BooksCategoryHandler)
	r.GET("/books/:category/:id", BooksHandler)

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
