package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/e1leet/liber/pkg/errors"
)

func (r *Postgresql) UserByID(ctx context.Context, id int) (user *model.Storage, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to get by id") }()

	sql, args, err := r.queryBuilder.
		Select().
		Columns(
			model.IDField,
			model.EmailField,
			model.UsernameField,
			model.PasswordField,
			model.CreatedAtField,
			model.UpdatedAtField,
		).From(tableScheme).
		Where(sq.Eq{model.IDField: id}).
		ToSql()

	if err != nil {
		r.logger.Error().Err(err).Send()

		return nil, err
	}

	r.logger.Trace().Str("sql", sql).Interface("args", args).Send()

	return r.selectFrom(ctx, sql, args)
}
