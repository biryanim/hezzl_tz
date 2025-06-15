package goods

import (
	"context"
	"fmt"
	"github.com/biryanim/hezzl_tz/internal/model"
)

func (s *serv) Update(ctx context.Context, goodsUpdatingParams *model.GoodUpdateParams) (*model.Good, error) {
	var (
		res *model.Good
	)
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		res, err = s.goodsRepository.Update(ctx, goodsUpdatingParams)
		return err
	})
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("good:%d:%d", goodsUpdatingParams.ID, goodsUpdatingParams.ProjectID)
	err = s.cache.Delete(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("failed to delete key: %w", err)
	}
	return res, nil
}
