package shutdown

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloser_Add(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		closer := &Closer{}
		assert.Equal(t, 0, len(closer.funcs))
		closer.Add(func(ctx context.Context) error {
			return nil
		})
		assert.Equal(t, 1, len(closer.funcs))
	})
}

func TestCloser_Close(t *testing.T) {
	t.Run("deadline_exceeded", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()

		closer := &Closer{}
		closer.Add(func(ctx context.Context) error {
			time.Sleep(time.Second)
			return nil
		})
		err := closer.Close(ctx)
		require.ErrorIs(t, err, context.DeadlineExceeded)
	})

	t.Run("close_func_throws_error", func(t *testing.T) {
		const errorText = "empty error"
		closer := &Closer{}
		closer.Add(func(ctx context.Context) error {
			return errors.New(errorText)
		})
		err := closer.Close(context.Background())
		require.ErrorContains(t, err, errorText)
	})

	t.Run("close", func(t *testing.T) {
		closer := &Closer{}
		closer.Add(func(ctx context.Context) error {
			return nil
		})
		err := closer.Close(context.Background())
		require.NoError(t, err)
	})
}
