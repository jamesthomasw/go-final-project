package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jamesthomasw/go-mygram-v1/models"
)

type UsersResponse struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Age int `json:"age"`
}

func CreateResponseUsers(userModel models.User) UsersResponse {
	return UsersResponse{ID: userModel.ID, Email: userModel.Email, Username: userModel.Username, Age: userModel.Age}
}

func (h Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users := []models.User{}
	h.DB.Find(&users)
	responseUsers := []UsersResponse{}

	for _, user := range users {
		responseUser := CreateResponseUsers(user)
		responseUsers = append(responseUsers, responseUser)
	} 

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseUsers)
}