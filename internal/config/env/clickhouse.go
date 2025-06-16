package env

import (
	"github.com/biryanim/hezzl_tz/internal/config"
	"github.com/pkg/errors"
	"os"
)

const (
	clickhouseDSNEnvName = "CH_DSN"
)

type clickhouseConfig struct {
	dsn string
}

func NewClickhouseConfig() (config.ClickhouseConfig, error) {
	dsn := os.Getenv(clickhouseDSNEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("clickhouse DSN is empty")
	}

	return &clickhouseConfig{
		dsn: dsn,
	}, nil
}

func (c *clickhouseConfig) DSN() string {
	return c.dsn
}
