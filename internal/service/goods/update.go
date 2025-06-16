package goods

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/model"
	"log"
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

	if err = s.cache.DeleteByPattern(ctx, "goods:list:*"); err != nil {
		log.Printf("delete goods cache err: %v", err)
	}

	s.publishLogEvent(ctx, res)

	return res, nil
}
