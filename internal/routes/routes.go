package routes

import (
	"byteknot-api/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func ByteKnotRouter(articleHandler *handler.ArticleHandler, commentsHandler *handler.CommentsHandler) *mux.Router {
	router := mux.NewRouter()
	articleSubrouter := router.PathPrefix("/api/v1").Subrouter()
	articleSubrouter.HandleFunc("/article/all", articleHandler.GetTop5ArticleOfDayHandler).Methods(http.MethodGet).Name("Fetch top 5 Articles")
	articleSubrouter.HandleFunc("/article/search", articleHandler.SearchArticleHandler).Methods(http.MethodGet).Queries("q", "{q}").Name("Search Articles")
	articleSubrouter.HandleFunc("/article", articleHandler.ArticleCRUDHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete).Name("Fetch article by ID")

	commentsSubroutine := router.PathPrefix("/api/v1").Subrouter()
	commentsSubroutine.HandleFunc("/comment/all", commentsHandler.GetComments).Methods(http.MethodGet).Name("Fetch comments by ArticleID")
	commentsSubroutine.HandleFunc("/comment", commentsHandler.WriteComment).Methods(http.MethodPost).Name("Write comment on Article")
	commentsSubroutine.HandleFunc("/comment", commentsHandler.UpdateComment).Methods(http.MethodPut).Name("Update comment by commentId")
	commentsSubroutine.HandleFunc("/comment", commentsHandler.DeleteComment).Methods(http.MethodDelete).Name("Delete comment by commentId")

	return router
}
