package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jamesthomasw/go-mygram-v1/helpers"
	"github.com/jamesthomasw/go-mygram-v1/models"
	"github.com/dgrijalva/jwt-go"
)

type AllSocialMediasResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserRefer int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User SocialMediaUser `json:"user"`
}

func CreateAllSocialMediasResponse(s models.SocialMedia, user SocialMediaUser) AllSocialMediasResponse{
	return AllSocialMediasResponse{ID: s.ID, Name: s.Name, SocialMediaURL: s.SocialMediaURL, UserRefer: s.UserRefer, CreatedAt: *s.CreatedAt, UpdatedAt: *s.UpdatedAt, User: user}
}

func (h Handler) GetAllSocialMedias(w http.ResponseWriter, r *http.Request) {
	
	var socialMedias []models.SocialMedia
	responseSocialMedias := []AllSocialMediasResponse{}

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	userEmail := claims.(jwt.MapClaims)["email"].(string)

	var user models.User
	h.DB.Where("email = ?", userEmail).First(&user)

	h.DB.Where("user_refer = ?", user.ID).Find(&socialMedias)

	var userData SocialMediaUser
	userData.ID = user.ID
	userData.Username = user.Username

	for _, socialMedia := range socialMedias {
		responseSocialMedia := CreateAllSocialMediasResponse(socialMedia, userData)
		responseSocialMedias = append(responseSocialMedias, responseSocialMedia)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseSocialMedias)
}