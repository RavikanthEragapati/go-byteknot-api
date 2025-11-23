package repository

import "time"

type Users struct {
	UserId      int64
	FireBaseUid string
	UserName    string
	Email       string
	CreatedAt   time.Time
}
