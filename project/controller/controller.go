package controller

import (
	"M_GO/GO_gin/project/dao"
	"M_GO/GO_gin/project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateToDo(c *gin.Context) {
	// 前端页面填写代办事项, 点击提交, 会发送请求到这里
	//  1. 将数据从请求中拿出来
	var todo models.Todo
	c.BindJSON(&todo)

	//  2. 数据存入数据库
	err := dao.DB.Create(&todo).Error

	//  3. 返回响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	// 查询todo表中所有的数据
	var todolist []models.Todo
	err := dao.DB.Find(&todolist).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func ModifyTodo(c *gin.Context) {
	// 1.先查询数据
	//    -- 拿到需要修改的数据的 id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"err": "id not found"})
		return
	}

	//   -- 拿到这个id对应的一行数据
	var todo models.Todo
	err := dao.DB.Where("id=?", id).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}

	// 2.后修改数据
	c.BindJSON(&todo)
	err = dao.DB.Save(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	// 先拿到需要删除的数据的id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"err": "id not exist"})
		return
	}

	// 删除
	var temp []models.Todo
	err := dao.DB.Where("id=?", id).Delete(temp).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
