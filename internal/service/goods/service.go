package goods

import (
	"github.com/biryanim/hezzl_tz/internal/client/cache"
	"github.com/biryanim/hezzl_tz/internal/client/db"
	"github.com/biryanim/hezzl_tz/internal/repository"
	"github.com/biryanim/hezzl_tz/internal/service"
)

var _ service.GoodsService = (*serv)(nil)

type serv struct {
	goodsRepository repository.GoodsRepository
	cache           cache.RedisClient
	txManager       db.TxManager
}

func NewService(goodsRepository repository.GoodsRepository, cache cache.RedisClient, txManager db.TxManager) *serv {
	return &serv{
		goodsRepository: goodsRepository,
		cache:           cache,
		txManager:       txManager,
	}
}
