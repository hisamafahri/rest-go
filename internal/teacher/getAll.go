package teacher

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
)

var dbCon = db.DB()

func GetAllTeacher(c *gin.Context) {
	var teacher []db.Teacher

	if err := dbCon.Find(&teacher).Error; err != nil {
		c.JSON(500, gin.H{"error": true, "details": err.Error()})
	} else {
		c.JSON(200, &teacher)
	}
}
