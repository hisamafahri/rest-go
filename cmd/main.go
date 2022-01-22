package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/internal/root"
	"github.com/hisamafahri/rest-go/internal/teacher"
)

// var (teacher = &db.Teacher{FullName: "JK Rownling", Email: "jk@rowling.com"})

func main() {
	route := gin.Default()
	route.GET("/", root.GetRoot)
	route.GET("/teacher", teacher.GetAllTeacher)
	route.GET("/teacher/:id", teacher.GetOneTeacher)
	route.DELETE("/teacher/:id", teacher.DeleteOneTeacher)
	route.POST("/teacher", teacher.CreateOneTeacher)
	route.Run(":3000")
}
