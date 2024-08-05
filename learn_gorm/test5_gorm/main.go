package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm 更新

// 建立模型
type Student struct {
	gorm.Model
	Name    string
	Level   int
	Marrige bool
}

func main() {

	// 打开数据库
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/my_2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer db.Close()

	// 建立对应关系
	db.AutoMigrate(&Student{})

	// 插入数据
	u1 := Student{Name: "aa", Level: 3, Marrige: true}
	u2 := Student{Name: "tt", Level: 5, Marrige: false}
	db.Create(&u1)
	db.Create(&u2)

	// 更新
	var p1 Student
	db.First(&p1)

	p1.Name = "pp"
	p1.Level = 9
	db.Debug().Save(&p1) // save 会默认更新所有的字段
	db.Debug().Model(&p1).Update("Name", "pp2")

	// 更新一个字段用 update  , 多个字段用 updates  +  map[string]interface{}
	p2 := map[string]interface{}{
		"Name":    "ii",
		"Level":   7,
		"Marrige": true,
	}

	db.Debug().Model(&p2).Updates(p2)                 // 更新所有字段
	db.Debug().Model(&p2).Select("Level").Update(p2)  // 只更新Level字段
	db.Debug().Model(&p2).Omit("Marrige").Updates(p2) // 只更新排除 Marrige 字段外的所有字段

	// 使用 .RowsAffected 计算影响的记录总数 , 返回一个记录数字 , Model(struct_name{})表示选中所有记录
	rowsnum := db.Model(Student{}).Update("Name", "ss").RowsAffected
	fmt.Printf("rowsaffected num is : %d", rowsnum)

	// 将表中的所有 level 在原来的基础上进行相同的变化
	db.Debug().Model(Student{}).Updates(map[string]interface{}{"Level": gorm.Expr("level * ? + ?", 2, 1)})

}
