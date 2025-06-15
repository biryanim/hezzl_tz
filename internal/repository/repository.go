package repository

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/model"
)

type GoodsRepository interface {
	Create(ctx context.Context, good *model.GoodCreateParams) (*model.Good, error)
	GetByIds(ctx context.Context, id, projectId int) (*model.Good, error)
	Update(ctx context.Context, good *model.GoodUpdateParams) (*model.Good, error)
	RemoveGood(ctx context.Context, id, projectId int) error
	Delete(ctx context.Context, good *model.GoodDRemoveParams) error
	List(ctx context.Context, good *model.GoodListParams) (*model.GoodsList, error)
	Reprioritize(ctx context.Context, good *model.GoodReprioritizeParams) (*model.GoodsPrioritize, error)
}
