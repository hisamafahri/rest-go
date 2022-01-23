package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/authToken"
)

func GenerateToken(c *gin.Context) {
	token, err := authToken.GenerateToken()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"token": token,
	})
}
