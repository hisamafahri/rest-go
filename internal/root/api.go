package root

import (
	"github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API is up and running!",
	})
}
