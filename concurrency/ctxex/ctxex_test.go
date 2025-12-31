package ctxex

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoWorkSuccess(t *testing.T) {
	synctest.Run(func() {
		workDuration := 1 * time.Second

		start := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		err := DoWork(ctx, workDuration)
		assert.NoError(t, err)
		assert.Equal(t, 1*time.Second, time.Since(start))
	})
}

func TestDoWorkDeadline(t *testing.T) {
	synctest.Run(func() {
		workDuration := 3 * time.Second

		start := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		err := DoWork(ctx, workDuration)
		assert.Error(t, err, context.DeadlineExceeded)
		assert.Equal(t, 1*time.Second, time.Since(start))
	})
}

func TestDoWorkCancel(t *testing.T) {
	synctest.Run(func() {
		workDuration := 3 * time.Second

		start := time.Now()
		ctx, cancel := context.WithCancel(context.Background())
		time.AfterFunc(1*time.Second, cancel)

		err := DoWork(ctx, workDuration)
		assert.Error(t, err, context.Canceled)
		assert.Equal(t, 1*time.Second, time.Since(start))
	})
}

func TestCancelCause(t *testing.T) {
	causeErr := errors.New("cause")
	ctx, cancel := context.WithCancelCause(t.Context())
	cancel(causeErr)

	assert.ErrorIs(t, ctx.Err(), context.Canceled)
	assert.ErrorIs(t, context.Cause(ctx), causeErr)
}

func TestDeadlineCancelCause(t *testing.T) {
	t.Run("cancel_after_timeout", func(t *testing.T) {
		causeErr := errors.New("timeout cause")
		// set to 0 makes the context timeout immediately
		ctx, cancel := context.WithTimeoutCause(t.Context(), 0, causeErr)
		defer cancel()

		assert.ErrorIs(t, ctx.Err(), context.DeadlineExceeded)
		assert.ErrorIs(t, context.Cause(ctx), causeErr)
		assert.NotErrorIs(t, ctx.Err(), context.Canceled)
	})

	t.Run("cancel_before_timeout", func(t *testing.T) {
		synctest.Run(func() {
			causeErr := errors.New("timeout cause")
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
	})

	t.Run("child_context_timeout_before_parent", func(t *testing.T) {
		synctest.Run(func() {
			parentCauseErr := errors.New("parent timeout cause")
			parentCtx, cancel := context.WithTimeoutCause(context.Background(), 1*time.Second, parentCauseErr)
			defer cancel()

			childCauseErr := errors.New("child timeout cause")
			childCtx, cancel := context.WithTimeoutCause(parentCtx, 0, childCauseErr)
			defer cancel()

			assert.NoError(t, parentCtx.Err())
			assert.NoError(t, context.Cause(parentCtx))

			assert.ErrorIs(t, childCtx.Err(), context.DeadlineExceeded)
			assert.ErrorIs(t, context.Cause(childCtx), childCauseErr)
		})
	})

	t.Run("both_parent_and_child_timeout", func(t *testing.T) {
		synctest.Run(func() {
			parentCauseErr := errors.New("parent timeout cause")
			parentCtx, cancel := context.WithTimeoutCause(context.Background(), 1*time.Second, parentCauseErr)
			defer cancel()

			childCauseErr := errors.New("child timeout cause")
			childCtx, cancel := context.WithTimeoutCause(parentCtx, 0, childCauseErr)
			defer cancel()

			time.Sleep(1500 * time.Millisecond)

			assert.ErrorIs(t, parentCtx.Err(), context.DeadlineExceeded)
			assert.ErrorIs(t, context.Cause(parentCtx), parentCauseErr)
			assert.NotErrorIs(t, context.Cause(parentCtx), childCauseErr)

			assert.ErrorIs(t, childCtx.Err(), context.DeadlineExceeded)
			assert.ErrorIs(t, context.Cause(childCtx), childCauseErr)
		})
	})
}
