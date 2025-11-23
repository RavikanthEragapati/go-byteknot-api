package handler

import (
	"byteknot-api/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	ArticleService service.ArticleService
}

func NewArticleHandler(articleService service.ArticleService) *ArticleHandler {
	return &ArticleHandler{ArticleService: articleService}
}

func (h ArticleHandler) GetTop5ArticleOfDayHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Top 5 Articles</h1>")
}

func (h ArticleHandler) SearchArticleHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Top 5 Articles</h1>")
}

func (h ArticleHandler) ArticleCRUDHandler(writer http.ResponseWriter, request *http.Request) {
	method := request.Method

	vars := mux.Vars(request)
	idStr, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, "Missing 'data' query parameter.", http.StatusBadRequest)
		return
	}

	switch method {
	case "GET":
		h.ArticleService.GetArticleByIDHandler(idStr)
	case "POST":
		h.ArticleService.CreateArticleHandler(idStr, request.Body)
	case "PUT":
		h.ArticleService.UpdateArticleHandler(idStr, request.Body)
	case "DELETE":
		h.ArticleService.DeleteArticleHandler(idStr)
	default:
		http.Error(writer, "Missing 'data' query parameter.", http.StatusBadRequest)
		return
	}
}
