package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("login.html")
	r.GET("/ping", func(context *gin.Context) {
		context.HTML(200, "login.html", nil)
	})
	r.POST("/ping_add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(200, gin.H{
			"username": username,
			"password": password,
		})
	})
	fmt.Println("http://127.0.0.1:9000")
	r.Run(":9000")
}
