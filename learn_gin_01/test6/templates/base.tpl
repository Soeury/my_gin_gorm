<!DOCTYPE html>
<html lang="zh-CN">
<head>

    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>home</title>

    <style>
        * {
            margin: 0;
        }

        .nav {
            height: 50px;
            width: 100%;
            position: fixed;
            top: 0;
            background-color: burlywood;
        }

        .main {
            margin-top: 50px;
        }

        .menu {
            height: 100%;
            width: 20%;
            position: fixed;
            left: 0;
            background-color: cornflowerblue;
        }

        .center {
            text-align: center;
        }
    </style>

</head>
<body>

   <div class="nav"></div>
   <div class="main">
       <div class="menu"></div>
       <div class="content center">
            {{block "content" .}} {{end}} 
       </div>
   </div>

</body>
</html>
