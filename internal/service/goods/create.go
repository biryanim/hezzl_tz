package goods

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/model"
	"log"
)

func (s *serv) Create(ctx context.Context, goodsCreatingParams *model.GoodCreateParams) (*model.Good, error) {
	res, err := s.goodsRepository.Create(ctx, goodsCreatingParams)
	if err != nil {
		return nil, err
	}
	if err = s.cache.DeleteByPattern(ctx, "goods:list:*"); err != nil {
		log.Printf("delete goods cache err: %v", err)
	}

	s.publishLogEvent(ctx, res)

	return res, nil
}
