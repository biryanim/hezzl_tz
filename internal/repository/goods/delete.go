package goods

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperrors "github.com/biryanim/hezzl_tz/internal/errors"
	"github.com/biryanim/hezzl_tz/internal/model"
)

func (r *repo) Delete(ctx context.Context, good *model.GoodDRemoveParams) error {
	query, args, err := r.qb.
		Delete("goods").
		Where(squirrel.Eq{"id": good.ID}).
		Where(squirrel.Eq{"project_id": good.ProjectID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build delete query: %w", err)
	}

	tag, err := r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return apperrors.ErrGoodsNotFound
	}

	return nil
}
