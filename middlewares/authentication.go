package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/jamesthomasw/go-mygram-v1/helpers"
)

type ErrMessage struct {
	Error string `json:"error"`
	Message string `json:"message"`
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		verifyToken, err := helpers.VerifyToken(w, r)
		_ = verifyToken

		if err != nil {
			errMessage := ErrMessage{
				Error: "Unauthenticated",
				Message: err.Error(),
			}

			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errMessage)
			return
		}

		next.ServeHTTP(w, r)
	})
}