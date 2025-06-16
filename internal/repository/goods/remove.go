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

func (r *repo) RemoveGood(ctx context.Context, id, projectId int) (*model.Good, error) {
	query, args, err := r.qb.
		Update("goods").
		Set("removed", true).
		Where(squirrel.Eq{
			"id":         id,
			"project_id": projectId,
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
