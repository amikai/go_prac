package goroutineex

import (
	"context"
	"errors"
	"os/signal"
	"syscall"
	"time"
)

// Copy from https://github.com/NTHU-LSALAB/NTHU-Distributed-System/blob/fe7ef4e110874817ee092ab9776fa0500d0cc44f/pkg/runkit/graceful.go#L22

var (
	ErrGracefullyTimeout = errors.New("gracefully shutdown timeout")
)

type GracefulConfig struct {
	Timeout time.Duration
}

func GracefulRun(runFunc func(context.Context) error, config *GracefulConfig) error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	errCh := make(chan error)
	go func() {
		errCh <- runFunc(ctx)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		if config == nil {
			return <-errCh
		}

		select {
		case err := <-errCh:
			// gracefully shutdown
			return err
		case <-time.After(config.Timeout):
			// timeout shutdown
			return ErrGracefullyTimeout
		}

	}
}
