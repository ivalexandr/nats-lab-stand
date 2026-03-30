package services

import (
	"reader/broker"
	"reader/common"
	"reader/database"

	"github.com/nats-io/nats.go"
)

type Handlers struct {
	Pg   *database.Pg
	nats *broker.Nats
	cfg  *common.Config
}

func NewHandlers() (*Handlers, error) {
	cfg, err := common.NewConfig()
	if err != nil {
		return nil, err
	}

	pg, err := database.NewPg(*cfg)
	if err != nil {
		return nil, err
	}

	nats, err := broker.NewNatsConnection(*cfg)
	if err != nil {
		return nil, err
	}

	return &Handlers{
		nats: nats,
		Pg:   pg,
		cfg:  cfg,
	}, nil
}

func (t *Handlers) Reader() (*nats.Subscription, error) {
	sub, err := t.Pg.SubscribeOrders(t.nats)
	if err != nil {
		return nil, err
	}

	return sub, nil
}
