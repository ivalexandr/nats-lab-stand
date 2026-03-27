package services

import (
	"publisher/broker"
	"publisher/common"
	"publisher/database"
)

type Handlers struct {
	nats *broker.Nats
	pg   *database.Pg
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
		pg:   pg,
		cfg:  cfg,
	}, nil
}

func (t *Handlers) Publisher() error {
	if err := t.pg.PublishData(t.nats); err != nil {
		return err
	}

	defer t.pg.CloseConnect()
	return nil
}
