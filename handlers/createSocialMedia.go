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

type SocialMediaUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type ResponseSocialMedia struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	SocialMediaURL string          `json:"social_media_url"`
	UserRefer      int             `json:"user_id"`
	User           SocialMediaUser `json:"-"`
	CreatedAt      time.Time       `json:"created_at"`
}

func CreateSocialMediaResponse(s models.SocialMedia) ResponseSocialMedia {
	return ResponseSocialMedia{ID: s.ID, Name: s.Name, SocialMediaURL: s.SocialMediaURL, UserRefer: s.UserRefer, CreatedAt: *s.CreatedAt}
}

func (h Handler) CreateSocialMedia(w http.ResponseWriter, r *http.Request) {

	var data map[string]string

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

	userEmail := claims.(jwt.MapClaims)["email"].(string)

	var user models.User
	h.DB.Where("email = ?", userEmail).First(&user)

	socialMedia := models.SocialMedia{
		Name:           data["name"],
		SocialMediaURL: data["social_media_url"],
		User:           user,
	}

	if result := h.DB.Create(&socialMedia); result.Error != nil {
		panic(result.Error)
	}

	responseCreateSocialMedia := CreateSocialMediaResponse(socialMedia)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseCreateSocialMedia)
}
