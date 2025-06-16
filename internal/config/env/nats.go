package env

import (
	"github.com/biryanim/hezzl_tz/internal/config"
	"github.com/pkg/errors"
	"os"
)

const (
	natsURLEnvName          = "NATS_URL"
	natsGoodsSubjectEnvName = "NATS_GOODS_SUBJECT"
)

type natsConfig struct {
	url          string
	goodsSubject string
}

func NewNatsConfig() (config.NatsConfig, error) {
	url := os.Getenv(natsURLEnvName)
	if len(url) == 0 {
		return nil, errors.New("nats url not found")
	}

	subject := os.Getenv(natsGoodsSubjectEnvName)
	if len(subject) == 0 {
		return nil, errors.New("nats subject not found")
	}

	return &natsConfig{
		url:          url,
		goodsSubject: subject,
	}, nil
}

func (c *natsConfig) URL() string {
	return c.url
}

func (c *natsConfig) Subject() string {
	return c.goodsSubject
}
