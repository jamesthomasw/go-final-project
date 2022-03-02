package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jamesthomasw/go-mygram-v1/helpers"
	"github.com/jamesthomasw/go-mygram-v1/models"
)

func (h Handler) DeleteSocialMedia(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["socialMediaId"])

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	userEmail := claims.(jwt.MapClaims)["email"].(string)

	var user models.User
	h.DB.Where("email = ?", userEmail).First(&user)

	var socialMedia models.SocialMedia
	h.DB.Where("id = ?", id).First(&socialMedia)

	if socialMedia.UserRefer == user.ID {
		h.DB.Delete(&socialMedia)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Your social media has been successfully deleted")
}