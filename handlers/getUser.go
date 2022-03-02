package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jamesthomasw/go-mygram-v1/models"
)

type FindResponse struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Age int `json:"age"`
}

func CreateFindResponse(userModel models.User) FindResponse {
	return FindResponse{ID: userModel.ID, Email: userModel.Email, Username: userModel.Username, Age: userModel.Age}
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["userId"])

	var user models.User

	if result := h.DB.Find(&user, id); result.Error != nil {
		panic(result.Error)
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("user does not exist")
		return
	}

	findResponse := CreateFindResponse(user)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(findResponse)
}