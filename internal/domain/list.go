package domain

import "time"

type List struct {
	ID        int64
	Title     string
	User_ID   int64
	Archived  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
