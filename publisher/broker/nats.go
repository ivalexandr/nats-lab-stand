package broker

import (
	"publisher/common"

	"github.com/nats-io/nats.go"
)

type Nats struct {
	Instance *nats.Conn
}

func NewNatsConnection(cfg common.Config) (*Nats, error) {
	conn, err := nats.Connect(cfg.NATS_URL)
	if err != nil {
		return nil, err
	}

	return &Nats{
		Instance: conn,
	}, nil
}
