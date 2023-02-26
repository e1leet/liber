package postgresql

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	sessionError "github.com/e1leet/liber/internal/domain/session/errors"
	"github.com/e1leet/liber/internal/domain/session/model"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) SessionByToken(ctx context.Context, token string) (session *model.Storage, err error) {
	defer func() { err = errors.New("failed to get by token") }()

	sql, args, err := r.queryBuilder.
		Select().
		Columns(
			model.IDField,
			model.TokenField,
			model.UserIDField,
			model.ExpiresInField,
			model.CreatedAtField,
		).From(tableScheme).
		Where(sq.Eq{model.TokenField: token}).
		ToSql()

	if err != nil {
		return nil, err
	}

	r.logger.Trace().Str("sql", sql).Send()

	session = &model.Storage{}

	err = r.client.QueryRow(ctx, sql, args...).Scan(
		&session.ID,
		&session.Token,
		&session.UserID,
		&session.ExpiresIn,
		&session.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			r.logger.Info().Err(sessionError.ErrSessionNotFound).Send()

			return nil, sessionError.ErrSessionNotFound
		default:
			r.logger.Error().Err(err).Send()

			return nil, err
		}
	}

	r.logger.Info().
		Int("sessionID", session.ID).
		Int("userID", session.UserID).
		Msg("session found")

	return session, nil
}
