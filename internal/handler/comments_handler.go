package handler

import (
	"byteknot-api/internal/service"
	"net/http"
)

type CommentsHandler struct {
	CommentsService service.CommentsService
}

func NewCommentsHandler(commentsService service.CommentsService) *CommentsHandler {
	return &CommentsHandler{CommentsService: commentsService}
}

func (h *CommentsHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *CommentsHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *CommentsHandler) WriteComment(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *CommentsHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
