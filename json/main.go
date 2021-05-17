package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "helloworld",
	})
}
func main() {
	type userInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		ID   int    `json:"id"`
	}
	m := userInfo{
		Name: "谢",
		Age:  18,
		ID:   20313,
	}
	r := gin.Default()
	r.GET("/", test)
	r.GET("/userInfo", func(c *gin.Context) {
		c.JSON(200, m)
	})
	fmt.Printf("开始启动服务,端口如下\n")
	fmt.Printf("127.0.0.1:9000\\n")
	r.Run(":9000")
}
