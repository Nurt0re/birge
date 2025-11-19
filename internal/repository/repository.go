package repository

import (
	"birge/internal/model"
	"context"
	"database/sql"
	"log/slog"
)

type BillRepo interface {
	CreateBill(ctx context.Context, chatID int64) (int64, error)
	AddUserToBill(ctx context.Context, billID int64, userID int64, username string) error
	GetBillParticipants(ctx context.Context, billID int64) ([]model.User, error)
}

type Repository struct {
	BillRepo BillRepo
}

func NewRepository(db *sql.DB, log *slog.Logger) *Repository {
	return &Repository{
		BillRepo: NewBillRepository(db, log),
	}
}
