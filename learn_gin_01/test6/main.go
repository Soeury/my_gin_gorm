package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//  模板继承 :  当需要使用到多个模板并且模板之间相似度非常高，可以考虑模板继承 ，建立一个根模板 ， 剩余内容套用即可
//  步骤 :

//     1-  建立一个根模板 ，将要使用的多个模板之间的'相同部分'保留
//     2-  根模板内的不同部分用   {{block "content" .}} {{end}}   代替
//     3-  注意这里的 "content" 取名 ,这是对区域的命名 , 在我们的子模板内用来区分某块区域应该填上什么内容
//     4-  接下来是对子模板的操作了
//     5-  子模板内部，需要有两个部分的操作
//         1- 继承根模板 :   {{template  "根模板的名字.tpl" .}}  这里的 . 用来继承传进来的数据对象
//         2- 重新定义不同的部分 :  {{define "区域名字"}} 需要填上的内容 {{end}}
//     6-  采用模板继承的方式，在解析模板的时候，需要写上多个模板 ， 并且把根模板写在前面 ， 子模板写在后面
//     7-  采用模板继承的方式，在渲染的时候要改成  t.ExecuteTemplate( w  ,  子模板名字.tpl  ,  要传入的数据)

func index(w http.ResponseWriter, r *http.Request) {

	// 定义模板 - 各自定义的方式
	// 模板渲染
	t, err := template.ParseFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test6\\index_1.tpl")
	if err != nil {
		fmt.Printf("found error in tempalte.ParseFiles : %v\n", err)
		return
	}

	name := "Index"

	// 模板继承
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("found error in t.Execute : %v\n", err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {

	// 定义模板 - 各自定义的方式
	// 模板渲染
	t, err := template.ParseFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test6\\home_1.tpl")
	if err != nil {
		fmt.Printf("found error in tempalte.ParseFiles : %v\n", err)
		return
	}

	name := "Home"

	// 模板继承
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("found error in t.Execute : %v\n", err)
		return
	}
}

func index_2(w http.ResponseWriter, r *http.Request) {

	// 定义模板 - 模板继承的方式

	// 解析模板 :  根模板写在前面   子模板写在后面
	t, err := template.ParseFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test6\\templates\\base.tpl", "D:\\M_GO\\GO_gin\\learn_gin_01\\test6\\templates\\index_2.tpl")
	if err != nil {
		fmt.Printf("found error in tempalte.ParseFiles : %v\n", err)
		return
	}

	name := "index_2"

	// 渲染模板  :  采用继承模板的方式 , 渲染模板时 需要使用   t.ExecuteTempalte(  w , "name.tpl" , 传入的对象)
	err = t.ExecuteTemplate(w, "index_2.tpl", name)
	if err != nil {
		fmt.Printf("found error in t.ExecuteTemplate : %v\n", err)
		return
	}
}

func home_2(w http.ResponseWriter, r *http.Request) {

	// 定义模板 - 模板继承的方式

	// 解析模板
	t, err := template.ParseFiles("D:\\M_GO\\GO_gin\\learn_gin_01\\test6\\templates\\base.tpl", "D:\\M_GO\\GO_gin\\learn_gin_01\\test6\\templates\\home_2.tpl")
	if err != nil {
		fmt.Printf("found error in template.ParseFiles : %v\n", err)
		return
	}

	name := "home_2"

	// 渲染模板
	err = t.ExecuteTemplate(w, "home_2.tpl", name)
	if err != nil {
		fmt.Printf("found error in t.ExecuteTemplate : %v\n", err)
		return
	}
}

func main() {

	// 这里 , 每一个 http.HandleFunc("/name" , name) 对应一个函数 , 表示一次请求响应
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index_2", index_2)
	http.HandleFunc("/home_2", home_2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("found error in http.ListenAndServe : %v\n", err)
		return
	}
}
