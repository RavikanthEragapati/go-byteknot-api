package repository

import (
	"database/sql"
	"time"
)

type Comments struct {
	CommentId       int64
	ArticleId       int64
	UserId          int64
	Content         string
	ParentCommentId int64
	CreatedAt       time.Time
}

type CommentsRepository interface {
	FindAll() ([]Comments, error)
	FindByID(id int) (Comments, error)
}

// Create Instance of type commentsRepository by injecting *sql.DB
type commentsRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentsRepository(db *sql.DB) CommentsRepository {
	return &commentsRepositoryImpl{DB: db}
}

// Implementation of Repository functions

func (c commentsRepositoryImpl) FindAll() ([]Comments, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsRepositoryImpl) FindByID(id int) (Comments, error) {
	//TODO implement me
	panic("implement me")
}
