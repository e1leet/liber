package postgresql

import (
	"context"
	"fmt"

	"github.com/e1leet/liber/internal/domain/session/model"
	"github.com/e1leet/liber/pkg/errors"
)

func (r *Repository) Create(ctx context.Context, m map[string]interface{}) (session *model.Storage, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to create session") }()

	session, err = model.NewStorageFromCreateMapper(m)
	if err != nil {
		r.logger.Error().Err(err).Send()

		return nil, err
	}

	sql, args, err := r.queryBuilder.
		Insert(tableScheme).
		SetMap(m).
		Suffix(fmt.Sprintf("RETURNING %s, %s", model.IDField, model.CreatedAtField)).
		ToSql()

	if err != nil {
		r.logger.Error().Err(err).Int("userID", session.UserID).Send()

		return nil, err
	}

	r.logger.Trace().Str("sql", sql).Send()

	if err := r.client.QueryRow(ctx, sql, args...).Scan(&session.ID, &session.CreatedAt); err != nil {
		r.logger.Error().Err(err).Send()

		return nil, err
	}

	r.logger.Info().Int("sessionID", session.ID).Int("userID", session.UserID).Send()

	return session, nil
}
