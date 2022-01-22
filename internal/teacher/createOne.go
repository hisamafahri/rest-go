package teacher

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
)

type teacher struct {
	TeacherID uint   `json:"teacherID" binding:"required"`
	FullName  string `json:"fullName" binding:"required"`
	Email     string `json:"email"`
}

func CreateOneTeacher(c *gin.Context) {
	var body teacher

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"details": err.Error(),
		})
		return
	}

	data := db.Teacher{TeacherID: body.TeacherID, FullName: body.FullName, Email: body.Email}

	if err := dbCon.Create(&data).Error; err != nil {
		c.JSON(500, gin.H{"error": true, "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"details": "Data created successfully!",
		"data": gin.H{
			"ID":        data.ID,
			"FullName":  data.FullName,
			"TeacherID": data.TeacherID,
			"Email":     data.Email,
		},
	})

}
