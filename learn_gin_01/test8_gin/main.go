package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//   一 ---- gin 框架的模板引擎 , 同原生 http 包 引擎
//   二 ---- 存在多个同名模板文件时 , 为了区分 , 可以建立一个模板文件夹 , 在其下再建立多个文件夹 , 分别存放对应的模板文件
//   三 ---- 自定义函数 : gin中自定义函数同原生 http 包 , 先提前定义一个函数 ， 然后再模板解析之前告诉模板有这个函数就可以
//   四 ---- 静态文件处理 : 静态文件处理 如 html 文件用到的 .css  .js 文件
//   五 ---- gin 不支持 block template 模板继承 (因为很少用 !), 若要实现模板继承 , 需要引入第三方的库

/*

     一 ,
        模板定义
          模板解析
              r.LoadHTMLFiles(路径)
          模板渲染
              func index(c *gin.Context) {
                 name := "chen"
	               c.HTML(http.StatusOK, "index.tpl", gin.H{
		               "name": name,
	               })
              }

   三 ,
          解析模板之前加上自定义函数
              r.SetFuncMap(template.FuncMap{
                  "name" : func_name ,
              })

   四 ,
          再模板解析之前加上  r.Static("/xxx" , "static文件路径")
               .css 文件 一般在<head></head>里面加上 <link ref="stylesheet" href="/xxx/filename.css">
               .js  文件 一般在<body></body>最底下加上 <script src="/xxx/filename.js"></script>


*/

// 模板渲染
func index(c *gin.Context) {

	name := "chen"
	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"name": name,
	})
}

func indexs(c *gin.Context) {

	name := "<a href='http://www.baidu.com' target='_blank'>点我,百度搜索</a>"
	c.HTML(http.StatusOK, "indexs/index.tpl", gin.H{
		"name": name,
	})
}

func posts(c *gin.Context) {

	name := "my name is posts"
	c.HTML(http.StatusOK, "posts/index.tpl", gin.H{
		"name": name,
	})
}

func main() {

	r := gin.Default()

	// 静态文件处理
	r.Static("/xxx", "D:\\M_GO\\GO_gin\\learn_gin_01\\test8_gin\\statics")

	// 自定义函数
	change := func(str string) template.HTML {
		return template.HTML(str)
	}

	r.SetFuncMap(template.FuncMap{
		"change": change,
	})

	//模板解析
	r.LoadHTMLFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test8_gin\\index.tpl")
	//r.LoadHTMLFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test8_gin\\templates\\indexs\\index.tpl")
	//r.LoadHTMLFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test8_gin\\templates\\posts\\index.tpl")

	// 也可以使用这种集群解析的方式 :
	r.LoadHTMLGlob("D:\\M_GO\\GO_gin\\learn_gin_01\\test8_gin\\templates\\**\\*")

	r.GET("/index", index)
	r.GET("/indexs/index", indexs)
	r.GET("/posts/index", posts)
	r.Run(":8080")
}
