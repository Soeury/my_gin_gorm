<!DOCTYPE html>
<html lang="zh-CN">
<head>

    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>index</title>

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
            width: 20%;
            height: 100%;
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
           <h1>this is in page 'index'</h1>
           <p> hello , {{ . }} </p>
       </div>
   </div>

</body>
</html>

