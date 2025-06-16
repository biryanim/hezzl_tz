package goods

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/model"
	"log"
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

	if err = s.cache.DeleteByPattern(ctx, "goods:list:*"); err != nil {
		log.Printf("delete goods cache err: %v", err)
	}

	return &model.GoodRemove{
		ID:        goodsRemovingParams.ID,
		ProjectID: goodsRemovingParams.ProjectID,
		Removed:   true,
	}, nil
}
