package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"publisher/models"

	"publisher/common"

	"publisher/broker"

	"github.com/jackc/pgx/v5/pgxpool"
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

func (p *Pg) PublishData(nt *broker.Nats) error {
	rows, err := p.instance.Query(p.ctx, `
		SELECT id, external_id, status, amount, created_at
		FROM orders
		WHERE sent_at IS NULL
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.ExternalID, &order.Status, &order.Amount, &order.CreatedAt)
		if err != nil {
			log.Println("Row scan error", err)
			continue
		}

		data, err := json.Marshal(order)
		if err != nil {
			log.Println("Marshall error", err)
			continue
		}

		if err := nt.Instance.Publish(models.Subject, data); err != nil {
			log.Println("publish error", err)
			continue
		}

		if err := p.setSendAt(order.ID); err != nil {
			log.Println("update error", err)
			continue
		}

		log.Print("Send order: ", order.ExternalID)
	}

	log.Println("Send done")
	return nil
}

func (p *Pg) setSendAt(id int64) error {
	_, err := p.instance.Exec(p.ctx, `UPDATE orders SET sent_at = NOW() WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("update error", err)
	}
	return nil
}

func (p *Pg) CloseConnect() {
	p.instance.Close()
}
