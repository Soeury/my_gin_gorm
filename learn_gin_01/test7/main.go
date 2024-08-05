package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//    模板标识符  {{ }}   的重定义
//    text/template 和 html.template 的区别

//  1.   模板标识符的重定义 : 在解析模板的时候将 template.ParseFiles()  换成  template.New("name.tpl").Delims("左" , "右").ParseFiles("path")
//  2.   html.template 会将所有的 html , css , js 类型的对象 转义为字符串类型 , 目的是为了防止 xss 攻击
//       要实现将指定对象不进行转义成为字符串，我们需要自定义一个函数，将字符串类型的对象转换成 template.HTML类型
//       最后 , 需要还原的对象在模板文件中进行调用即可 , 调用方法  {{  .对象名称 | 函数名称 }}

func sayhello(w http.ResponseWriter, r *http.Request) {

	// 模板定义

	// 模板解析
	t, err := template.New("hello.tpl").Delims("{[", "]}").ParseFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test7\\hello.tpl")
	if err != nil {
		fmt.Printf("found error in template.ParseFiles : %v\n", err)
		return
	}

	// 模板渲染

	name := "chen"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("found error in t.execute : %v\n", err)
		return
	}
}

func xss(w http.ResponseWriter, r *http.Request) {

	// 定义模板

	// 自定义函数
	change := func(str string) template.HTML {
		return template.HTML(str)
	}

	// 解析模板
	t, err := template.New("hello.tpl").Funcs(template.FuncMap{
		"change": change,
	}).ParseFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test7\\hello.tpl")

	if err != nil {
		fmt.Printf("found error in tempalte.ParseFiles : %v\n", err)
		return
	}

	// 渲染模板
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='http://www.baidu.com' target='_blank'>百度搜索</a>"

	err = t.Execute(w, map[string]interface{}{
		"str1": str1,
		"str2": str2,
	})

	if err != nil {
		fmt.Printf("found error in t.execute : %v\n", err)
		return
	}
}

func main() {

	http.HandleFunc("/hello", sayhello)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("found err in http.ListenAndServe : %v\n", err)
		return
	}
}
