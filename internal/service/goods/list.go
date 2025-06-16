package goods

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/biryanim/hezzl_tz/internal/model"
	"time"
)

func (s *serv) List(ctx context.Context, goodsListingParams *model.GoodListParams) (*model.GoodsList, error) {
	var goods *model.GoodsList
	cacheKey := fmt.Sprintf("goods:list:limit=%d&offset=%d", goodsListingParams.Limit, goodsListingParams.Offset)
	values, err := s.cache.Get(ctx, cacheKey)
	if err == nil && values != nil {
		switch v := values.(type) {
		case []byte:
			err = json.Unmarshal(v, goods)
			if err == nil {
				return goods, nil
			}
		case string:
			err = json.Unmarshal([]byte(v), goods)
			if err == nil {
				return goods, nil
			}
		}
	}

	err = s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		goods, err = s.goodsRepository.List(ctx, goodsListingParams)
		return err
	})
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(goods)
	if err != nil {
		return nil, err
	}

	err = s.cache.Set(ctx, cacheKey, data, 1*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("failed to set goods list cache: %w", err)
	}

	return goods, nil
}
