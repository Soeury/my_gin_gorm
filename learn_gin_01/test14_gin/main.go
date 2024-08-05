package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//  gin 参数绑定  :     err := c.ShouldBlind(&结构体对象)
//  把query中的用户输入的字段一个一个拿出来有时候会很繁琐 ， 这时候需要用到 gin 的参数绑定 ， 能够更加方便的获取请求中携带的参数

// 1.原始一个一个拿出数据的方式
func user(c *gin.Context) {

	type UserInfo struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}

	// 注意 : 这里返回的都是string类型的数据
	username2 := c.Query("username")
	password2 := c.Query("password")

	u := UserInfo{
		Username: username2,
		Password: password2,
	}

	fmt.Printf("%#v\n", u)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// 2.使用 ShouldBlind 取出 QueryString 类型的数据
func user2(c *gin.Context) {

	// 注意这里的 tag 使用 form 还是 json  :   json 是返回给前端的 tag 标签 , form 是c *gin.context 中取出对应的绑定参数
	//    注意格式 :      `form:"字段名"`
	type UserInfo struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}

	// 注意 : 这里   err := c.ShouldBind(&u) 记得用引用的方式传入结构体对象
	var u UserInfo
	err := c.ShouldBind(&u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	} else {
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

// 3.使用 ShouldBlind 取出 Form 类型的数据
func form(c *gin.Context) {

	type UserInfo struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}

	var u2 UserInfo
	err := c.ShouldBind(&u2)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	} else {
		fmt.Printf("%#v\n", u2)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

func index(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
}

func json(c *gin.Context) {

	type UserInfo struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}

	var u3 UserInfo
	err := c.ShouldBind(&u3)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	} else {
		fmt.Printf("%#v\n", u3)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

func main() {

	r := gin.Default()

	// 解析模板
	r.LoadHTMLGlob("D:\\M_GO\\GO_gin\\learn_gin_01\\test14_gin\\templates\\*")

	r.GET("/user", user)
	r.GET("/index", index)
	r.GET("/user2", user2)
	r.POST("/form", form)
	r.POST("/json", json)
	r.Run(":8080")
}
