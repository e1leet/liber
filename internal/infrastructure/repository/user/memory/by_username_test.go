package memory

import (
	"context"
	"io"
	"testing"

	"github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	log.Logger = log.Output(io.Discard)
}

func TestMemory_UserByUsername(t *testing.T) {
	var (
		ctx = context.Background()
	)

	t.Run("user_not_found", func(t *testing.T) {
		repository := New(log.Logger)

		_, err := repository.UserByEmail(ctx, "test@test.com")
		require.ErrorIs(t, err, errors.ErrUserNotFound)
	})

	t.Run("user_found", func(t *testing.T) {
		repository := New(log.Logger)
		m := model.NewCreateMapper("test@test.com", "test", "test")

		usr, err := repository.Create(ctx, m)
		require.NoError(t, err)

		actual, err := repository.UserByUsername(ctx, usr.Username)
		require.NoError(t, err)

		assert.EqualValues(t, usr, actual)
	})
}
