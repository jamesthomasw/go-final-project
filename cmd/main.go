package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jamesthomasw/go-mygram-v1/database"
	"github.com/jamesthomasw/go-mygram-v1/handlers"
	"github.com/jamesthomasw/go-mygram-v1/middlewares"
)

func main() {
	DB := database.Connect()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/users/register", h.UserRegister).Methods(http.MethodPost)
	router.HandleFunc("/users/login", h.UserLogin).Methods(http.MethodPost)
	router.HandleFunc("/users", h.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}", h.GetUser).Methods(http.MethodGet)
	router.Handle("/users", middlewares.Authentication(http.HandlerFunc(h.UserUpdate))).Methods(http.MethodPut)
	router.Handle("/users", middlewares.Authentication(http.HandlerFunc(h.UserDelete))).Methods(http.MethodDelete)

	router.Handle("/photos", middlewares.Authentication(http.HandlerFunc(h.CreatePhoto))).Methods(http.MethodPost)
	router.Handle("/photos", middlewares.Authentication(http.HandlerFunc(h.GetAllPhotos))).Methods(http.MethodGet)
	router.Handle("/photos/{photoId}", middlewares.Authentication(http.HandlerFunc(h.GetPhoto))).Methods(http.MethodGet)
	router.Handle("/photos/{photoId}", middlewares.Authentication(http.HandlerFunc(h.UpdatePhoto))).Methods(http.MethodPut)
	router.Handle("/photos/{photoId}", middlewares.Authentication(http.HandlerFunc(h.DeletePhoto))).Methods(http.MethodDelete)

	router.Handle("/comments", middlewares.Authentication(http.HandlerFunc(h.CreateComment))).Methods(http.MethodPost)
	router.Handle("/comments", middlewares.Authentication(http.HandlerFunc(h.GetAllComments))).Methods(http.MethodGet)
	router.Handle("/comments/{commentId}", middlewares.Authentication(http.HandlerFunc(h.GetComment))).Methods(http.MethodGet)
	router.Handle("/comments/{commentId}", middlewares.Authentication(http.HandlerFunc(h.UpdateComment))).Methods(http.MethodPut)
	router.Handle("/comments/{commentId}", middlewares.Authentication(http.HandlerFunc(h.DeleteComment))).Methods(http.MethodDelete)

	router.Handle("/socialmedias", middlewares.Authentication(http.HandlerFunc(h.CreateSocialMedia))).Methods(http.MethodPost)
	router.Handle("/socialmedias", middlewares.Authentication(http.HandlerFunc(h.GetAllSocialMedias))).Methods(http.MethodGet) // This need fix
	router.Handle("/socialmedias/{socialMediaId}", middlewares.Authentication(http.HandlerFunc(h.GetSocialMedia))).Methods(http.MethodGet) // This need fix
	router.Handle("/socialmedias/{socialMediaId}", middlewares.Authentication(http.HandlerFunc(h.UpdateSocialMedia))).Methods(http.MethodPut) // This need fix
	router.Handle("/socialmedias/{socialMediaId}", middlewares.Authentication(http.HandlerFunc(h.DeleteSocialMedia))).Methods(http.MethodDelete) // This need fix

	fmt.Println("Listening to port :8080")
	http.ListenAndServe(":8080", router)
}