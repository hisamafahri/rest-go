package teacher

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	} else {
		c.JSON(http.StatusOK, &body)
	}

}
