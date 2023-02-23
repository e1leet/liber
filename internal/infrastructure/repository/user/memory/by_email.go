package memory

import (
	"context"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
)

func (r *Memory) UserByEmail(ctx context.Context, email string) (usr *model.Storage, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to get by email") }()

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Email == email {
			r.logger.Info().
				Int("userID", u.ID).
				Str("username", u.Username).
				Msg("user found")

			return u, nil
		}
	}

	r.logger.Info().Err(userError.ErrUserNotFound).Send()

	return nil, userError.ErrUserNotFound
}
