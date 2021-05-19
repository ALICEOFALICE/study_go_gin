package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func m1(c *gin.Context) {
	fmt.Println("进入m1中间件")
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Println("耗时:", cost)
	fmt.Println("退出中间件")
}
func result(c *gin.Context) {
	fmt.Println("进入返回数据函数")
	c.JSON(http.StatusOK, gin.H{
		"message": "您已经通过中间件校验",
	})
}
func main() {
	r := gin.Default()
	r.GET("/zjj", m1, result)
	r.Run(":9000")
}
