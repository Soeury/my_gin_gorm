package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//  ---- gin 请求重定向 和 路由重定向
//  用途 : 一般在第一次去到某个网站的时候，是一个未登录的状态，这时候点击个人中心，会跳转到登录的界面，登录完成之后跳转回来

//   我觉得这里有必要解释一下 重定向 和 请求转发
//
//       * 重定向   : 产生两次请求 两次响应 通过浏览器的重新请求来产生跳转 可以跳转到服务器外部
//       * 请求转发 : 产生一次请求 一次响应 只能在服务器内部跳转

func index(c *gin.Context) {

	// 这是正常定向 ， 客户端指定哪里就是哪里
	//c.JSON(http.StatusOK, gin.H{
	//	"status": "ok",
	//})

	// 请求重定向 :  发送两次请求
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}

func main() {

	r := gin.Default()

	r.GET("/index", index)

	// 路由重定向    "/a"     ->    "/b"
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" // 修改请求的 url
		r.HandleContext(c)        // 继续后续的处理
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":8080")
}
