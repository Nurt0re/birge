package model

type User struct {
	ID         int64
	Username   string
	AmountOwed float64
	HasPaid    bool
}
