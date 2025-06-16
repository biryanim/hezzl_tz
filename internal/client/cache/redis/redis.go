package redis

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/client/cache"
	"github.com/biryanim/hezzl_tz/internal/config"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var _ cache.RedisClient = (*client)(nil)

type handler func(ctx context.Context, conn redis.Conn) error

type client struct {
	pool   *redis.Pool
	config config.RedisConfig
}

func NewClient(pool *redis.Pool, config config.RedisConfig) *client {
	return &client{
		pool:   pool,
		config: config,
	}
}

func (c *client) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("SETEX", key, ttl.Seconds(), value)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) Get(ctx context.Context, key string) (interface{}, error) {
	var value interface{}
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		var errEx error
		value, errEx = conn.Do("GET", key)
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *client) DeleteByPattern(ctx context.Context, pattern string) error {
	return c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		var keys []string
		var cursor uint64
		for {
			reply, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", pattern, "COUNT", 100))
			if err != nil {
				return err
			}

			cursor, _ = redis.Uint64(reply[0], nil)
			chunk, _ := redis.Strings(reply[1], nil)
			keys = append(keys, chunk...)

			if cursor == 0 {
				break
			}
		}
		if len(keys) > 0 {
			_, err := conn.Do("DEL", redis.Args{}.AddFlat(keys)...)
			return err
		}
		return nil
	})
}

func (c *client) Ping(ctx context.Context) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("PING")
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) execute(ctx context.Context, handler handler) error {
	conn, err := c.getConnect(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Printf("failed to close redis connection: %v\n", err)
		}
	}()

	err = handler(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) getConnect(ctx context.Context) (redis.Conn, error) {
	getConnTimeoutCtx, cancel := context.WithTimeout(ctx, c.config.ConnectionTimeout())
	defer cancel()

	conn, err := c.pool.GetContext(getConnTimeoutCtx)
	if err != nil {
		log.Printf("failed to get redis connection: %v\n", err)

		_ = conn.Close()
		return nil, err
	}

	return conn, nil
}
