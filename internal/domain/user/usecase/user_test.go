package usecase

import (
	"context"
	"testing"
	"time"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsecase_User(t *testing.T) {
	var (
		repository = &userRepositoryMock{}
		manager    = &passwordManagerMock{}
		usecase    = New(repository, manager, log.Logger)
	)

	var (
		ctx = context.Background()
		id  = 1
	)

	t.Run("user_not_found", func(t *testing.T) {
		repository.On("UserByID", ctx, id).
			Return(&model.Storage{}, userError.ErrUserNotFound).Once()

		_, err := usecase.User(ctx, id)
		require.ErrorIs(t, err, userError.ErrUserNotFound)
	})

	t.Run("unexpected_error", func(t *testing.T) {
		unexpected := errors.New("error")

		repository.On("UserByID", ctx, id).
			Return(&model.Storage{}, unexpected).Once()

		_, err := usecase.User(ctx, id)
		require.ErrorIs(t, err, unexpected)
	})

	t.Run("success", func(t *testing.T) {
		expected := &model.User{
			ID:        1,
			Email:     "test",
			Username:  "test",
			Password:  "test",
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		repository.On("UserByID", ctx, id).
			Return(&model.Storage{
				ID:        expected.ID,
				Email:     expected.Email,
				Username:  expected.Username,
				Password:  expected.Password,
				CreatedAt: expected.CreatedAt,
				UpdatedAt: expected.UpdatedAt,
			}, nil).Once()

		user, err := usecase.User(ctx, id)
		require.NoError(t, err)

		assert.EqualValues(t, expected, user)
	})
}
