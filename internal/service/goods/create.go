package goods

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/model"
)

func (s *serv) Create(ctx context.Context, goodsCreatingParams *model.GoodCreateParams) (*model.Good, error) {
	res, err := s.goodsRepository.Create(ctx, goodsCreatingParams)
	if err != nil {
		return nil, err
	}

	return res, nil
}
