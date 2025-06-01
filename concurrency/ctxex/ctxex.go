package ctxex

import (
	"context"
	"time"
)

func DoWork(ctx context.Context, duration time.Duration) error {
	select {
	case <-time.After(duration):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
