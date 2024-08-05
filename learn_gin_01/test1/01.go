package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func r_get(c *gin.Context) {
	// gin.H 实际上是一个map    map[string]any

	c.JSON(200, gin.H{
		"message": "R_GET",
	})
}

func r_post(c *gin.Context) {
	// 这里的状态码 200 改成 http.StatusOK 也可以

	c.JSON(http.StatusOK, gin.H{
		"message": "R_POST",
	})
}

func r_put(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "R_PUT",
	})
}

func r_delete(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "R_DELETE",
	})
}

func main() {

	r := gin.Default()
	r.GET("/hello", r_get)
	r.POST("/hello", r_post)
	r.PUT("/hello", r_put)
	r.DELETE("/hello", r_delete)

	// gin 框架支持 restful api 开发 ：用 url 定位资源，用 http 动词 get , post , put , delete 描述操作

	r.Run(":8080")
}
