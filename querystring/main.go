package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func find(c *gin.Context) {
	//获取携带的query string
	//name:=c.QueryArray("name")
	name, ok := c.GetQuery("name")
	if !ok {
		name = "获取不到数值"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": name,
	})
}
func main() {
	r := gin.Default()
	r.GET("/find", find)
	r.Run(":9000")
}
