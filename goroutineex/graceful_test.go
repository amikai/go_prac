package goroutineex

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGracefulRunReturnEarly(t *testing.T) {
	time.AfterFunc(100*time.Millisecond, func() {
		p, err := os.FindProcess(os.Getpid())
		assert.NoError(t, err)
		err = p.Signal(syscall.SIGINT)
		assert.NoError(t, err)

	})

	errCause := errors.New("cause")
	err := GracefulRun(func(ctx context.Context) error {
		return errCause
	}, nil)
	assert.Equal(t, err, errCause)

}

func TestGracefulSignal(t *testing.T) {
	time.AfterFunc(100*time.Millisecond, func() {
		p, err := os.FindProcess(os.Getpid())
		assert.NoError(t, err)
		err = p.Signal(syscall.SIGINT)
		assert.NoError(t, err)

	})

	err := GracefulRun(func(ctx context.Context) error {
		<-ctx.Done()
		return nil
	}, nil)
	assert.NoError(t, err)
}

func TestGracefullyShudownTimeout(t *testing.T) {
	time.AfterFunc(100*time.Millisecond, func() {
		p, err := os.FindProcess(os.Getpid())
		assert.NoError(t, err)
		err = p.Signal(syscall.SIGINT)
		assert.NoError(t, err)

	})

	err := GracefulRun(func(ctx context.Context) error {
		<-ctx.Done()
		return nil
	}, &GracefulConfig{Timeout: 200 * time.Millisecond})
	assert.NoError(t, err)
}

func TestGracefullyShudownToolLong(t *testing.T) {
	time.AfterFunc(100*time.Millisecond, func() {
		p, err := os.FindProcess(os.Getpid())
		assert.NoError(t, err)
		err = p.Signal(syscall.SIGINT)
		assert.NoError(t, err)

	})

	err := GracefulRun(func(ctx context.Context) error {
		<-ctx.Done()
		time.Sleep(300 * time.Millisecond)
		return nil
	}, &GracefulConfig{Timeout: 200 * time.Millisecond})
	assert.ErrorIs(t, err, ErrGracefullyTimeout)
}

func TestHttpGracefullyShudown(t *testing.T) {
	time.AfterFunc(100*time.Millisecond, func() {
		p, err := os.FindProcess(os.Getpid())
		assert.NoError(t, err)
		err = p.Signal(syscall.SIGINT)
		assert.NoError(t, err)

	})

	err := GracefulRun(func(ctx context.Context) error {
		srv := http.Server{
			Addr: ":0",
		}

		go func() {
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()
		<-ctx.Done()
		return srv.Shutdown(context.Background())
	}, nil)
	assert.NoError(t, err)
}
