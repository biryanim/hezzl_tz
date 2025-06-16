package log

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/biryanim/hezzl_tz/internal/model"
	"github.com/biryanim/hezzl_tz/internal/repository"
)

const logGoodsTableName = "goods_logs"

type repo struct {
	db clickhouse.Conn
}

func NewRepository(db clickhouse.Conn) repository.LogRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Add(ctx context.Context, entries []model.LogEntry) error {
	if len(entries) == 0 {
		return nil
	}

	query := `INSERT INTO ` + logGoodsTableName

	batch, err := r.db.PrepareBatch(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}

	for _, entry := range entries {
		if err = batch.AppendStruct(&entry); err != nil {
			return fmt.Errorf("failed to append struct to batch: %w", err)
		}
	}

	return batch.Send()
}
