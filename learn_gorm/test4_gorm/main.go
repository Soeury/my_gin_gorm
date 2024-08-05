package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm 实现数据查询

// 2. 创建一个模型
type Animal struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {

	// 1. 连接数据库
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/my_2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("found err in gorm.Open : %v\n", err)
		return
	}
	defer db.Close()

	// 3. 结构体与数据库表对应
	db.AutoMigrate(&Animal{})

	// 4. 创建
	a1 := Animal{Name: "alice", Age: 17}
	a2 := Animal{Name: "robin", Age: 21}
	db.Create(&a1)
	db.Create(&a2)

	//  5.查询  :  db.First(&name , num) 只有在主键为int 类型时才能使用
	// 下面两行代码改成  p1 := new(Animal)  db.First(p1) 也可以
	var p1 Animal
	db.First(&p1) // 获取第一条记录

	var p2 Animal
	db.Take(&p2) // 随机获取一条记录

	var p3 Animal
	db.Last(&p3) // 获取最后一条记录

	var p4 []Animal // 获取所有记录
	db.Find(&p4)

	var p5 Animal // 获取指定某一条记录
	db.First(&p5, 2)

	var p6 Animal
	db.FirstOrInit(&p6, Animal{Name: "saray"}) // 存在就返回该条记录 , 不存在就创建一条新的记录然后返回
	fmt.Printf("%#v\n", p6)

	var p7 Animal
	var p8 Animal
	db.Attrs(Animal{Age: 33}).FirstOrInit(&p7, Animal{Name: "chen"})   // Attrs里面的内容在创建一条新的记录时生效
	db.Assign(Animal{Age: 44}).FirstOrInit(&p8, Animal{Name: "zhang"}) // 不管返回什么, Assign里面的内容都会生效

	// struct or map 查询
	// not 条件
	// or 条件
	// 内连条件
	// 立即执行方法
	// 范围
	// 额外查询选项
	// firstorcreate
	// attrs
	// assign
	// 子查询
	// 选择字段
	// 排序
	// 数量
	// 偏移
	// 总数
	// group & having
	// 连接
	// pluck
	// 扫描
	// 链式操作
	// 多个立即执行方法

}
