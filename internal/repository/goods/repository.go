package goods

import (
	"github.com/Masterminds/squirrel"
	"github.com/biryanim/hezzl_tz/internal/client/db"
	"github.com/biryanim/hezzl_tz/internal/repository"
)

type repo struct {
	db db.Client
	qb squirrel.StatementBuilderType
}

func NewRepository(db db.Client) repository.GoodsRepository {
	return &repo{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
