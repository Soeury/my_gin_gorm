package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm 模型定义  :

/*
    gorm.Model里面有四个默认字段 :

	type Model struct {
	    ID uint `gorm:"primary_key"`
        CreateAt time.Time
		UpdateAt time.Time
		DeleteAt *time.Time
	}

	调用了 gorm.Model 之后 每次进行 create update delete 的时间就会被  CreateAt  UpdateAt  DeleteAt 记录下来
	软删除 :  调用 delete 之后, DeleteAt 将会设置为当前时间, 不是直接将记录从数据库中删除
*/

//  ----默认表名就是结构体名称的复数, 可以使用函数重新指定一个新的表名, 也可以使用 db.Table("name").CreateTable() 重新指定表名
//  ----默认使用字段名为 ID 的作为主键,比如嵌入了 gorm.Model 之后里面有一个 ID 会成为默认主键,要指定主键需要用tag  `gorm:"primary_key"`
//  ----给默认表名统一加上前缀或者后缀
//  ----默认列名用下划线连接, 可以用 tag 指定列名   `gorm:"column:col_name"`

type Three struct { // 这里默认表名是 Threes
	gorm.Model
	Name    string `gorm:"not null"`
	Age     int    `gorm:"unique"`
	Marrige bool
	Career  string
}

// tag重新指定列名 : 如果在指定列名之前创建了这张表，然后重新指定一个列名, 这时候会重新创建一个列, 不会把原列删除
type Four struct {
	Two_id int    `gorm:"primary_key"`
	Name   string `gorm:"not null"`
	Age    int    `gorm:"column:four_age"`
}

// 函数重新指定表名，如果在指定表明之前创建了一个表 ，然后指定一个新的表名, 这时候会重新创建一个表，不会把原表删除
func (Three) TableName() string {
	return "Three_new_name"
}

func (Four) TableName() string {
	return "Four_new_name"
}

func main() {

	// 给默认表名统一加上前缀或者后缀
	//  *注意 : 这里只会给默认的表名添加 ， 自己重新指定的表名不会添加
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sms_" + defaultTableName
	}

	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/my_2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("found error in  gorm.Open : %v\n", err)
		return
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&Three{})
	db.AutoMigrate(&Four{})

	// table重新指定表名   (这样在添加默认前缀的时候好像会出问题......以后只用方法 1.)
	//db.Table("Four_new_name").CreateTable(&Four{})

}
