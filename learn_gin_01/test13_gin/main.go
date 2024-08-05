package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin 获取 url 路径参数

// 请求的参数通过 url 路径传递  ,  例如 :  /web/student/clotehs/size/red

func people(c *gin.Context) {

	// 获取 url 参数  :   temp := c.Param("路径上每两个 / 之间对应的字段名")
	// 这里返回的都是字符串
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func blog(c *gin.Context) {

	year := c.Param("year")
	month := c.Param("month")

	// 这里用 json 表示主要为了拿到 url 中传入的参数 ， 这里主要是为了拿到 ， 不是为了使用
	c.JSON(http.StatusOK, gin.H{
		"year":  year,
		"month": month,
	})
}

func main() {

	r := gin.Default()

	// 获取路径参数  在需要获取的字段前加上:
	r.GET("/:name/:age", people)
	r.GET("/blog/:year/:month", blog)

	r.Run(":8080")
}
