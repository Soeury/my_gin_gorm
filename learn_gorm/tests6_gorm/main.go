package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm 删除

// 建立模型
type Teacher struct {
	gorm.Model
	Name     string
	Standard int
	Salary   int
}

func main() {

	// 连接数据库
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/my_2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("found err in form.Open : %v\n", err)
		return
	}
	defer db.Close()

	// 建立对应关系
	db.AutoMigrate(&Teacher{})

	// 创建数据
	u1 := Teacher{Name: "hh", Standard: 1, Salary: 5000}
	u2 := Teacher{Name: "ff", Standard: 3, Salary: 9000}
	u3 := Teacher{Name: "dd", Standard: 7, Salary: 18000}
	db.Create(&u1)
	db.Create(&u2)
	db.Create(&u3)

	// 删除 : 一般都是通过主键删除 , gorm通过主键去删除记录 , 如果主键为空 ， 则会删除该 model 里面的所有记录
	// (如果添加了 gorm.Model 都是软删除 : 设置 delete_at 时间为当前时间 ， 查找数据时 ， 不会去查找 delete_at 有值的)
	var u4 = Teacher{}
	u4.ID = 4
	db.Debug().Delete(&u4)

	// 也可以不通过主键删除
	db.Debug().Where("Standard=?", 7).Delete(Teacher{})

	// 查询被软删除的记录   db.Unscoped().Where().Find()
	var u5 = Teacher{}
	db.Debug().Unscoped().Where("ID=?", 2).Find(&u5)
	fmt.Printf("%#v\n", u5)

	// 如果不想实现软删除, 而是将数据库中的数据直接移除      db.Unscoped().Where().Delete()
	db.Debug().Unscoped().Where("Name=?", "ff").Delete(Teacher{})

}
