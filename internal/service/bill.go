package service

import "birge/internal/model"

func NewBill(chatID int64, userID int64) (int64, error) {
	// Logic to create a new bill and return its ID
	return 0, nil
}

func AddUserToBill(billID int64, userID int64, username string) error {
	// Logic to add a user to the bill
	return nil
}
func GetBillParticipants(billID int64) ([]model.User, error) {
	// Logic to get participants of the bill
	return []model.User{}, nil
}
