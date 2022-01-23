package authToken

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenData := strings.Split(authHeader, " ")

		if authHeader == "" || tokenData[1] == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   true,
				"details": "authorization token required",
			})
			c.Abort()
		} else {
			// to get the token decoded data, change the _ below with `token`
			_, err := jwt.Parse(tokenData[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return mySigningKey, nil
			})

			// if token, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 	// fmt.Println(claims["authorized"], claims["user"])
			// 	c.Next()
			// }
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"error":   true,
					"details": "authentication failed: " + err.Error(),
				})
				c.Abort()
			}
		}
	}
}
