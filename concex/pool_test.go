package concex

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/sourcegraph/conc/pool"
	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	var p *pool.Pool = pool.New().WithMaxGoroutines(runtime.NumCPU())

	for _, url := range searchURLs {
		url := url
		// Be careful!!! if goroutines in pool are busy, p.Go will be blocked
		p.Go(func() {
			_, _ = fakeSearch(url)
		})
	}
	p.Wait()
}

func TestErrPool(t *testing.T) {
	// ErrPool will wait all task done, and join (errors.Join) the error
	var p *pool.ErrorPool = pool.New().WithErrors().WithMaxGoroutines(runtime.NumCPU())

	for i, url := range searchURLs {
		i, url := i, url
		p.Go(func() error {
			if i == 0 {
				time.Sleep(time.Second)
			}
			_, err := fakeSearchMustErr(url)
			return err
		})
	}
	retErr := p.Wait()

	for _, url := range searchURLs {
		assert.ErrorIs(t, retErr, searchErr(url))
	}
}

func TestContextPoolSuccess(t *testing.T) {
	var p *pool.ContextPool = pool.New().WithContext(context.Background()).WithMaxGoroutines(runtime.NumCPU())

	for _, url := range searchURLs {
		url := url
		p.Go(func(ctx context.Context) error {
			_, err := fakeSearchCtx(ctx, url)
			if err != nil {
				return err
			}
			return nil
		})
	}
	err := p.Wait()
	assert.NoError(t, err)
}

func TestContextPoolTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()
	// ContextPool runs tasks that take a context, and join (errors.Join) the error
	var p *pool.ContextPool = pool.New().WithContext(ctx).WithMaxGoroutines(runtime.NumCPU())

	for _, url := range searchURLs {
		url := url
		p.Go(func(ctx context.Context) error {
			_, err := fakeSearchCtxWithDuration(ctx, url, 1*time.Second)
			if err != nil {
				// searchErr will wrap err
				return searchErr(url, err)
			}
			return nil
		})
	}

	retErr := p.Wait()
	assert.ErrorIs(t, retErr, context.DeadlineExceeded)
	for _, url := range searchURLs {
		assert.ErrorIs(t, retErr, searchErr(url))
	}
}

func TestContextPoolCancelOnError(t *testing.T) {
	// ContextPool runs tasks that take a context, and join (errors.Join) the error
	// WithCancelOnError() setting will cancal context when one task return error
	var p *pool.ContextPool = pool.New().
		WithContext(context.Background()).
		WithCancelOnError().
		WithMaxGoroutines(runtime.NumCPU())

	failedIndex := 3
	for i, url := range searchURLs {
		i, url := i, url
		p.Go(func(ctx context.Context) error {
			if i == failedIndex {
				_, err := fakeSearchMustErr(url)
				return err
			}

			_, err := fakeSearchCtxWithDuration(ctx, url, 1*time.Second)
			if err != nil {
				return err
			}
			return nil
		})
	}

	retErr := p.Wait()
	assert.ErrorIs(t, retErr, context.Canceled)
	assert.ErrorIs(t, retErr, searchErr(searchURLs[failedIndex]))
}
