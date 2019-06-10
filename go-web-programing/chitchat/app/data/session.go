package data

import (
	"time"
)

type (
	Session struct {
		Id       int
		Uuid     string
		Email    string
		UserID   int
		CreateAt time.Time
	}
)

func (sess Session) Check() (ok bool, err error) {
	return
}
