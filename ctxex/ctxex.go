package ctxex

import (
	"context"
	"time"
)

var workDuration time.Duration = 5 * time.Second

func DoWork(ctx context.Context) error {
	select {
	case <-time.After(workDuration):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
