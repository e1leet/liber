package postgresql

import (
	"context"
	"fmt"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Postgresql) Create(ctx context.Context, m map[string]interface{}) (user *model.Storage, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to create user") }()

	user, err = model.NewStorageFromCreateMapper(m)
	if err != nil {
		r.logger.Error().Err(err).Send()

		return nil, err
	}

	sql, args, err := r.queryBuilder.
		Insert(tableScheme).
		SetMap(m).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s", model.IDField, model.CreatedAtField, model.UpdatedAtField)).
		ToSql()

	if err != nil {
		r.logger.Error().Str("username", user.Username).Err(err).Send()

		return nil, err
	}

	r.logger.Trace().Str("sql", sql).Send()

	if err = r.client.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt); err != nil {
		var pgErr *pgconn.PgError

		switch {
		case errors.As(err, &pgErr):
			if pgErr.Code == pgerrcode.UniqueViolation && pgErr.ConstraintName == emailConstraintName {
				r.logger.Info().Err(userError.ErrEmailAlreadyExists).Send()

				return nil, userError.ErrEmailAlreadyExists
			}

			if pgErr.Code == pgerrcode.UniqueViolation && pgErr.ConstraintName == usernameConstraintName {
				r.logger.Info().Err(userError.ErrUsernameAlreadyExists).Send()

				return nil, userError.ErrUsernameAlreadyExists
			}

			fallthrough
		default:
			r.logger.Error().Err(err).Send()

			return nil, nil
		}
	}

	return user, nil
}
