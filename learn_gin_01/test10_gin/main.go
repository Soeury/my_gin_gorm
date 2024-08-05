package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin 获取 querystring 参数 : 一般用在 get 请求中

//   querystring 参数 : 指的是 http 请求 url 中 ? 后面的参数
//   采用 key = value 的形式 , 多个 key = value 之间用 & 符号连接
//   e.g.      localhost:8080/web?query=chen&age=18

func web(c *gin.Context) {

	// 1.   temp := c.Query( "key_name" ) 获取请求中携带的 querystring 参数
	// 注意 这里面的 "key_name" 表示 url 中 key = value 的 key 的名字 , 传入的参数会被赋值给返回值  ， 返回值使用在 json 中

	name := c.Query("query")
	age := c.Query("age")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func web2(c *gin.Context) {

	// 2.  temp := c.DefaultQuery( "key_name" , "default")

	name := c.DefaultQuery("query", "sb")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func web3(c *gin.Context) {

	// 3.  temp , ok := c.GetQuery( "key_name" )

	name, ok := c.GetQuery("query")
	if !ok {
		fmt.Printf("found error in c.GetQuery")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func main() {

	r := gin.Default()

	r.GET("/web", web)
	r.GET("/web2", web2)
	r.GET("/web3", web3)

	r.Run(":8080")
}
