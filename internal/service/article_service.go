package service

import (
	"byteknot-api/internal/db/repository"
)

type ArticleService interface {
	GetTop5ArticleList() ([]repository.Article, error)
	SearchArticle() ([]repository.Article, error)
	GetArticleByID(id int) (repository.Article, error)
	CreateArticle(article repository.Article) (repository.Article, error)
	UpdateArticle(id int, article repository.Article) (repository.Article, error)
	DeleteArticle(id int) error
}

// Create Instance of ArticleService by injecting ArticleRepo instance
type articleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) ArticleService {
	return &articleServiceImpl{ArticleRepository: articleRepo}
}

// Implementations Of ArticleService Interface

func (a articleServiceImpl) GetTop5ArticleList() ([]repository.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleServiceImpl) SearchArticle() ([]repository.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleServiceImpl) GetArticleByID(id int) (repository.Article, error) {
	article, err := a.ArticleRepository.FindByID(id)
	if err != nil {
		return repository.Article{}, err
	}
	return article, nil
}

func (a articleServiceImpl) CreateArticle(article repository.Article) (repository.Article, error) {

	articleId, err := a.ArticleRepository.InsertArticle(article)
	if err != nil {
		return repository.Article{}, err
	}
	return repository.Article{ArticleId: articleId}, nil
}

func (a articleServiceImpl) UpdateArticle(id int, article repository.Article) (repository.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleServiceImpl) DeleteArticle(id int) error {
	//TODO implement me
	panic("implement me")
}
