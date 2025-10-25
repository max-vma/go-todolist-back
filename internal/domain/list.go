package domain

import "time"

type List struct {
	ID        int64
	title     string
	Archived  bool
	CreatedAt time.Time
}
