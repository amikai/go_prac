package ctxex

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoWorkSuccess(t *testing.T) {
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

func TestCancelCause(t *testing.T) {
	causeErr := errors.New("cause")
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(causeErr)

	assert.ErrorIs(t, ctx.Err(), context.Canceled)
	assert.ErrorIs(t, context.Cause(ctx), causeErr)
}

func TestDeadlineCancelCause(t *testing.T) {
	t.Run("cancel after timeout", func(t *testing.T) {
		causeErr := errors.New("cause")
		ctx, cancel := context.WithTimeoutCause(context.Background(), 0, causeErr)
		defer cancel()

		assert.ErrorIs(t, ctx.Err(), context.DeadlineExceeded)
		assert.ErrorIs(t, context.Cause(ctx), causeErr)
		assert.NotErrorIs(t, ctx.Err(), context.Canceled)
	})

	t.Run("cancel before timeout", func(t *testing.T) {
		causeErr := errors.New("cause")
		ctx, cancel := context.WithTimeoutCause(context.Background(), 1*time.Second, causeErr)
		cancel()

		assert.ErrorIs(t, ctx.Err(), context.Canceled)
		assert.NotErrorIs(t, ctx.Err(), context.DeadlineExceeded)
		assert.NotErrorIs(t, context.Cause(ctx), causeErr)

		time.Sleep(1500 * time.Millisecond)

		assert.ErrorIs(t, ctx.Err(), context.Canceled)
		assert.NotErrorIs(t, ctx.Err(), context.DeadlineExceeded)
		assert.NotErrorIs(t, context.Cause(ctx), causeErr)
	})

	t.Run("verify child or parent context I", func(t *testing.T) {
		parentCauseErr := errors.New("parent cause")
		parentCtx, cancel := context.WithTimeoutCause(context.Background(), 1*time.Second, parentCauseErr)
		defer cancel()

		childCauseErr := errors.New("child cause")
		childCtx, cancel := context.WithTimeoutCause(parentCtx, 0, childCauseErr)
		defer cancel()

		assert.NoError(t, parentCtx.Err())
		assert.NoError(t, context.Cause(parentCtx))

		assert.ErrorIs(t, childCtx.Err(), context.DeadlineExceeded)
		assert.ErrorIs(t, context.Cause(childCtx), childCauseErr)
	})

	t.Run("verify child or parent context II", func(t *testing.T) {
		parentCauseErr := errors.New("parent cause")
		parentCtx, cancel := context.WithTimeoutCause(context.Background(), 1*time.Second, parentCauseErr)
		defer cancel()

		childCauseErr := errors.New("child cause")
		childCtx, cancel := context.WithTimeoutCause(parentCtx, 0, childCauseErr)
		defer cancel()

		time.Sleep(1500 * time.Millisecond)

		assert.ErrorIs(t, parentCtx.Err(), context.DeadlineExceeded)
		assert.ErrorIs(t, context.Cause(parentCtx), parentCauseErr)
		assert.NotErrorIs(t, context.Cause(parentCtx), childCauseErr)

		assert.ErrorIs(t, childCtx.Err(), context.DeadlineExceeded)
		assert.ErrorIs(t, context.Cause(childCtx), childCauseErr)
	})
}
