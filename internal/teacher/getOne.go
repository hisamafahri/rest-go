package teacher

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
)

func GetOneTeacher(c *gin.Context) {
	var teacher []db.Teacher
	id := c.Param("id")

	if err := dbCon.Find(&teacher, id).Error; err != nil {
		c.JSON(500, gin.H{"error": true, "details": err.Error()})
		return
	}

	c.JSON(200, &teacher)
}
