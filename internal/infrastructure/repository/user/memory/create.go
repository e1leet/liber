package memory

import (
	"context"
	"time"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
)

func (r *Memory) Create(ctx context.Context, m map[string]interface{}) (usr *model.Storage, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to create user") }()

	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Info().Msg("create user")

	usr, err = model.NewStorageFromCreateMapper(m)
	if err != nil {
		r.logger.Error().Err(err).Send()

		return nil, err
	}

	for _, u := range r.users {
		if u.Email == usr.Email {
			r.logger.Info().Err(userError.ErrEmailAlreadyExists).Send()

			return nil, userError.ErrEmailAlreadyExists
		}

		if u.Username == usr.Username {
			r.logger.Info().Err(userError.ErrUsernameAlreadyExists).Send()

			return nil, userError.ErrUsernameAlreadyExists
		}
	}

	usr.ID = len(r.users) + 1
	usr.CreatedAt = time.Now().UTC()
	usr.UpdatedAt = time.Now().UTC()

	r.users[usr.ID] = usr

	r.logger.Info().
		Int("userID", usr.ID).
		Str("username", usr.Username).
		Msg("user created")

	return
}
