package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin 路由和路由组

//   get    : 查询
//   post   : 创建
//   put    : 更新
//   delete : 删除
//   any    : 请求方法的集合 (不建议使用)

//     NoRoute 专门设置用户访问不存在的路由时返回的界面

//    name := r.Group("/共同的路由名字"){  路由集合  }

func main() {

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	//   ANY   请求方法的集合
	r.Any("/any", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET in any"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST in any"})
		case "PUT":
			c.JSON(http.StatusOK, gin.H{"method": "PUT in any"})
		case "DELETE":
			c.JSON(http.StatusOK, gin.H{"method": "DELETE in any"})
		}
	})

	// 用户访问没有定义路由的地址时返回的界面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "404040404040404040404 not found !",
		})
	})

	// 我们可以将拥有共同 url 前缀的路由划分为一个路由组 ，用 {} 包裹同组的路由
	// 路由组 : 多用于区分不同的业务线和 API 版本
	//   注意 : 路由组可以嵌套
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/index",
			})
		})

		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/xx",
			})
		})

		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/oo",
			})
		})
	}

	r.Run(":8080")
}
