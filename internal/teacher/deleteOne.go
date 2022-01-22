package teacher

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
)

func DeleteOneTeacher(c *gin.Context) {
	var teacher []db.Teacher
	id := c.Param("id")

	if err := dbCon.Delete(&teacher, id).Error; err != nil {
		c.JSON(500, gin.H{"error": true, "details": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"error":   false,
			"details": "Item deleted!",
		})
	}
}
