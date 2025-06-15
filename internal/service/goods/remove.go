package goods

import (
	"context"
	"fmt"
	"github.com/biryanim/hezzl_tz/internal/model"
)

func (s *serv) Remove(ctx context.Context, goodsRemovingParams *model.GoodDRemoveParams) (*model.GoodRemove, error) {
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		err = s.goodsRepository.RemoveGood(ctx, goodsRemovingParams.ID, goodsRemovingParams.ProjectID)
		return err
	})
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("goods:%d:%d", goodsRemovingParams.ID, goodsRemovingParams.ProjectID)
	err = s.cache.Delete(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("failed to remove goods: %w", err)
	}

	return &model.GoodRemove{
		ID:        goodsRemovingParams.ID,
		ProjectID: goodsRemovingParams.ProjectID,
		Removed:   true,
	}, nil
}
