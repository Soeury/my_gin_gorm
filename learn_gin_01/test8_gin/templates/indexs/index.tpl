{{define "indexs/index.tpl"}}
    <!DOCTYPE html>
    <html lang="zh-CN">
    <head>

        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <link rel="stylesheet" href="/xxx/css/index.css">
        <title>hello</title>

    </head>
    <body>

        <p> hello  , {{ .name | change }} </p>


        <script src="/xxx/js/index.js"></script>
    </body>
    </html>
{{end}}