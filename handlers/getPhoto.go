package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jamesthomasw/go-mygram-v1/models"
)

type PhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserRefer int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateGetPhotoResponse(photo models.Photo) PhotoResponse {
	return PhotoResponse{ID: photo.ID, Title: photo.Title, Caption: photo.Caption, PhotoURL: photo.PhotoURL, UserRefer: photo.UserRefer, CreatedAt: *photo.CreatedAt, UpdatedAt: *photo.UpdatedAt}
}

func (h Handler) GetPhoto(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["photoId"])

	var photo models.Photo

	if result := h.DB.Find(&photo, id); result.Error != nil {
		panic(result.Error)
	}

	responseGetPhoto := CreateGetPhotoResponse(photo)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseGetPhoto)
}
