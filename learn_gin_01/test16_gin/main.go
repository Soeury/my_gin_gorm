package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

// gin 多个文件上传

func temp(c *gin.Context) {

	c.HTML(http.StatusOK, "file.html", nil)
}

func up(c *gin.Context) {

	/*
		1. 多文件上传  :

		form , err := c.MultipartForm()
		files := form.File["name"]

		2. 文件保存到本地 :

		for _, file := range files {

		    dst := path.Join("路径", file.Filename)
		    c.SaveUploadedFile(file, dst)

		}


	*/

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	files := form.File["fn"]

	for _, file := range files {

		dst := path.Join("D:\\M_GO\\GO_gin\\learn_gin_01\\test16_gin\\upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func main() {

	r := gin.Default()

	//模板解析
	r.LoadHTMLGlob("D:\\M_GO\\GO_gin\\learn_gin_01\\test16_gin\\templates\\*")

	r.GET("/temp", temp)
	r.POST("/up", up)
	r.Run(":8080")
}
