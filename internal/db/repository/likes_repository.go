package repository

import "time"

type Likes struct {
	LikeId    int64
	ArticleId int64
	UserId    int64
	CreatedAt time.Time
}
