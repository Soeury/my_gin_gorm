package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//  ---- gin 中间件  :
//   路由中间件的设置
//   路由组中间件的设置
//   中间件常用闭包结构
//   c.Set()  c.Get() 取出中间件中设置的上下文的值

//  ---- gin.Default() 会默认采用两个中间件 logger 和 recovery
//      ---- 1. logger 中间件会将日志写入 Gin.DefaultWritter
//      ---- 2. recovery 中间件会recover 任何 panic , 会写入 500 响应码
//  ---- 如果不想采用这两个中间件，请使用   r := gin.New()  配置默认路由

//  *注意 : 在中间件中使用  goroutine 时 ， 不能使用上下文中的 (c *gin.Context) , 只能使用  c.Copy() 这个对象

func testFunc(c *gin.Context) {

	// 取出上文中设置的值    value , ok := c.Get(key)
	temp, ok := c.Get("age")
	if !ok {
		temp = 20
	}

	fmt.Println("this is in testFunc")
	c.JSON(http.StatusOK, gin.H{
		"message": temp,
	})
}

// 定义一个中间件 medium_1 : 统计请求处理函数的耗时
func medium_1(c *gin.Context) {

	fmt.Println("medium_1 in ...")
	start := time.Now()

	c.Next() // 表示调用后续的处理函数
	// c.Abort() // 表示阻止调用后续的处理函数

	cost := time.Since(start)
	fmt.Printf("cost : %v\n", cost)
	fmt.Println("medium_1 end ...")
}

func medium_2(c *gin.Context) {
	fmt.Println("medium_2 in ...")

	// c.Set(key , value) 表示在上下文中设置值 key = value
	c.Set("age", 18)

	c.Next()
	fmt.Println("medium_2 end ...")
}

func medium_3(c *gin.Context) {
	fmt.Println("medium_3 in ...")
	// c.Abort() // 阻止调用后续函数
	// return
	fmt.Println("medium_3 end ...")
}

// 中间件一般采用闭包的结构 : 传入 bool 参数 , true 表示检查 ， false 表示不检查
func medium_4(doCheck bool) gin.HandlerFunc {

	// 数据库的连接...
	// 其他操作...

	return func(c *gin.Context) {
		if doCheck {
			//   是否登录的判断
			//   if 是登陆用户
			//   c.Next()
			//   else
			//   c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {

	r := gin.Default()

	// 这里可以传入多个函数，但是会按顺序执行
	// 在处理函数之前加上中间件函数就可以，不需要加上中间件的可以不用加上

	//  r.Use(func_name) 表示全局注册中间件函数, 也是按照顺序执行的(最先调用的是这些中间件函数)
	r.Use(medium_1, medium_2, medium_3, medium_4(false))

	r.GET("/test", testFunc)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "shop",
		})
	})

	r.GET("/video", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "video",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "40404 not found",
		})
	})

	// 路由组添加中间件  方法1 :   group_name := r.Group("/name" , middleware_name...)
	videoGroup := r.Group("/video", medium_1)
	{
		videoGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/login",
			})
		})

		videoGroup.GET("/myself", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/myself",
			})
		})
	}

	// 路由组添加中间件  方法2 :    group_name.Use(middleware_name ...)
	bookGroup := r.Group("/book")
	bookGroup.Use(medium_2, medium_3)
	{
		bookGroup.GET("/buy", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/book/buy",
			})
		})

		bookGroup.GET("/read", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/book/read",
			})
		})

	}

	r.Run(":8080")
}
