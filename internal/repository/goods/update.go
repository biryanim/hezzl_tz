package goods

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperrors "github.com/biryanim/hezzl_tz/internal/errors"
	"github.com/biryanim/hezzl_tz/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *repo) Update(ctx context.Context, goodsUpdate *model.GoodUpdateParams) (*model.Good, error) {
	query, args, err := r.qb.
		Update("goods").
		Set("name", goodsUpdate.Info.Name).
		Set("description", goodsUpdate.Info.Description).
		Where(squirrel.Eq{
			"id":         goodsUpdate.ID,
			"project_id": goodsUpdate.ProjectID,
			"removed":    false,
		}).Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build update query: %w", err)
	}

	var good model.Good
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(
		&good.ID,
		&good.ProjectID,
		&good.Info.Name,
		&good.Info.Description,
		&good.Priority,
		&good.Removed,
		&good.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrGoodsNotFound
		}
		return nil, fmt.Errorf("failed to update goods: %w", err)
	}

	return &good, nil
}
