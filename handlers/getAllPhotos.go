package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jamesthomasw/go-mygram-v1/helpers"
	"github.com/jamesthomasw/go-mygram-v1/models"
	"github.com/dgrijalva/jwt-go"
)

type AllPhotoResponse struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserRefer int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User UserData `json:"user"`
}

func CreateAllPhotoResponse(photoModel models.Photo, user UserData) AllPhotoResponse {
	return AllPhotoResponse{ID: photoModel.ID, Title: photoModel.Title, Caption: photoModel.Caption, PhotoURL: photoModel.PhotoURL, UserRefer: photoModel.UserRefer, User: user, CreatedAt: *photoModel.CreatedAt, UpdatedAt: *photoModel.UpdatedAt}
}

func (h Handler) GetAllPhotos(w http.ResponseWriter, r *http.Request) {
	
	photos := []models.Photo{}
	responsePhotos := []AllPhotoResponse{}

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	userEmail := claims.(jwt.MapClaims)["email"].(string)

	var user models.User
	h.DB.Where("email = ?", userEmail).First(&user)

	h.DB.Where("user_refer = ?", user.ID).Find(&photos)

	var userData UserData
	userData.Email = user.Email
	userData.Username = user.Username

	for _, photo := range photos {
		responsePhoto := CreateAllPhotoResponse(photo, userData)
		responsePhotos = append(responsePhotos, responsePhoto)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responsePhotos)
}