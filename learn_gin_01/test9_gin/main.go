package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin 框架返回 json 格式的数据    ->    json 渲染

//    1. map[string]interface{}
//    2. gin.H{}
//    3. type name struct {}

func get_json1(c *gin.Context) {

	// 1. 使用 map[string]interface{} 返回 json 数据
	data := map[string]interface{}{
		"name":  "chen",
		"age":   18,
		"money": 26000,
	}
	c.JSON(http.StatusOK, data)
}

func get_json2(c *gin.Context) {

	// 2. 使用 gin.H{} 返回  json 格式的数据  ( 实际上gin.H本来就是 map[string]interface{} 类型)
	data := gin.H{
		"name":  "boliang",
		"age":   20,
		"level": 100,
	}

	c.JSON(http.StatusOK, data)
}

func get_json3(c *gin.Context) {

	// 3. 使用 struct 结构体返回 json 格式的数据 (原来结构体里面不需要使用 , 分开)
	type Student struct {
		Name   string
		Age    int
		Height int
	}

	data := Student{"piki", 18, 168}
	c.JSON(http.StatusOK, data)
}

func main() {

	r := gin.Default()

	r.GET("/json1", get_json1)
	r.GET("/json2", get_json2)
	r.GET("/json3", get_json3)

	r.Run(":8080")
}
