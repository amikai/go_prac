package ctxex

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoWorkSucess(t *testing.T) {
	workDuration := 3 * time.Millisecond
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Millisecond)
	defer cancel()

	err := DoWork(ctx, workDuration)
	assert.NoError(t, err)
}

func TestDoWorkDeadline(t *testing.T) {
	workDuration := 3 * time.Millisecond
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	defer cancel()

	err := DoWork(ctx, workDuration)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestDoWorkCancel(t *testing.T) {
	workDuration := 3 * time.Millisecond
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(1 * time.Millisecond)
		cancel()
	}()
	err := DoWork(ctx, workDuration)
	assert.ErrorIs(t, err, context.Canceled)
}
