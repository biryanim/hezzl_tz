package goods

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/biryanim/hezzl_tz/internal/model"
)

func (r *repo) List(ctx context.Context, good *model.GoodListParams) (*model.GoodsList, error) {
	query, args, err := r.qb.
		Select("g.id",
			"g.project_id",
			"g.name",
			"g.description",
			"g.priority",
			"g.removed",
			"g.created_at",
			"total.total", "removed.removed",
		).
		From("goods g").
		JoinClause(
			"CROSS JOIN (SELECT COUNT(*) as total FROM goods) total",
		).
		JoinClause(
			"CROSS JOIN (SELECT COUNT(*) as removed FROM goods WHERE removed=true) removed",
		).
		Where(squirrel.Eq{"g.removed": false}).
		OrderBy("g.priority").
		Limit(uint64(good.Limit)).
		Offset(uint64(good.Offset)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	rows, err := r.db.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list goods: %w", err)
	}
	defer rows.Close()

	var goods model.GoodsList
	var total, removed int

	for rows.Next() {
		var g model.Good
		err = rows.Scan(
			&g.ID, &g.ProjectID, &g.Info.Name, &g.Info.Description, &g.Priority, &g.Removed, &g.CreatedAt, &total, &removed)
		if err != nil {
			return nil, fmt.Errorf("failed to list goods: %w", err)
		}

		goods.Goods = append(goods.Goods, g)
	}

	goods.MetaInfo = model.Meta{
		Total:   total,
		Removed: removed,
		Limit:   good.Limit,
		Offset:  good.Offset,
	}

	return &goods, nil
}
