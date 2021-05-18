package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type userInfo struct {
	Userid   string `form:"userid" json:"userid"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()
	r.POST("/user", func(context *gin.Context) {
		var user userInfo
		err := context.ShouldBind(&user)
		if err != nil {
			context.JSON(403, gin.H{
				"message": "鉴权错误，错误信息如下",
				"data":    err,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"message":  "success",
				"username": user,
			})
		}

	})
	r.Run(":9000")
}
