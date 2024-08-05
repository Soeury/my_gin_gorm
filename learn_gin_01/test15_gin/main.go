package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

// gin 单个文件上传

func index(c *gin.Context) {

	c.HTML(http.StatusOK, "file.html", nil)
}

func upload(c *gin.Context) {

	/*
		1.  读取上传的文件

		    f , err := c.FormFile("name")

		2.  将文件保存到本地

			dst := path.Join("路径" , f.Filename)
			c.SaveUploadFile(f , dst)

	*/

	f, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	dst := path.Join("D:\\M_GO\\GO_gin\\learn_gin_01\\test15_gin\\uploadFiles", f.Filename)
	c.SaveUploadedFile(f, dst)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func main() {

	r := gin.Default()

	// 解析模板
	r.LoadHTMLGlob("D:\\M_GO\\GO_gin\\learn_gin_01\\test15_gin\\templates\\*")

	r.GET("/index", index)
	r.POST("/upload", upload)
	r.Run(":8080")
}
