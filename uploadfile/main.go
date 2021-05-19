package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("up.html")
	r.GET("/image", func(context *gin.Context) {
		context.HTML(200, "up.html", nil)
	})

	r.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("upload_file")
		if err != nil {
			fmt.Print(err)

			context.JSON(http.StatusOK, gin.H{
				"message": "发生错误，具体信息如下",
				"data":    err,
			})
		} else {
			dst := fmt.Sprintf("./uploads/%s", file.Filename)
			context.SaveUploadedFile(file, dst)
			context.JSON(http.StatusOK, gin.H{
				"message":   "上传成功",
				"fileName":  file.Filename,
				"file_size": file.Size,
			})
		}

	})
	r.Run(":9000")
}
