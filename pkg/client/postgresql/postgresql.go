package postgresql

import (
	"context"
	"time"

	"github.com/e1leet/liber/pkg/errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewClient(ctx context.Context, uri string, pingTimeout time.Duration) (pool *pgxpool.Pool, err error) {
	defer func() { err = errors.WrapIfErr(err, "failed to connect to postgresql") }()

	pool, err = pgxpool.New(ctx, uri)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
