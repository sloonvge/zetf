package main

import (
	"github.com/gin-gonic/gin"
	"response"
	)

func main() {
	r := gin.Default()
	r.GET("/test", response.Login)
	r.Run()
}
