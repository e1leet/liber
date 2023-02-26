package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	sessionError "github.com/e1leet/liber/internal/domain/session/errors"
	"github.com/e1leet/liber/internal/domain/session/model"
	"github.com/e1leet/liber/pkg/errors"
)

func (r *Repository) DeleteByToken(ctx context.Context, token string) (err error) {
	defer func() { err = errors.New("failed to delete by token") }()

	sql, args, err := r.queryBuilder.
		Delete(tableScheme).
		Where(sq.Eq{model.TokenField: token}).
		ToSql()
	if err != nil {
		r.logger.Error().Err(err).Send()

		return err
	}

	r.logger.Trace().Str("sql", sql).Send()

	if exec, err := r.client.Exec(ctx, sql, args); err != nil {
		r.logger.Error().Err(err).Send()

		return err
	} else if !exec.Delete() || exec.RowsAffected() == 0 {
		r.logger.Warn().Err(sessionError.ErrSessionDeleteFailed).Send()

		return sessionError.ErrSessionDeleteFailed
	}

	return nil
}
