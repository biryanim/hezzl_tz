package env

import (
	"errors"
	"github.com/biryanim/hezzl_tz/internal/config"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

func NewPGConfig() (config.PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg DSN not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (p *pgConfig) DSN() string {
	return p.dsn
}
