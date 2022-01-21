package teacher

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/db"
	"github.com/lib/pq"
)

var dbCon = db.DB()

func GetAllTeacher(c *gin.Context) {
	var teacher []db.Teacher

	if err := dbCon.Find(&teacher).Error; err != nil {
		pqErr := err.(*pq.Error)
		c.JSON(500, gin.H{"error": true, "details": pqErr.Message})
	} else {
		c.JSON(200, &teacher)
	}
}

func GetATeacher(c *gin.Context) {
	var teacher []db.Teacher
	id := c.Param("id")

	if err := dbCon.Find(&teacher, id).Error; err != nil {
		pqErr := err.(*pq.Error)
		c.JSON(500, gin.H{"error": true, "details": pqErr.Message})
	} else {
		c.JSON(200, &teacher)
	}
}
