package core

import (
	"context"
	"time"
)

type UserID int

type Transfer struct {
	ID          int
	From        UserID
	To          UserID
	Amount      int
	Type        string
	PerformedAt time.Time
}

type TransferRepository interface {
	Add(ctx context.Context, tfs *Transfer) error
}
