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

func (r *repo) GetByIds(ctx context.Context, id, projectId int) (*model.Good, error) {
	query, args, err := r.qb.
		Select("id", "project_id", "name", "description", "priority", "removed", "created_at").
		From("goods").
		Where(squirrel.Eq{"id": id}).
		Where(squirrel.Eq{"project_id": projectId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
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
		return nil, fmt.Errorf("failed to get good: %w", err)
	}

	return &good, nil
}
