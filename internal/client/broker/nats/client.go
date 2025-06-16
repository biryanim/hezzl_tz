package nats

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/client/broker"
	"github.com/nats-io/nats.go"
)

var _ broker.Publisher = (*client)(nil)

type client struct {
	nc *nats.Conn
}

func NewClient(nc *nats.Conn) *client {
	return &client{nc: nc}
}

func (c *client) Publish(ctx context.Context, subject string, data []byte) error {
	return c.nc.Publish(subject, data)
}

func (c *client) Close() {
	if c.nc != nil {
		c.nc.Close()
	}
}
