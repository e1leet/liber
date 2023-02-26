package postgresql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog"
)

const (
	scheme      = "public"
	table       = "user_session"
	tableScheme = scheme + "." + table
)

type Repository struct {
	client       postgresqlClient
	queryBuilder sq.StatementBuilderType
	logger       zerolog.Logger
}

func New(client postgresqlClient, logger zerolog.Logger) *Repository {
	return &Repository{
		client:       client,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		logger:       logger,
	}
}
