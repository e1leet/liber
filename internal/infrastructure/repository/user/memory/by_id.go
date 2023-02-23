package memory

import (
	"context"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
)

func (r *Memory) UserByID(ctx context.Context, id int) (usr *model.Storage, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to get by id") }()

	r.mu.Lock()
	defer r.mu.Unlock()

	usr, ok := r.users[id]
	if !ok {
		r.logger.Info().Err(userError.ErrUserNotFound).Send()

		return nil, userError.ErrUserNotFound
	}

	r.logger.Info().
		Int("userID", usr.ID).
		Str("username", usr.Username).
		Msg("user found")

	return usr, nil
}
