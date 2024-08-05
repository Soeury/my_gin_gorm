{{/*  继承根模板  */}}

{{template "base.tpl" .}}

{{/*  重新定义不同的部分  */}}

{{define "content"}}

    <h1>this is in page 'HOME_2'</h1>
    <p>hello , {{.}}</p>

{{end}}