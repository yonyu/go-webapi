package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yonyu/go-webapi/internal/domain"
)

type CommentService interface {
	PostComment(ctx context.Context, comment domain.Comment) (domain.Comment, error)
	GetComment(ct context.Context, ID string) (domain.Comment, error)
	UpdateComment(ct context.Context, ID string, comment domain.Comment) (domain.Comment, error)
	DeleteComment(ct context.Context, ID string) error
}

type Response struct {
	Message string
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	// get/create domain.Comment
	//comment := r.Body.Read()
	var comment domain.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return
	}

	comment, err := h.Service.PostComment(r.Context(), comment)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	// get comment id
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	// get id
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// decode body to get new comment
	var comment domain.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return
	}

	// update the existing comment with the new comment
	comment, err := h.Service.UpdateComment(r.Context(), id, comment)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// encode the updated comment into response
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID := vars["id"]

	if commentID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteComment(r.Context(), commentID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}
}
