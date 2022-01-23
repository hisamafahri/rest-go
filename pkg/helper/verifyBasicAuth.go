package helper

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyBasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := os.Getenv("AUTH_USERNAME")
		password := os.Getenv("AUTH_PASSWORD")

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   true,
				"details": "authorization credentials required",
			})
			c.Abort()
		} else {
			authData := strings.Split(authHeader, " ")
			var decodedByte, _ = base64.StdEncoding.DecodeString(authData[1])
			if string(decodedByte) == username+":"+password {
				c.Next()
			} else {
				c.JSON(http.StatusForbidden, gin.H{
					"error":   true,
					"details": "authorization failed",
				})
				c.Abort()
			}
		}
	}
}
