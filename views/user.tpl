<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>網站 {{.Website}}</h1>
    {{ if .user }}
    用戶名: {{ .user.Username }}
    {{ else }}
    查無此用戶
    {{ end }}
</body>
</html>