<!DOCTYPE html>
<html lang="zh-CN">
<head>

    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>hello</title>

</head>
<body>

    <h1> 测试嵌套 template 语法   </h1>
    <br>
    <br>

    {{/*  嵌套另外一个单独的 son.tpl 模板文件 */}}
    {{ template "son.tpl" }}
    <br>
    <br>

    {{/*  嵌套内置 define 定义的模板文件   */}}
    {{ template "def.tpl"}}
    <br>
    <br>

    <div>hello {{ . }}</div>
   

</body>
</html>



{{/* 通过define定义一个模板  */}}
{{ define "def.tpl" }}

    <ul>
        <li>def1</li>
        <li>def2</li>
        <li>def3</li>
        <li>def4</li>
    </ul>

{{end}}