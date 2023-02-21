package memory

import (
	"context"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
)

func (r *Memory) UserByUsername(ctx context.Context, username string) (usr *model.Storage, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to get by username") }()

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Username == username {
			r.logger.Info().
				Int("userID", u.ID).
				Str("username", u.Username).
				Msg("user found")

			return u, nil
		}
	}

	r.logger.Info().
		Str("username", username).
		Err(userError.ErrUserNotFound).Send()

	return nil, userError.ErrUserNotFound
}
