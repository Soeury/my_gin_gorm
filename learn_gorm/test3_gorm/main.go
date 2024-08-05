package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//  ---- gorm 创建记录及字段默认值相关
//  ----  1.设置默认值记得在最开始设置, 否则就相当于更改表结构
//  ----  2. 设置了默认值之后, 所有字段的零值, 如 0 false "" ...都不会被保存到数据库中, 都会采用默认值
//           避免这种情况有两种方法
//           -1 指针 : 在字段类型前加上* , 传入数据时, 记得传入指针
//           -2 databases/sql 包 : 将字段类型设置为 sql.NullType , 传入数据时, 记得传入一个结构体 sql.NullType{Type : value , Valid : true}

// 2.定义一个模型
type User struct {
	ID   int           // 默认成为主键
	Name *string       `gorm:"default:'robot'"`
	Age  sql.NullInt64 `gorm:"default:10"`
}

func main() {

	// 1. 连接数据库
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/my_2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("found err in gorm.Open : %v\n", err)
		return
	}
	defer db.Close()

	// 3. 将模型与数据库进行对应
	db.AutoMigrate(&User{})

	// 4. 创建对象    通过 db.Create(&name) 可以查看主键状态   db.Debug()会将运行过程打印出来
	u1 := User{Name: new(string), Age: sql.NullInt64{Int64: 0, Valid: true}}
	fmt.Println(db.NewRecord(&u1))
	db.Debug().Create(&u1)
	fmt.Println(db.NewRecord(&u1))

}
