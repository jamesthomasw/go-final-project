package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jamesthomasw/go-mygram-v1/helpers"
	"github.com/jamesthomasw/go-mygram-v1/models"
)

type PhotoModel struct {
	ID int `json:"photo_id"`
}

type UserModel struct {
	ID int `json:"user_id"`
}

type CommentResponse struct {
	ID         int        `json:"id"`
	Message    string     `json:"message"`
	PhotoRefer int        `json:"photo_id"`
	UserRefer  int        `json:"user_id"`
	Photo      PhotoModel `json:"-"`
	User       UserModel  `json:"-"`
	CreatedAt  time.Time  `json:"created_at"`
}

func CreateCommentResponse(c models.Comment) CommentResponse {
	return CommentResponse{ID: c.ID, Message: c.Message, PhotoRefer: c.PhotoRefer, UserRefer: c.UserRefer, CreatedAt: *c.CreatedAt}
}

func (h Handler) CreateComment(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &data)

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	var user models.User
	userEmail := claims.(jwt.MapClaims)["email"].(string)

	h.DB.Where("email = ?", userEmail).First(&user)

	photoId := int(data["photo_id"].(float64))

	var photo models.Photo
	h.DB.Where("id = ?", photoId).First(&photo)

	photo.User = user

	comment := models.Comment{
		Message: data["message"].(string),
		Photo:   photo,
		User:    user,
	}

	if res := h.DB.Create(&comment); res.Error != nil {
		panic(res.Error)
	}

	responseCreateComment := CreateCommentResponse(comment)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseCreateComment)
}
