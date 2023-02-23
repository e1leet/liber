package usecase

import (
	"context"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
)

func (u *Usecase) User(ctx context.Context, id int) (user *model.User, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to get user") }()

	storage, err := u.userRepository.UserByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, userError.ErrUserNotFound):
			u.logger.Warn().Err(err).Send()

			return nil, err
		default:
			u.logger.Error().Err(err).Send()

			return nil, err
		}
	}

	return storage.ToDomain(), nil
}
