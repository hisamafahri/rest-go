package teacher

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
)

func DeleteOneTeacher(c *gin.Context) {
	var teacher []db.Teacher
	id := c.Param("id")

	if err := dbCon.Delete(&teacher, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"details": "Item deleted!",
	})
}
