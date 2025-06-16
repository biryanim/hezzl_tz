package goods

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/model"
	"log"
)

func (s *serv) Reprioritize(ctx context.Context, goodReprioritizingParams *model.GoodReprioritizeParams) (*model.GoodsPrioritize, error) {
	var goods *model.GoodsPrioritize
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		goods, err = s.goodsRepository.Reprioritize(ctx, goodReprioritizingParams)
		return err
	})
	if err != nil {
		return nil, err
	}

	if err = s.cache.DeleteByPattern(ctx, "goods:list:*"); err != nil {
		log.Printf("delete goods cache err: %v", err)
	}

	return goods, nil
}
