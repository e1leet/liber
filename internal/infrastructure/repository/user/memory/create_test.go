package memory

import (
	"context"
	"io"
	"testing"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	log.Logger = log.Output(io.Discard)
}

func TestMemory_Create(t *testing.T) {
	var (
		ctx = context.Background()
		m   = model.NewCreateMapper("test@test.com", "test", "test")
	)

	t.Run("not_specified", func(t *testing.T) {
		testcases := []struct {
			name          string
			m             map[string]interface{}
			expectedError error
		}{
			{
				name: "email_not_specified",
				m: map[string]interface{}{
					model.UsernameField: "test",
					model.PasswordField: "test",
				},
				expectedError: userError.ErrEmailNotSpecified,
			},
			{
				name: "username_not_specified",
				m: map[string]interface{}{
					model.EmailField:    "test@test.com",
					model.PasswordField: "test",
				},
				expectedError: userError.ErrUsernameNotSpecified,
			},
			{
				name: "password_not_specified",
				m: map[string]interface{}{
					model.EmailField:    "test@test.com",
					model.UsernameField: "test",
				},
				expectedError: userError.ErrPasswordNotSpecified,
			},
		}

		for _, tt := range testcases {
			t.Run(tt.name, func(t *testing.T) {
				repository := New(log.Logger)
				_, err := repository.Create(context.Background(), tt.m)
				require.ErrorIs(t, err, tt.expectedError)
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		repository := New(log.Logger)

		usr, err := repository.Create(ctx, m)
		require.NoError(t, err)

		assert.Equal(t, 1, usr.ID)
		assert.Equal(t, usr.Email, m[model.EmailField])
		assert.Equal(t, usr.Username, m[model.UsernameField])
		assert.Equal(t, usr.Password, m[model.PasswordField])
	})

	t.Run("email_already_exists", func(t *testing.T) {
		repository := New(log.Logger)

		usr, err := repository.Create(ctx, m)
		require.NoError(t, err)

		m2 := model.NewCreateMapper(usr.Email, "test", "test")

		_, err = repository.Create(ctx, m2)
		require.ErrorIs(t, err, userError.ErrEmailAlreadyExists)
	})

	t.Run("username_already_exists", func(t *testing.T) {
		repository := New(log.Logger)

		usr, err := repository.Create(ctx, m)
		require.NoError(t, err)

		m2 := model.NewCreateMapper("new@test.com", usr.Username, "test")

		_, err = repository.Create(ctx, m2)
		require.ErrorIs(t, err, userError.ErrUsernameAlreadyExists)
	})
}
