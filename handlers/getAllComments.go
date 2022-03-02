package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jamesthomasw/go-mygram-v1/helpers"
	"github.com/jamesthomasw/go-mygram-v1/models"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Photo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserRefer int    `json:"user_id"`
	User      User   `json:"-"`
}

type AllCommentResponse struct {
	ID         int       `json:"id"`
	Message    string    `json:"message"`
	PhotoRefer int       `json:"photo_id"`
	UserRefer  int       `json:"user_id"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
	User       User      `json:"user"`
	Photo      Photo     `json:"photo"`
	
}

func CreateAllCommentResponse(comment models.Comment, user User, photo Photo) AllCommentResponse {
	return AllCommentResponse{ID: comment.ID, Message: comment.Message, PhotoRefer: comment.PhotoRefer, UserRefer: comment.UserRefer, UpdatedAt: *comment.UpdatedAt, CreatedAt: *comment.CreatedAt, Photo: photo, User: user}
}

func (h Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {

	comments := []models.Comment{}
	responseComments := []AllCommentResponse{}

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	userEmail := claims.(jwt.MapClaims)["email"].(string)

	var user models.User
	h.DB.Where("email = ?", userEmail).First(&user)

	h.DB.Where("user_refer = ?", user.ID).Find(&comments)

	var photo models.Photo
	h.DB.Where("user_refer = ?", user.ID).First(&photo)

	userData := User{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
	}

	photoData := Photo{
		ID: photo.ID,
		Title: photo.Title,
		Caption: photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserRefer: user.ID,
		User: userData,
	}

	for _, comment := range comments {
		responseComment := CreateAllCommentResponse(comment, userData, photoData)
		responseComments = append(responseComments, responseComment)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseComments)
}
