package service

import (
	"byteknot-api/internal/db/repository"
	"io"
)

type ArticleService interface {
	GetTop5ArticleListHandler() ([]repository.Article, error)
	GetArticleSearchHandler() ([]repository.Article, error)
	GetArticleByIDHandler(id int) (repository.Article, error)
	CreateArticleHandler(id int, r io.ReadCloser) (repository.Article, error)
	UpdateArticleHandler(id int, r io.ReadCloser) (repository.Article, error)
	DeleteArticleHandler(id int) error
}

// Create Instance of ArticleService by injecting ArticleRepo instance
type articleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) ArticleService {
	return &articleServiceImpl{ArticleRepository: articleRepo}
}

// Implementations Of ArticleService Interface

func (a articleServiceImpl) GetTop5ArticleListHandler() ([]repository.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleServiceImpl) GetArticleSearchHandler() ([]repository.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleServiceImpl) GetArticleByIDHandler(id int) (repository.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleServiceImpl) CreateArticleHandler(id int, r io.ReadCloser) (repository.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleServiceImpl) UpdateArticleHandler(id int, r io.ReadCloser) (repository.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleServiceImpl) DeleteArticleHandler(id int) error {
	//TODO implement me
	panic("implement me")
}
