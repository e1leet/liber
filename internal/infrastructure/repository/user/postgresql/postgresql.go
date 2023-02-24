package postgresql

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/rs/zerolog"
)

const (
	scheme      = "public"
	table       = "usr"
	tableScheme = scheme + "." + table
)

var (
	emailConstraintName    = fmt.Sprintf("%s_%s_key", table, model.EmailField)
	usernameConstraintName = fmt.Sprintf("%s_%s_key", table, model.UsernameField)
)

// TODO Write unit test for methods
type Postgresql struct {
	client       postgresqlClient
	queryBuilder sq.StatementBuilderType
	logger       zerolog.Logger
}

func New(client postgresqlClient, logger zerolog.Logger) *Postgresql {
	return &Postgresql{
		client:       client,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		logger:       logger,
	}
}
