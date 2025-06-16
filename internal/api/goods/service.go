package goods

import "github.com/biryanim/hezzl_tz/internal/service"

type Implementation struct {
	goodsService service.GoodsService
}

func NewImplementation(goodsService service.GoodsService) *Implementation {
	return &Implementation{
		goodsService: goodsService,
	}
}
