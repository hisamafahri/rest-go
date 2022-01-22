package teacher

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
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
