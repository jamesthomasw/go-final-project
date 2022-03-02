package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jamesthomasw/go-mygram-v1/models"
)

type GetSocialMediaResponse struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	SocialMediaURL string          `json:"social_media_url"`
	UserRefer      int             `json:"user_id"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	User           SocialMediaUser `json:"-"`
}

func CreateGetSocialMediaResponse(s models.SocialMedia) GetSocialMediaResponse {
	return GetSocialMediaResponse{ID: s.ID, Name: s.Name, SocialMediaURL: s.SocialMediaURL, UserRefer: s.UserRefer, CreatedAt: *s.CreatedAt, UpdatedAt: *s.UpdatedAt}
}

func (h Handler) GetSocialMedia(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["socialMediaId"])

	var socialMedia models.SocialMedia

	if result := h.DB.Find(&socialMedia, id); result.Error != nil {
		panic(result.Error)
	}

	responseGetSocialMedia := CreateGetSocialMediaResponse(socialMedia)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseGetSocialMedia)
}
