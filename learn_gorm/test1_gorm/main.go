package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Name   string
	Age    int
	Gender string
	Hobby  string
}

func main() {

	// 连接 MySql 数据库
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/my_1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("found error in gorm.Open : %v\n", err)
		return
	}
	defer db.Close()

	// 创建表 自动迁移 (结构体 -> 数据库)
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	u1 := UserInfo{"chen", 18, "high", "love"}
	db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("u : %#v\n", u)

	// 更新
	db.Model(&u).Update("hobby", "dance")

	// 删除
	db.Delete(&u)

}
