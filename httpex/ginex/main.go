package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(h)
	gin.SetMode(gin.DebugMode)
	// Default router has gin builtin logger and recovery middleware
	// In production code, should use gin.New() with custom logger and recovery middleware manually
	r := gin.Default()
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
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("gracefully quit gorilla server")
	}
}
