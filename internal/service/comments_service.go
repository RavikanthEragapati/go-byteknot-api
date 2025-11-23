package service

import "byteknot-api/internal/db/repository"

type CommentsService interface {
	GetComments(id int) (repository.Comments, error)
	CreateComments(comments repository.Comments) (repository.Comments, error)
	UpdateComments(id int, comments repository.Comments) (repository.Comments, error)
	DeleteComments(id int) error
}

// Create Instance of CommentsService by injecting CommentsRepo instance
type commentsServiceImpl struct {
	commentsRepo repository.CommentsRepository
}

func NewCommentsService(commentsRepo repository.CommentsRepository) CommentsService {
	return commentsServiceImpl{commentsRepo: commentsRepo}
}

// Implementations Of CommentsService Interface

func (c commentsServiceImpl) GetComments(id int) (repository.Comments, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsServiceImpl) CreateComments(comments repository.Comments) (repository.Comments, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsServiceImpl) UpdateComments(id int, comments repository.Comments) (repository.Comments, error) {
	//TODO implement me
	panic("implement me")
}

func (c commentsServiceImpl) DeleteComments(id int) error {
	//TODO implement me
	panic("implement me")
}
