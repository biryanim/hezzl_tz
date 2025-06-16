package goods

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/model"
	"log"
)

func (s *serv) Remove(ctx context.Context, goodsRemovingParams *model.GoodDRemoveParams) (*model.GoodRemove, error) {
	var good *model.Good
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		good, err = s.goodsRepository.RemoveGood(ctx, goodsRemovingParams.ID, goodsRemovingParams.ProjectID)
		return err
	})
	if err != nil {
		return nil, err
	}

	if err = s.cache.DeleteByPattern(ctx, "goods:list:*"); err != nil {
		log.Printf("delete goods cache err: %v", err)
	}

	s.publishLogEvent(ctx, good)

	return &model.GoodRemove{
		ID:        good.ID,
		ProjectID: good.ProjectID,
		Removed:   true,
	}, nil
}
