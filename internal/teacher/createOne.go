package teacher

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/pkg/db"
)

type teacherCreate struct {
	TeacherID uint   `json:"teacher_id" binding:"required"`
	FullName  string `json:"full_name" binding:"required"`
	Email     string `json:"email"`
}

func CreateOneTeacher(c *gin.Context) {
	var body teacherCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"details": err.Error(),
		})
		return
	}

	data := db.Teacher{TeacherID: body.TeacherID, FullName: body.FullName, Email: body.Email}

	if err := dbCon.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"details": "Data created successfully!",
		"data": gin.H{
			"ID":         data.ID,
			"full_name":  data.FullName,
			"teacher_id": data.TeacherID,
			"email":      data.Email,
		},
	})

}
