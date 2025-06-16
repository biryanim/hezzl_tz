package app

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/biryanim/hezzl_tz/internal/api/goods"
	"github.com/biryanim/hezzl_tz/internal/client/broker"
	natsCl "github.com/biryanim/hezzl_tz/internal/client/broker/nats"
	"github.com/biryanim/hezzl_tz/internal/client/cache"
	"github.com/biryanim/hezzl_tz/internal/client/cache/redis"
	"github.com/biryanim/hezzl_tz/internal/client/db"
	"github.com/biryanim/hezzl_tz/internal/client/db/pg"
	"github.com/biryanim/hezzl_tz/internal/client/db/transaction"
	"github.com/biryanim/hezzl_tz/internal/config"
	"github.com/biryanim/hezzl_tz/internal/config/env"
	"github.com/biryanim/hezzl_tz/internal/repository"
	goodsRepository "github.com/biryanim/hezzl_tz/internal/repository/goods"
	logRepository "github.com/biryanim/hezzl_tz/internal/repository/log"
	"github.com/biryanim/hezzl_tz/internal/service"
	goodsService "github.com/biryanim/hezzl_tz/internal/service/goods"
	loggerService "github.com/biryanim/hezzl_tz/internal/service/logger"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type serviceProvider struct {
	pgConfig         config.PGConfig
	httpConfig       config.HTTPConfig
	redisConfig      config.RedisConfig
	clickhouseConfig config.ClickhouseConfig
	natsConfig       config.NatsConfig

	redisPool *redigo.Pool
	nats      *nats.Conn
	chConn    clickhouse.Conn

	redisClient  cache.RedisClient
	dbClient     db.Client
	txManager    db.TxManager
	brokerClient broker.Publisher

	goodsRepository repository.GoodsRepository
	logRepository   repository.LogRepository

	goodsService  service.GoodsService
	loggerService service.LoggerService

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

func (s *serviceProvider) CHConfig() config.ClickhouseConfig {
	if s.clickhouseConfig == nil {
		cfg, err := env.NewClickhouseConfig()
		if err != nil {
			log.Fatalf("failed to load clickhouse config: %v", err)
		}
		s.clickhouseConfig = cfg
	}
	return s.clickhouseConfig
}

func (s *serviceProvider) NatsConfig() config.NatsConfig {
	if s.natsConfig == nil {
		cfg, err := env.NewNatsConfig()
		if err != nil {
			log.Fatalf("failed to load nats config: %v", err)
		}
		s.natsConfig = cfg
	}

	return s.natsConfig
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

func (s *serviceProvider) CHConn() clickhouse.Conn {
	if s.chConn == nil {
		opt, err := clickhouse.ParseDSN(s.CHConfig().DSN())
		if err != nil {
			log.Fatalf("failed to parse clickhouse DSN: %v", err)
		}

		conn, err := clickhouse.Open(opt)
		if err != nil {
			log.Fatalf("failed to open clickhouse connection: %v", err)
		}
		s.chConn = conn
	}
	return s.chConn
}

func (s *serviceProvider) NatsConn() *nats.Conn {
	if s.nats == nil {
		fmt.Println(s.NatsConfig().URL())
		nc, err := nats.Connect(s.NatsConfig().URL(), nats.Timeout(5*time.Second), nats.PingInterval(20*time.Second), nats.MaxReconnects(3))
		if err != nil {
			log.Fatalf("failed to connect to nats server: %v", err)
		}

		s.nats = nc
	}

	return s.nats
}

func (s *serviceProvider) PublisherClient() broker.Publisher {
	if s.brokerClient == nil {
		s.brokerClient = natsCl.NewClient(s.NatsConn())
	}
	return s.brokerClient
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

func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewRepository(s.CHConn())
	}

	return s.logRepository
}

func (s *serviceProvider) GoodsService(ctx context.Context) service.GoodsService {
	if s.goodsService == nil {
		s.goodsService = goodsService.NewService(s.GoodsRepository(ctx), s.RedisClient(), s.TxManager(ctx), s.PublisherClient())
	}

	return s.goodsService
}

func (s *serviceProvider) LoggerService(ctx context.Context) service.LoggerService {
	if s.loggerService == nil {
		s.loggerService = loggerService.NewService(s.NatsConn(), s.LogRepository(ctx), s.NatsConfig().Subject())
	}

	return s.loggerService
}

func (s *serviceProvider) GoodsImpl(ctx context.Context) *goods.Implementation {
	if s.goodsImpl == nil {
		s.goodsImpl = goods.NewImplementation(s.GoodsService(ctx))
	}

	return s.goodsImpl
}
