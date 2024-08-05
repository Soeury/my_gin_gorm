package routers

import (
	"M_GO/GO_gin/project/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// 定义一个路由
	r := gin.Default()

	// 加载静态文件
	r.Static("/static", "D:\\M_GO\\GO_gin\\project\\static")

	// 解析模板
	r.LoadHTMLFiles("D:\\M_GO\\GO_gin\\project\\templates\\index.html")

	// 渲染模板
	r.GET("/", controller.IndexHandler)

	// 路由组完成所有工作
	v1Group := r.Group("v1")
	{
		//  ----添加记录
		v1Group.POST("/todo", controller.CreateToDo)

		//  ----查看所有记录
		v1Group.GET("/todo", controller.GetTodoList)

		//  ----修改某一条记录
		v1Group.PUT("/todo/:id", controller.ModifyTodo)

		//  ----删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return r
}
