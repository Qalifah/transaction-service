package postgres

import (
	"context"
	core "github.com/Qalifah/grey-challenge/transaction"
	"github.com/jackc/pgx/v4"
)

type transferRepository struct {
	client *pgx.Conn
}

func NewTransferRepository(client *pgx.Conn) *transferRepository {
	return &transferRepository{
		client: client,
	}
}

func (t *transferRepository) Add(ctx context.Context, tfs *core.Transfer) error {
	_, err := t.client.Exec(ctx, `INSERT INTO transfers ("from", "to", "amount", "type", "performed_at") VALUES ($1, $2, $3, $4, $5)`, tfs.From, tfs.To, tfs.Amount, tfs.Type, tfs.PerformedAt)
	return err
}
