package usecase

import (
	"context"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
)

func (u *Usecase) Register(ctx context.Context, dto model.RegisterDTO) (user *model.User, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to register") }()

	u.logger.Info().Str("username", dto.Username).Send()

	m := model.NewCreateMapper(
		dto.Email,
		dto.Password,
		u.passwordManager.Hash(dto.Password),
	)

	storage, err := u.userRepository.Create(ctx, m)
	if err != nil {
		switch {
		case errors.Is(err, userError.ErrEmailAlreadyExists), errors.Is(err, userError.ErrUsernameAlreadyExists):
			u.logger.Info().Err(err).Send()

			return nil, err
		default:
			u.logger.Error().Err(err).Send()

			return nil, err
		}
	}

	return storage.ToDomain(), nil
}
