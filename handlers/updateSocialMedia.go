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

type UpdateSocialMediaResponse struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	SocialMediaURL string          `json:"social_media_url"`
	UserRefer      int             `json:"user_id"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

func CreateUpdateSocialMediaResponse(s models.SocialMedia) UpdateSocialMediaResponse{
	return UpdateSocialMediaResponse{ID: s.ID, Name: s.Name, SocialMediaURL: s.SocialMediaURL, UserRefer: s.UserRefer, UpdatedAt: *s.UpdatedAt}
}

func (h Handler) UpdateSocialMedia(w http.ResponseWriter, r *http.Request) {

	type UpdateSocialMedia struct {
		Name string `json:"name"`
		SocialMediaURL string `json:"social_media_url"`
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["socialMediaId"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var updatedSocialMedia UpdateSocialMedia
	json.Unmarshal(body, &updatedSocialMedia)

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	var user models.User
	userEmail := claims.(jwt.MapClaims)["email"].(string)
	h.DB.First(&user, "email = ?", userEmail)

	var socialMedia models.SocialMedia
	if result := h.DB.First(&socialMedia, id); result.Error != nil {
		panic(result.Error)
	}

	if socialMedia.UserRefer == user.ID {
		socialMedia.Name = updatedSocialMedia.Name
		socialMedia.SocialMediaURL = updatedSocialMedia.SocialMediaURL

		

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	h.DB.Save(&socialMedia)

	responseUpdateData := CreateUpdateSocialMediaResponse(socialMedia)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseUpdateData)
}