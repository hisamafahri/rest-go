package teacher

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
)

type teacherUpdate struct {
	TeacherID uint   `json:"teacher_id" binding:"required"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
}

func UpdateOneTeacher(c *gin.Context) {
	var body teacherUpdate
	var teacher []db.Teacher

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"details": err.Error(),
		})
		return
	}

	if err := dbCon.Model(&teacher).Where("teacher_id = ?", body.TeacherID).Updates(teacherUpdate{FullName: body.FullName, Email: body.Email}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"details": "Data updated successfully!",
	})
}
