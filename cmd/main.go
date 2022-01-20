package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hisamafahri/rest-go/internal/root"
)

func main() {
	r := gin.Default()
	r.GET("/", root.GetRoot)
	r.Run(":3000")
}
