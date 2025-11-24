package handler

import (
	"byteknot-api/internal/db/repository"
	"byteknot-api/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

// ArticleHandler Create Instance by injecting ArticleRepo instance
type ArticleHandler struct {
	ArticleService service.ArticleService
}

func NewArticleHandler(articleService service.ArticleService) *ArticleHandler {
	return &ArticleHandler{ArticleService: articleService}
}

// Implementations Of ArticleHandler Interface

func (h ArticleHandler) GetTop5ArticleOfDayHandler(writer http.ResponseWriter, request *http.Request) {
	h.ArticleService.GetTop5ArticleList()
}

func (h ArticleHandler) SearchArticleHandler(writer http.ResponseWriter, request *http.Request) {
	h.ArticleService.SearchArticle()
}

func (h ArticleHandler) ArticleCRUDHandler(writer http.ResponseWriter, request *http.Request) {
	method := request.Method

	queryParams := request.URL.Query()
	id, err := strconv.Atoi(queryParams.Get("id"))
	if err != nil {
		http.Error(writer, "Missing 'data' query parameter.", http.StatusBadRequest)
		return
	}

	switch method {
	case "GET":
		article, err := h.ArticleService.GetArticleByID(id)
		if err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}
		articleJSON, err := json.Marshal(article)
		if err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(articleJSON)
		return

	case "POST":
		var article repository.Article
		defer request.Body.Close()
		if err := json.NewDecoder(request.Body).Decode(&article); err != nil {
			http.Error(writer, "Invalid request body", http.StatusBadRequest)
		}
		result, err := h.ArticleService.CreateArticle(article)
		if err != nil {
			http.Error(writer, "Failed to create article.", http.StatusInternalServerError)
			return
		}
		responseBody, err := json.Marshal(result)
		if err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		writer.Write(responseBody)
		return

	case "PUT":
		var article repository.Article
		defer request.Body.Close()
		if err := json.NewDecoder(request.Body).Decode(&article); err != nil {
			http.Error(writer, "Invalid request body", http.StatusBadRequest)
		}

		_, err := h.ArticleService.UpdateArticle(id, article)
		if err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}

	case "DELETE":
		err := h.ArticleService.DeleteArticle(id)
		if err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusNoContent)
		return

	default:
		http.Error(writer, "Missing 'data' query parameter.", http.StatusBadRequest)
		return
	}
}
