<!DOCTYPE html>
<html lang="zh-CN">
<head>

    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>hello</title>

</head>
<body>

    {{/*   我是 template 里面的注释     */}}




    {{/*  自定义变量    $变量名 := 值   */}}
    {{ $a := 100 }}
    {{ $age := .s1.Age }}




    {{ range  $k , $v := .n1 }}
        <p> {{$k}} - {{$v}} </p>
    {{else}}
        <p> slice is empty </p>
    {{end}}
    <br>




    {{/*  index 用法  :    在某对象中取出指定下标的值   */}}
    {{index  .n1  2}}
    <br>




    {{/*  if - else 语句中比较函数的用法   先写函数 , 再写两个值   */}}
    {{if lt .s1.Age 18}}
    <p>study</p>
    {{else}}
    <p>work</p>
    {{end}}
    <br>
   



    <p> Demo - s1 </p>
    <p> Name : {{ .s1.Name }} </p>
    <p> Age :{{ .s1.Age }} </p>
    <p> Height :{{ .s1.Height }} </p>
    <br>




    {{/* with作用域  :    通过 with end 语句 , 创建一个作用域 , 可以简化   [ .对象.字段 ]   的写法为    [ .字段 ]    */}}

    {{with .m1}}
    <p> Map - m1 </p>
    <p> Standard :{{ .Standard }} </p>
    <p> Salary :{{ .Salary }} </p>
    <p> Age : {{ .Age }} </p>
    {{end}}
    



</body>
</html>

   