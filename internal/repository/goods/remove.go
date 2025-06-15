package goods

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r *repo) RemoveGood(ctx context.Context, id, projectId int) error {
	query, args, err := r.qb.
		Update("goods").
		Set("removed", true).
		Where(squirrel.Eq{
			"id":         id,
			"project_id": projectId,
			"removed":    false,
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build update query: %w", err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("the project doesn't exist yet: %w", err)
		}
		return fmt.Errorf("failed to execute update query: %w", err)
	}

	return nil
}
