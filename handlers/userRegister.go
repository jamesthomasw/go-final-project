package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/jamesthomasw/go-mygram-v1/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterResponse struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Age int `json:"age"`
}

func CreateUserRegisterResponse(userModel models.User) UserRegisterResponse {
	return UserRegisterResponse{ID: userModel.ID, Email: userModel.Email, Username: userModel.Username, Age: userModel.Age}
}

func (h Handler) UserRegister(w http.ResponseWriter, r *http.Request) {

	var data map[string]string

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Failed to read body while registering user")
	}

	json.Unmarshal(body, &data)

	var user models.User

	pass, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		log.Fatalln("Failed to hash password")
	}

	age, err := strconv.Atoi(data["age"])
	if err != nil {
		panic(err)
	}

	user = models.User{
		Username: data["username"],
		Email: data["email"],
		Password: string(pass),
		Age: age,
	}

	if result := h.DB.Create(&user); result.Error != nil {
		panic(result.Error)
	}

	responseRegister := CreateUserRegisterResponse(user)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseRegister)
}