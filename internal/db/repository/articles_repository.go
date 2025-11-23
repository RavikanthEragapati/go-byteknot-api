package repository

import (
	"database/sql"
	"time"
)

type Article struct {
	ArticleId int64
	Title     string
	Content   string
	AuthorId  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ArticleRepository interface {
	FindAll() ([]Article, error)
	FindByID(id int) (Article, error)
}

// Create Instance of type commentsRepository by injecting *sql.DB
type articleRepositoryImpl struct {
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepositoryImpl{DB: db}
}

// Implementation of Repository functions

func (a articleRepositoryImpl) FindAll() ([]Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleRepositoryImpl) FindByID(id int) (Article, error) {
	//TODO implement me
	panic("implement me")
}
