package domain

import "time"

type Todo struct {
	ID        int64
	Text      string
	Checked   bool
	CreatedAt time.Time
}
