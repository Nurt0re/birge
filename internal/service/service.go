package service

import (
	"birge/internal/model"
	"birge/internal/repository"
	"context"
	"log/slog"
)

type BillService interface {
	NewBill(ctx context.Context, chatID int64, creatorID int64, username string) (int64, error)
	AddUserToBill(ctx context.Context, billID int64, userID int64, username string) error
	GetBillParticipants(ctx context.Context, billID int64) ([]model.User, error)
}

type Service struct {
	BillService BillService
}

func NewService(r *repository.Repository, log *slog.Logger) *Service {
	return &Service{
		BillService: NewBillService(r.BillRepo, log),
	}
}
