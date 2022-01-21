package teacher

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/db"
)

func GetAllTeacher(c *gin.Context) {

	dbCon := db.DB()

	var teacher []db.Teacher
	dbCon.Find(&teacher)

	c.JSON(200, &teacher)
}
