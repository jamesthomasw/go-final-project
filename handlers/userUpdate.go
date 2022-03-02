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

type UserUpdateResponse struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Age int `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateResponseUserUpdate(newUser models.User) UserUpdateResponse {
	return UserUpdateResponse{ID: newUser.ID, Email: newUser.Email, Username: newUser.Username, Age: newUser.Age, UpdatedAt: *newUser.UpdatedAt}
}

func (h Handler) UserUpdate(w http.ResponseWriter, r *http.Request) {

	type UpdateUser struct {
		Email string `json:"email"`
		Username string `json:"username"`
	}

	var updatedUser UpdateUser

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Failed to read body")
	}

	json.Unmarshal(body, &updatedUser)

	claims, err := helpers.VerifyToken(w, r)
	if err != nil {
		panic(err)
	}

	email := claims.(jwt.MapClaims)["email"].(string)

	var user models.User
	h.DB.Where("email = ?", email).First(&user)

	if user.Email == email {
		user.Email = updatedUser.Email
		user.Username = updatedUser.Username
	
		
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	h.DB.Save(&user)

	responseUpdate := CreateResponseUserUpdate(user)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseUpdate)
}