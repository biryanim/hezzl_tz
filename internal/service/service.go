package service

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/model"
)

type GoodsService interface {
	Create(ctx context.Context, goodsCreatingParams *model.GoodCreateParams) (*model.Good, error)
	Update(ctx context.Context, goodsUpdatingParams *model.GoodUpdateParams) (*model.Good, error)
	Remove(ctx context.Context, goodsRemovingParams *model.GoodDRemoveParams) (*model.GoodRemove, error)
	List(ctx context.Context, goodsListingParams *model.GoodListParams) (*model.GoodsList, error)
	Reprioritize(ctx context.Context, goodReprioritizingParams *model.GoodReprioritizeParams) (*model.GoodsPrioritize, error)
}
