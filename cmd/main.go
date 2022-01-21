package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/db"
	"github.com/hisamafahri/rest-go/internal/root"
	"github.com/hisamafahri/rest-go/internal/teacher"
)

// var (teacher = &db.Teacher{FullName: "JK Rownling", Email: "jk@rowling.com"})

func main() {
	db.DB()
	route := gin.Default()
	route.GET("/", root.GetRoot)
	route.GET("/teacher", teacher.GetAllTeacher)
	route.Run(":3000")

}
