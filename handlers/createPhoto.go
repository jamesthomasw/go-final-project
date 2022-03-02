package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jamesthomasw/go-mygram-v1/helpers"
	"github.com/jamesthomasw/go-mygram-v1/models"
)

type UserData struct {
	Email string `json:"email"`
	Username string `json:"username"`
}

type ResponseCreate struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserRefer int `json:"user_id"`
	User UserData `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

func CreatePhotoResponse(photo models.Photo, user UserData) ResponseCreate {
	return ResponseCreate{ID: photo.ID, Title: photo.Title, Caption: photo.Caption, PhotoURL: photo.PhotoURL, UserRefer: photo.UserRefer, User: user, CreatedAt: *photo.CreatedAt}
}

func (h Handler) CreatePhoto(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Failed to read body")
	}

	json.Unmarshal(body, &data)

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	userEmail := claims.(jwt.MapClaims)["email"].(string)

	var user models.User
	h.DB.Where("email = ?", userEmail).First(&user)
	
	photo := models.Photo{
		Title: data["title"].(string),
		Caption: data["caption"].(string),
		PhotoURL: data["photo_url"].(string),
		User: user,
	}

	if result := h.DB.Create(&photo); result.Error != nil {
		panic(result.Error)
	}

	var userData UserData
	userData.Email = user.Email
	userData.Username = user.Username

	responseCreatePhoto := CreatePhotoResponse(photo, userData)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseCreatePhoto)
}