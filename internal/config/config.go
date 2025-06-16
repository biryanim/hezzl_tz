package config

import (
	"github.com/joho/godotenv"
	"time"
)

type PGConfig interface {
	DSN() string
}

type HTTPConfig interface {
	Address() string
}

type NatsConfig interface {
	URL() string
	Subject() string
}

type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}

type ClickhouseConfig interface {
	DSN() string
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
