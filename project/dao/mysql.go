package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 初始化一个全局对象 db
var DB *gorm.DB

func InitMySql() (err error) {
	dsn := "root:123456@(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("found err in gorm.Open : %v\n", err)
		return
	}
	return DB.DB().Ping()
}
