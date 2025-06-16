package app

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/client/cache/redis"
	goodsRepository "github.com/biryanim/hezzl_tz/internal/repository/goods"
	goodsService "github.com/biryanim/hezzl_tz/internal/service/goods"
	"github.com/biryanim/hezzl_tz/internal/api/goods"
	"github.com/biryanim/hezzl_tz/internal/client/cache"
	"github.com/biryanim/hezzl_tz/internal/client/db"
	"github.com/biryanim/hezzl_tz/internal/client/db/pg"
	"github.com/biryanim/hezzl_tz/internal/client/db/transaction"
	"github.com/biryanim/hezzl_tz/internal/config"
	"github.com/biryanim/hezzl_tz/internal/config/env"
	"github.com/biryanim/hezzl_tz/internal/repository"
	"github.com/biryanim/hezzl_tz/internal/service"
	redigo "github.com/gomodule/redigo/redis"
	"log"
)

type serviceProvider struct {
	pgConfig    config.PGConfig
	httpConfig  config.HTTPConfig
	redisConfig config.RedisConfig

	redisPool *redigo.Pool

	redisClient cache.RedisClient
	dbClient    db.Client
	txManager   db.TxManager

	goodsRepository repository.GoodsRepository

	goodsService service.GoodsService

	goodsImpl *goods.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to load pg config: %v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to load http config: %v", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := env.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to load redis config: %v", err)
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) RedisPool() *redigo.Pool {
	if s.redisPool == nil {
		s.redisPool = &redigo.Pool{
			MaxIdle:     s.RedisConfig().MaxIdle(),
			IdleTimeout: s.RedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", s.RedisConfig().Address())
			},
		}
	}

	return s.redisPool
}

func (s *serviceProvider) RedisClient() cache.RedisClient {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(s.RedisPool(), s.RedisConfig())
	}

	return s.redisClient
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %v", err)
		}

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) GoodsRepository(ctx context.Context) repository.GoodsRepository {
	if s.goodsRepository == nil {
		s.goodsRepository = goodsRepository.NewRepository(s.DBClient(ctx))
	}

	return s.goodsRepository
}

func (s *serviceProvider) GoodsService(ctx context.Context) service.GoodsService {
	if s.goodsService == nil {
		s.goodsService = goodsService.NewService(s.GoodsRepository(ctx), s.RedisClient(), s.TxManager(ctx))

	}
	return s.goodsService
}

func (s *serviceProvider) GoodsImpl(ctx context.Context) *goods.Implementation {
	if s.goodsImpl == nil {
		s.goodsImpl = goods.NewImplementation(s.GoodsService(ctx))
	}

	return s.goodsImpl
}
