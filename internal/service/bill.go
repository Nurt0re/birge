package service

import (
	"birge/internal/model"
	"birge/internal/repository"
	"context"
	"log/slog"
)

type BillServiceImpl struct {
	repo repository.BillRepo
	log  *slog.Logger
}

func NewBillService(repo repository.BillRepo, log *slog.Logger) *BillServiceImpl {
	return &BillServiceImpl{
		repo: repo,
		log:  log,
	}
}

func (s *BillServiceImpl) NewBill(ctx context.Context, chatID int64, creatorUserID int64, username string) (int64, error) {
	billID, err := s.repo.CreateBill(ctx, chatID)
	if err != nil {
		return 0, err
	}
	
	// Automatically add creator as first participant
	err = s.repo.AddUserToBill(ctx, billID, creatorUserID, username)
	if err != nil {
		s.log.Warn("Failed to add creator to bill", "error", err, "bill_id", billID)
	}
	
	return billID, nil
}

func (s *BillServiceImpl) AddUserToBill(ctx context.Context, billID int64, userID int64, username string) error {
	return s.repo.AddUserToBill(ctx, billID, userID, username)
}
func (s *BillServiceImpl) GetBillParticipants(ctx context.Context, billID int64) ([]model.User, error) {
	return s.repo.GetBillParticipants(ctx, billID)
}
