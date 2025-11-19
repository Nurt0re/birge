package repository

import (
	"birge/internal/model"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type BillRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func NewBillRepository(db *sql.DB, log *slog.Logger) *BillRepository {
	return &BillRepository{
		db:  db,
		log: log,
	}
}

func (r *BillRepository) CreateBill(ctx context.Context, chatID int64) (int64, error) {

	query := `
	INSERT INTO bills (chat_id, created_at)
	VALUES ($1, NOW())
	RETURNING id;
	`
	var billID int64
	err := r.db.QueryRowContext(ctx, query, chatID).Scan(&billID)
	if err != nil {
		return 0, fmt.Errorf("failed to create a bill: %w", err)
	}
	return billID, nil
}

func (r *BillRepository) AddUserToBill(ctx context.Context, billID int64, userID int64, username string) error {
	query := `
	INSERT INTO participants (bill_id, user_id, username)
	VALUES ($1, $2, $3);
	`
	_, err := r.db.ExecContext(ctx, query, billID, userID, username)
	if err != nil {
		return fmt.Errorf("error at adding participants to the bill: %w", err)
	}
	return nil

}

func (r *BillRepository) GetBillParticipants(ctx context.Context, billID int64) ([]model.User, error) {
	query := `
	SELECT user_id, username
	FROM participants
	WHERE bill_id = $1
	ORDER BY joined_at;
	`
	rows, err := r.db.QueryContext(ctx, query, billID)
	if err != nil {
		return nil, fmt.Errorf("failed to get bill participants: %w", err)
	}
	defer rows.Close()

	var participants []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, fmt.Errorf("failed to scan participant: %w", err)
		}
		participants = append(participants, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return participants, nil
}
