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

func (h Handler) DeletePhoto(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["photoId"])

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	userEmail := claims.(jwt.MapClaims)["email"].(string)

	var user models.User
	h.DB.Where("email = ?", userEmail).First(&user)

	var photo models.Photo
	h.DB.Where("id = ?", id).First(&photo)

	if photo.UserRefer == user.ID {
		h.DB.Delete(&photo)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Your photo has been successfully deleted")
}