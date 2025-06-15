package goods

import (
	"context"
	"fmt"
	"github.com/biryanim/hezzl_tz/internal/model"
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

	keys := make([]string, len(goods.Priorities))
	for _, good := range goods.Priorities {
		keys = append(keys, fmt.Sprintf("goods:%d:%d", good.ID, goodReprioritizingParams.ProjectID))
	}

	err = s.cache.Delete(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("delete goods prioritizing cache: %w", err)
	}

	return goods, nil
}
