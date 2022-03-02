package helpers

import "github.com/dgrijalva/jwt-go"

const SecretKey = "SuperSecretPhrase"

func GenerateToken(email, username string) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"username": username,
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		panic(err)
	}

	return token
}