package teacher

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
)

var dbCon = db.DB()

func GetAllTeacher(c *gin.Context) {
	var teacher []db.Teacher

	if err := dbCon.Find(&teacher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &teacher)
}
