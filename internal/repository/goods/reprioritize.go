package goods

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperrors "github.com/biryanim/hezzl_tz/internal/errors"
	"github.com/biryanim/hezzl_tz/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *repo) Reprioritize(ctx context.Context, good *model.GoodReprioritizeParams) ([]*model.Good, error) {
	query, args, err := r.qb.
		Select("priority").
		From("goods").
		Where(squirrel.Eq{"id": good.ID, "project_id": good.ProjectID, "removed": false}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	var curPriority int
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(&curPriority)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("goods not found: %w", err)
		}
		return nil, fmt.Errorf("failed to query goods: %w", err)
	}

	if good.NewPriority > curPriority {
		query, args, err = r.qb.
			Update("goods").
			Set("priority", squirrel.Expr("priority - 1")).
			Where(squirrel.And{
				squirrel.Eq{"project_id": good.ProjectID},
				squirrel.Gt{"priority": curPriority},
				squirrel.LtOrEq{"priority": good.NewPriority},
				squirrel.Eq{"removed": false},
			}).Suffix("RETURNING *").ToSql()
	} else if good.NewPriority < curPriority {
		query, args, err = r.qb.
			Update("goods").
			Set("priority", squirrel.Expr("priority + 1")).
			Where(squirrel.And{
				squirrel.Eq{"project_id": good.ProjectID},
				squirrel.GtOrEq{"priority": good.NewPriority},
				squirrel.Lt{"priority": curPriority},
				squirrel.Eq{"removed": false},
			}).
			Suffix("RETURNING *").
			ToSql()
	} else {
		query, args, err = r.qb.
			Select("id", "project_id", "name", "description", "priority", "removed", "created_at").
			From("goods").
			Where(squirrel.Eq{"id": good.ID, "project_id": good.ProjectID}).ToSql()
		if err != nil {
			return nil, fmt.Errorf("failed to build select query: %w", err)
		}
		var g model.Good
		err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(
			&g.ID,
			&g.ProjectID,
			&g.Info.Name,
			&g.Info.Description,
			&g.Priority,
			&g.Removed,
			&g.CreatedAt,
		)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, apperrors.ErrGoodsNotFound
			}
			return nil, fmt.Errorf("failed to get good: %w", err)
		}
		return []*model.Good{&g}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to build update query: %w", err)
	}

	rows, err := r.db.DB().QueryContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("goods not found: %w", err)
		}
		return nil, fmt.Errorf("failed to query rows: %w", err)
	}
	defer rows.Close()

	var updated []*model.Good

	for rows.Next() {
		var g model.Good
		if err = rows.Scan(
			&g.ID,
			&g.ProjectID,
			&g.Info.Name,
			&g.Info.Description,
			&g.Priority,
			&g.Removed,
			&g.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		updated = append(updated, &g)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate rows: %w", err)
	}

	query, args, err = r.qb.
		Update("goods").
		Set("priority", good.NewPriority).
		Where(squirrel.Eq{
			"id":         good.ID,
			"project_id": good.ProjectID,
		}).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build update query: %w", err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update goods: %w", err)
	}

	return updated, nil
}
