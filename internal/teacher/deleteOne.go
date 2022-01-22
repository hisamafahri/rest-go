package teacher

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
	"github.com/lib/pq"
)

func DeleteOneTeacher(c *gin.Context) {
	var teacher []db.Teacher
	id := c.Param("id")

	if err := dbCon.Delete(&teacher, id).Error; err != nil {
		pqErr := err.(*pq.Error)
		c.JSON(500, gin.H{"error": true, "details": pqErr.Message})
	} else {
		fmt.Println("uo = ", teacher)
		c.JSON(200, gin.H{
			"error":   false,
			"details": "Item deleted!",
		})
	}
}
