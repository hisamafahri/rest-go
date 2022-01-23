package authToken

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var jwtSecret = os.Getenv("JWT_SECRET")
var mySigningKey = []byte(jwtSecret)

func GenerateToken() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claim := token.Claims.(jwt.MapClaims)

	claim["authorized"] = true
	claim["user"] = "Hisam Fahri"

	claim["iss"] = "http://api.hisamafahri.com/v1/golang"
	claim["exp"] = time.Now().Add(time.Minute * 30).Unix() // token expire in 30 minutes
	claim["iat"] = time.Now().Unix()
	claim["jti"] = uuid.New()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf("Something wrong with JWT token creation: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
