package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//   ---- 1. net/http模板初识(模板定义 ， 解析 ， 渲染)

//   go 语言的模板引擎

//   1. go 语言内置了 html/template 包作为模板引擎
//   2. 模板文件以  .tpl 为后缀 ， 必须使用 utf-8 编码
//   3. 模板文件中使用 {{ . }} 包裹需要传送的数据 ，  其中  .  表示我们传送给模板的数据
//   4. 除了 {{ }} 包裹的内容外 ， 其他内容都不做修改原样输出

//   go 语言模板引擎的使用
//   1. 定义模板
//      创建一个 .tpl 为后缀的文件 ，

//   2. 解析模板 - 两种方式:
//          t , error := template.ParseFiles("路径")
//          t , error := template.New("f").ParseFiles("路径")

//   3. 渲染模板
//  err  =  解析出来的模板.Execute( 要传给的对象 , 传给模板中 {{.}} 位置上的数据 )

func sayhello(w http.ResponseWriter, r *http.Request) {

	//  解析模板

	t, err := template.ParseFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test3\\hello.tpl")
	if err != nil {
		fmt.Printf("found err in template.ParseFiles : %v\n", err)
		return
	}

	//  渲染模板

	err = t.Execute(w, "My_String")
	if err != nil {
		fmt.Printf("found err in t.Execute : %v\n", err)
		return
	}

}

func main() {

	http.HandleFunc("/hello", sayhello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("found error in http.ListenAndServe : %v\n", err)
		return
	}

}
