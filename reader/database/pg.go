package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"reader/broker"
	"reader/common"

	models "github.com/ivalexander/nats-lab/models/orders"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

type Pg struct {
	instance *pgxpool.Pool
	ctx      context.Context
}

func NewPg(config common.Config) (*Pg, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	ctx := context.Background()

	pgx, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &Pg{
		instance: pgx,
		ctx:      ctx,
	}, nil
}

func (p *Pg) SubscribeOrders(nt *broker.Nats) (*nats.Subscription, error) {
	cb := func(m *nats.Msg) {
		var order models.Order
		msgData := m.Data

		if err := json.Unmarshal(msgData, &order); err != nil {
			log.Printf("order saved: %s", order.ExternalID)
			return
		}

		_, err := p.instance.Exec(p.ctx, `
      INSERT INTO orders_mirror (source_id,external_id,status,amount,created_at,received_at)
      VALUES ($1,$2,$3,$4,$5,NOW())
      ON CONFLICT (external_id) DO NOTHING
    `, order.ID, order.ExternalID, order.Status, order.Amount, order.CreatedAt)
		if err != nil {
			log.Printf("failed to insert order into db: %v", err)
			return
		}

		log.Printf("Data read and insert successfuly!")
	}

	sub, err := nt.Instance.Subscribe(models.Subject, cb)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (p *Pg) CloseConnect() {
	p.instance.Close()
}
