package teacher

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
	"github.com/lib/pq"
)

func GetOneTeacher(c *gin.Context) {
	var teacher []db.Teacher
	id := c.Param("id")

	if err := dbCon.Find(&teacher, id).Error; err != nil {
		pqErr := err.(*pq.Error)
		c.JSON(500, gin.H{"error": true, "details": pqErr.Message})
	} else {
		c.JSON(200, &teacher)
	}
}
