package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//  这里我用来测试一下 模板直接用  .html 文件是否可以

//    事实证明  :  可以 !

func test(c *gin.Context) {

	// 注意 : gin 框架里也会转义
	name := "<script>alert(123);</script>"
	c.HTML(http.StatusOK, "file.html", gin.H{
		"name": name,
	})
}

func main() {

	r := gin.Default()

	// 解析模板
	r.LoadHTMLFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test12_gin\\templates\\file.html")

	r.GET("/test", test)
	r.Run(":8080")

}
