package model

import "time"

type Bill struct {
	ID          int64
	TotalAmount float64
	CreatedAt   time.Time
}
