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

type UpdatePhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserRefer int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUpdatePhotoResponse(photo models.Photo) UpdatePhotoResponse {
	return UpdatePhotoResponse{ID: photo.ID, Title: photo.Title, Caption: photo.Caption, PhotoURL: photo.PhotoURL, UserRefer: photo.UserRefer, UpdatedAt: *photo.UpdatedAt}
}

func (h Handler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	
	type UpdatePhoto struct {
		Title     string    `json:"title"`
		Caption   string    `json:"caption"`
		PhotoURL  string    `json:"photo_url"`
	}
	
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["photoId"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var updatedPhoto UpdatePhoto
	json.Unmarshal(body, &updatedPhoto)

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	var user models.User

	userEmail := claims.(jwt.MapClaims)["email"].(string)
	h.DB.First(&user, "email = ?", userEmail)

	var photo models.Photo
	if result := h.DB.First(&photo, id); result.Error != nil {
		panic(result.Error)
	}

	if photo.UserRefer == user.ID {
		photo.Title = updatedPhoto.Title
		photo.Caption = updatedPhoto.Caption
		photo.PhotoURL = updatedPhoto.PhotoURL
		h.DB.Debug().Save(&photo)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	responseUpdate := CreateUpdatePhotoResponse(photo)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseUpdate)
}