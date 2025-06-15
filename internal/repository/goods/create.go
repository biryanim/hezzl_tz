package goods

import (
	"context"
	"fmt"

	apperrors "github.com/biryanim/hezzl_tz/internal/errors"
	"github.com/biryanim/hezzl_tz/internal/model"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

func (r *repo) Create(ctx context.Context, good *model.GoodCreateParams) (*model.Good, error) {
	query, args, err := r.qb.
		Insert("goods").
		Columns("project_id", "name", "description").
		Values(good.ProjectID, good.Info.Name, good.Info.Description).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build insert query: %w", err)
	}

	var res model.Good
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(
		&res.ID,
		&res.ProjectID,
		&res.Info.Name,
		&res.Info.Description,
		&res.Removed,
		&res.CreatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, apperrors.ErrGoodsAlreadyExists
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &res, nil
}
