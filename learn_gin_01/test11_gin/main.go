package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin 获取 form 参数

// 模板渲染
func login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil) // nil表示不需要传入对象 , 需要传入对象用  gin.H{ }
}

func login2(c *gin.Context) {

	// 1. 使用  temp := c.PostForm("key") 返回表单参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 2. 使用 temp := c.DefaultPostForm("key" , "default")
	username1 := c.DefaultPostForm("username", "somebody")
	password1 := c.DefaultPostForm("password", "***")

	// 3. 使用 temp , ok := c.GetPostForm("key")
	username2, ok := c.GetPostForm("username")
	if !ok {
		username2 = "somebody"
	}

	password2, ok := c.GetPostForm("password")
	if !ok {
		password2 = "***"
	}

	c.HTML(http.StatusOK, "postform.html", gin.H{
		"username":  username,
		"password":  password,
		"username1": username1,
		"password1": password1,
		"username2": username2,
		"password2": password2,
	})
}

func main() {

	r := gin.Default()

	// 模板解析
	r.LoadHTMLGlob("D:\\M_GO\\GO_gin\\learn_gin_01\\test11_gin\\templates\\*")

	r.GET("/login", login)
	r.POST("/login", login2)

	r.Run(":8080")
}
