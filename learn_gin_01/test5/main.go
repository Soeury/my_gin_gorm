package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//  模板引擎  :    自定义函数    +    模板嵌套

//  1. 自定义函数 :   首先需要自己定义一个函数 ， 然后需要在模板解析之前告诉模板有一个自定义的函数
//     t.Funcs(template.FuncMap{
//         "name" : name ,
//    })
//  注意: 自定义函数在模板文件内使用时也放在  {{ 函数名 }} 里面

//  2. 模板嵌套 : 主要注意多个模板解析的顺序 ， 父模板放在前面 ， 子模板放在后面

func sayhello(w http.ResponseWriter, r *http.Request) {

	// 自定义函数 : 要么只有一个返回值 ， 有么有两个返回值 ， 但是第二个返回值必须是 error
	good := func(name string) (string, error) {
		return name + " is handsome", nil
	}

	// 定义模板
	t := template.New("hello.tpl")

	// 在模板解析之前，告诉模板引擎 ， 我们多了一个自定义的函数 good
	t.Funcs(template.FuncMap{
		"good": good,
	})

	//  模板解析
	//  注意这里用的是模板解析    template.New("模板文件的名字").ParseFiles("路径")    的方法
	_, err := t.ParseFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test5\\hello.tpl")
	if err != nil {
		fmt.Printf("found err in template.ParseFiles : %v\n", err)
		return
	}

	name := "chen"

	//  模板渲染
	//  所以说 ， 这里的 name 对象是传给了模板函数对吗 ?
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("found err in t.Execute : %v\n", err)
		return
	}
}

func saydouble(w http.ResponseWriter, r *http.Request) {

	// 定义模板 ： 定义两个父子模板  father.tpl   son.tpl

	// 解析模板 : 多个模板文件注意书写的顺序 ， 父模板写在前面 ， 子模板写在后面
	t, err := template.ParseFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test5\\father.tpl", "D:\\M_GO\\GO_gin\\learn_gin_01\\test5\\son.tpl")
	if err != nil {
		fmt.Printf("found err in template.ParseFiles : %v\n", err)
		return
	}

	name := "Chen"

	// 嵌套模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("found err in t.Execute : %v\n", err)
	}
}

func main() {

	http.HandleFunc("/hello", sayhello)
	http.HandleFunc("/double", saydouble)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("found error in http.ListenAndServe : %v\n", err)
		return
	}
}
