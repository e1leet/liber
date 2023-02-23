package usecase

import (
	"context"
	"io"
	"testing"
	"time"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	log.Logger = log.Output(io.Discard)
}

func TestUsecase_Register(t *testing.T) {
	var (
		repository = &userRepositoryMock{}
		manager    = &passwordManagerMock{}
		usecase    = New(repository, manager, log.Logger)
	)

	var (
		ctx = context.Background()
		dto = model.RegisterDTO{
			Email:    "test@test.com",
			Username: "test",
			Password: "test",
		}
		hash = "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
		m    = model.NewCreateMapper(dto.Email, dto.Username, hash)
	)

	manager.On("Hash", dto.Password).
		Return(hash)

	t.Run("email_already_exists", func(t *testing.T) {
		repository.On("Create", ctx, m).
			Return(&model.Storage{}, userError.ErrEmailAlreadyExists).Once()

		_, err := usecase.Register(ctx, dto)
		require.ErrorIs(t, err, userError.ErrEmailAlreadyExists)
	})

	t.Run("username_already_exists", func(t *testing.T) {
		repository.On("Create", ctx, m).
			Return(&model.Storage{}, userError.ErrUsernameAlreadyExists).Once()

		_, err := usecase.Register(ctx, dto)
		require.ErrorIs(t, err, userError.ErrUsernameAlreadyExists)
	})

	t.Run("unexpected_error", func(t *testing.T) {
		unexpected := errors.New("error")
		repository.On("Create", ctx, m).
			Return(&model.Storage{}, unexpected).Once()

		_, err := usecase.Register(ctx, dto)
		require.ErrorIs(t, err, unexpected)
	})

	t.Run("success", func(t *testing.T) {
		expected := &model.User{
			ID:        1,
			Email:     dto.Email,
			Username:  dto.Username,
			Password:  dto.Password,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}
		repository.On("Create", ctx, m).
			Return(&model.Storage{
				ID:        expected.ID,
				Email:     expected.Email,
				Username:  expected.Username,
				Password:  expected.Password,
				CreatedAt: expected.CreatedAt,
				UpdatedAt: expected.UpdatedAt,
			}, nil).Once()

		user, err := usecase.Register(ctx, dto)
		require.NoError(t, err)

		assert.EqualValues(t, expected, user)
	})
}
