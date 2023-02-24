package postgresql

import (
	"context"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
	"github.com/jackc/pgx/v5"
)

func (r *Postgresql) selectFrom(ctx context.Context, sql string, args []interface{}) (user *model.Storage, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to select user") }()

	user = &model.Storage{}

	err = r.client.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			r.logger.Info().Err(userError.ErrUserNotFound).Send()

			return nil, userError.ErrUserNotFound
		default:
			r.logger.Error().Err(err).Send()

			return nil, err
		}
	}

	return user, nil
}
