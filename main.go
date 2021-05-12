package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayhello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang",
	})
}
func book(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "POST",
	})
}
func deleteBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "delete",
	})
}
func creat_book(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "creat",
	})
}
func remove_book(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "put",
	})
}
func main() {
	r := gin.Default() //返回一个引擎
	r.GET("/hello", sayhello)
	r.DELETE("/book", deleteBook)
	r.POST("/creat_book", creat_book)
	r.PUT("/remove_book", remove_book)
	//启动服务
	r.Run(":9090")
}
