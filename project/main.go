package main

import (
	"M_GO/GO_gin/project/dao"
	"M_GO/GO_gin/project/models"
	"M_GO/GO_gin/project/routers"
	"fmt"
)

func main() {

	// 创建数据库  create database bubble;
	// 连接数据库
	err := dao.InitMySql()
	if err != nil {
		fmt.Printf("found err in initMysql : %v\n", err)
		return
	}
	defer dao.DB.Close()

	// 建立对应关系
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	r.Run(":8080")
}
