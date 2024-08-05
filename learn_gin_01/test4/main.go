package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//   ---- 2. net/http模板引擎 :   了解如何传入多个参数   +   常用模板语法介绍

type Demo struct {
	Name   string
	Age    int
	Height int
}

func sayhello(w http.ResponseWriter, r *http.Request) {

	// 解析模板
	t, err := template.ParseFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test4\\hello2.tpl")
	if err != nil {
		fmt.Printf("found error in template.ParseFiles : %v\n", err)
		return
	}

	// 一个结构体对象
	s1 := Demo{"chen", 18, 167}

	// 一个 map 对象
	m1 := map[string]int{
		"Standard": 1,
		"Salary":   18000,
		"Age":      27,
	}

	// 一个切片对象
	n1 := []string{"apple", "pine", "melon"}

	//  渲染模板   放入单个参数     err = t.Execute(w, 对象)
	//  模板渲染:  放入单个参数     {{ .字段名 }}

	//  渲染模板 放入多个参数     err = t.Execute(w , map[string]interface{}{ "key" : 对象  ,  "key" : 对象 })
	//  模板渲染:  放入单个参数   {{ .对象.字段名 }}

	err = t.Execute(w, map[string]interface{}{
		"m1": m1,
		"s1": s1,
		"n1": n1,
	})

	if err != nil {
		fmt.Printf("found error in t.Execute : %v\n", err)
		return
	}

}

func main() {

	http.HandleFunc("/hello", sayhello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("found error in http,ListenAndServe : %v\n", err)
		return
	}
}
