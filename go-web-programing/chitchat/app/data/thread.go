package data

import (
	"time"
)

type (
	Thread struct {
		ID int
		Uuid string
		Topic string
		UserID int
		CreatedAt time.Time
	}
)
