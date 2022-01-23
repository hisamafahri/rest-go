package authToken

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	var mySigningKey = []byte(jwtSecret)

	token := jwt.New(jwt.SigningMethodHS256)

	claim := token.Claims.(jwt.MapClaims)

	claim["authorized"] = true
	claim["user"] = "John Doe"
	claim["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf("Something wrong with JWT token creation: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
