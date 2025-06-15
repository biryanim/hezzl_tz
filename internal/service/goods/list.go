package goods

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/model"
)

func (s *serv) List(ctx context.Context, goodsListingParams *model.GoodListParams) (*model.GoodsList, error) {
	var goods *model.GoodsList
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		goods, err = s.goodsRepository.List(ctx, goodsListingParams)
		return err
	})
	if err != nil {
		return nil, err
	}

	return goods, nil
}
