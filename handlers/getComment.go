package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jamesthomasw/go-mygram-v1/models"
)

type ResponseGetComment struct {
	ID         int       `json:"id"`
	Message    string    `json:"message"`
	PhotoRefer int       `json:"photo_id"`
	UserRefer  int       `json:"user_id"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func CreateGetCommentResponse(comment models.Comment) ResponseGetComment {
	return ResponseGetComment{ID: comment.ID, Message: comment.Message, PhotoRefer: comment.PhotoRefer, UserRefer: comment.UserRefer, CreatedAt: *comment.CreatedAt, UpdatedAt: *comment.UpdatedAt}
}

func (h Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["commentId"])

	var comment models.Comment

	if result := h.DB.Find(&comment, id); result.Error != nil {
		panic(result.Error)
	}

	responseGetComment := CreateGetCommentResponse(comment)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseGetComment)
}