package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jamesthomasw/go-mygram-v1/helpers"
	"github.com/jamesthomasw/go-mygram-v1/models"
	"github.com/dgrijalva/jwt-go"
)

func (h Handler) UserDelete(w http.ResponseWriter, r *http.Request) {

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	email := claims.(jwt.MapClaims)["email"].(string)

	var user models.User
	h.DB.Where("email = ?", email).First(&user)
	h.DB.Delete(&user)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Your account has been succesfully deleted")
}