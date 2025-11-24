package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Article struct {
	ArticleId int64 `db:"article_id"`
	Title     string
	Content   string
	AuthorId  int64     `db:"author_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const insertSQL = `insert into articles (title, content, author_id) values (?, ?, ?)`
const selectSQL = `select article_id, title, content, author_id, created_at, updated_at from articles where article_id = ?`

type ArticleRepository interface {
	FindAll() ([]Article, error)
	FindByID(id int) (Article, error)
	InsertArticle(article Article) (int64, error)
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

// FindByID function implementation that takes in an article_id and returns the article record for the id
func (a articleRepositoryImpl) FindByID(id int) (Article, error) {
	log.Println("REPOSITORY: Executing SQL query to find article by ID...")

	// 1. Use QueryRow for single results
	// This is safer and simpler than Query() + Next() for finding a single row.
	row := a.DB.QueryRow(selectSQL, id)

	var article Article

	// 2. Scan directly from the row
	err := row.Scan(
		&article.ArticleId,
		&article.Title,
		&article.Content,
		&article.AuthorId,
		&article.CreatedAt,
		&article.UpdatedAt,
	)

	if err != nil {
		// 3. Handle specific 'No Rows Found' error
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("REPOSITORY: Article with ID %d not found.\n", id)
			// You might return a custom error type here for better handling in services
			return Article{}, err
		}
		log.Println("REPOSITORY: Scan error:", err)
		return Article{}, err
	}

	log.Printf("REPOSITORY: Successfully found article %d.\n", article.ArticleId)
	return article, nil
}

// InsertArticle function implementation that insert an Article to DB and returns the article_id
func (a articleRepositoryImpl) InsertArticle(article Article) (int64, error) {
	log.Println("REPOSITORY: Executing SQL query to insert article...")

	tx, err := a.DB.Begin()
	if err != nil {
		log.Println("REPOSITORY: Begin transaction error:", err)
		return 0, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(insertSQL, article.Title, article.Content, article.AuthorId)
	if err != nil {
		return 0, fmt.Errorf("error executing insert: %w", err)
	}
	var newID int64
	newID, err = result.LastInsertId()
	if err != nil {
		log.Printf("error retrieving last insert ID: %w", err)
		return 0, fmt.Errorf("error retrieving last insert ID: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("could not commit transaction: %w", err)
	}

	return newID, nil
}
