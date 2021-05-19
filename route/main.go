package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Any("/key", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	fmt.Printf("启动端口为\n")
	fmt.Print("127.0.0.1:9000")
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/xhp", func(context *gin.Context) {
			r.LoadHTMLFiles("login.html")
			context.HTML(http.StatusOK, "login.html", nil)
		})
		videoGroup.POST("/ping_add", func(context *gin.Context) {
			username := context.PostForm("username")
			password := context.PostForm("password")
			context.JSON(http.StatusOK, gin.H{
				"username": username,
				"password": password,
			})
		})
	}
	r.Run(":9000")

}
