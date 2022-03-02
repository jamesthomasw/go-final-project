package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jamesthomasw/go-mygram-v1/helpers"
	"github.com/jamesthomasw/go-mygram-v1/models"
)

type UpdateCommentResponse struct {
	ID         int       `json:"id"`
	Message    string    `json:"message"`
	PhotoRefer int       `json:"photo_id"`
	UserRefer  int       `json:"user_id"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func CreateUpdateCommentResponse(comment models.Comment) UpdateCommentResponse {
	return UpdateCommentResponse{ID: comment.ID, Message: comment.Message, PhotoRefer: comment.PhotoRefer, UserRefer: comment.UserRefer, UpdatedAt: *comment.UpdatedAt}
}

func (h Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	
	type UpdateComment struct {
		Message string `json:"message"`
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["commentId"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var updatedComment UpdateComment
	json.Unmarshal(body, &updatedComment)

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	var user models.User

	userEmail := claims.(jwt.MapClaims)["email"].(string)
	h.DB.First(&user, "email = ?", userEmail)

	var comment models.Comment
	if result := h.DB.First(&comment, id); result.Error != nil {
		panic(result.Error)
	}

	if comment.UserRefer == user.ID {
		comment.Message = updatedComment.Message
		
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	h.DB.Save(&comment)

	responseUpdate := CreateUpdateCommentResponse(comment)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseUpdate)
}