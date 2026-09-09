# Gin

## 安装
```bash
go mod init webdemo
go get github.com/gin-gonic/gin
```
## 初始化，挂载路由,启动
```go
// 隐藏终端的gin日志
gin.SetMode(gin.ReleaseMode)

// 1.初始化，会返回一个gind的引擎，engine

    r := gin.Default()
// 2.挂载路由
    r.GET("/", Index)

    r.Run(":80")
```

## json封装

## 常用功能

- 路由注册：GET、POST、PUT、DELETE、PATCH
- 参数绑定：Query、Path、Body
- JSON 响应：c.JSON()
- 文件上传：c.PostForm()
- 中间件：日志、鉴权、限流、跨域等


