package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jamesthomasw/go-mygram-v1/helpers"
	"github.com/jamesthomasw/go-mygram-v1/models"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	Token string `json:"token"`
}

func (h Handler) UserLogin(w http.ResponseWriter, r *http.Request) {

	var data map[string]string

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Failed to read body while user logging in")
	}

	json.Unmarshal(body, &data)

	var user models.User

	h.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		log.Fatalln("User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		log.Fatalln("Incorrect password")
	}

	token := Token{
		Token: helpers.GenerateToken(user.Email, user.Username),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}