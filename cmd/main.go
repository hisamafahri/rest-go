package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/internal/auth"
	"github.com/hisamafahri/rest-go/internal/root"
	"github.com/hisamafahri/rest-go/internal/teacher"
	"github.com/hisamafahri/rest-go/pkg/authToken"
	"github.com/hisamafahri/rest-go/pkg/db"
	"github.com/hisamafahri/rest-go/pkg/helper"
)

func main() {
	db.DB()
	route := gin.Default()

	authorized := route.Group("/")

	authorized.Use(authToken.VerifyToken())
	{
		authorized.GET("/", root.GetRoot)
		authorized.GET("/teacher", teacher.GetAllTeacher)
		authorized.GET("/teacher/:id", teacher.GetOneTeacher)
		authorized.DELETE("/teacher/:id", teacher.DeleteOneTeacher)
		authorized.POST("/teacher", teacher.CreateOneTeacher)
		authorized.PUT("/teacher", teacher.UpdateOneTeacher)
	}

	route.GET("/auth", helper.VerifyBasicAuth(), auth.GenerateToken)
	route.Run(":3000")
}
