package repository

import (
	"time"
)

type ArticleViews struct {
	ArticleViewId int64
	ArticleId     int64
	UserId        int64
	SessionId     string
	ViewedAt      time.Time
}

//
//type articleViewsRepository struct {
//	db *sql.DB
//}
//
//func articleViewsRepository(db *sql.DB) ArticleViewsRepository {
//	return &articleViewsRepository{db: db}
//}
//
//type ArticleViewsRepository interface {
//	FindByID(ctx context.Context, id int) (*ArticleViews, error)
//	CreateUser(ctx context.Context, name, email string) (*ArticleViews, error)
//}
