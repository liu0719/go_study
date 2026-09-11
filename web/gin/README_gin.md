# Gin
[Gin文档](https://gin-gonic.com/zh-cn/docs)
## 1.安装
```bash
go mod init webdemo
go get github.com/gin-gonic/gin
```
## 2.Gin中的请求方式
|方法|典型 REST 用途|
|---|---|
|**GET**|获取资源|
|**POST**|创建新资源|
|**PUT**|替换现有资源|
|**PATCH**|部分更新现有资源|
|**DELETE**|删除资源|
|**HEAD**|与 GET 相同但不返回响应体|
|**OPTIONS**|描述通信选项|
## 3.初始化，挂载路由,启动
```go
package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/liu0719/go_study/web/gin/res"
)
func Index(c *gin.Context) {
    fmt.Println(c.Request){
    // 3.响应json,gin.H是一个map[string]any类型
    // 响应码一般都是200，
    c.JSON(200, gin.H{
        // 这个code是业务码，判断业务层面是否正确
        "Code": 200,
        "Msg":  "后端数据",
        "Data": gin.H{},
    })
}

func main() {
    // 0.隐藏终端的gin日志
    gin.SetMode(gin.ReleaseMode)
    // 1.初始化，会返回一个gind的引擎，engine
    r := gin.Default() 
    // 2.挂载路由
    r.GET("/", Index)
    
    //调用json封装的方法，业务处理时可读性高也清晰
    r.GET("/login", func(c *gin.Context) {
        res.OkWithMsg(c, "登陆成功")
    })
    r.GET("/users", func(c *gin.Context) {
        res.OkwithData(c, map[string]any{
            "name": "张三",
        })
    })
    r.POST("/users", func(c *gin.Context) {

        res.FailWithMsg(c, "验证失败")

    })

    r.Run(":80")

}
```

## 4.json封装
```go
// json封装，把业务封装成函数，目的是为了实现业务更方便，直接调用方法，清晰一点

//json传输格式
type Response struct {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
    Data any    `json:"data"`
}
//codemap,查询code对应的msg,用于返回前端
var CodeMap = map[int]string{

    1001: "权限错误",

    1002: "角色错误",

}

// 封装响应方法，response是真正干活的，在response上层调用不同的方法，清晰点
func response(c *gin.Context, code int, msg string, data any) {
    c.JSON(200, Response{
        Code: code,
        Msg:  msg,
        Data: data,
    })
}

// 成功响应全部，传msg,data

func Ok(c *gin.Context, msg string, data any) {

    response(c, 0, msg, data)

}

// 只响应msg,其他参数采用默认
func OkWithMsg(c *gin.Context, msg string) {
    // 这里的data位置不能传nil,否则传到前端就是null，会报空指针
    response(c, 0, msg, gin.H{})
} 

// 只响应data
func OkwithData(c *gin.Context, data any) {
    response(c, 0, "成功", data)
}

// 失败响应全部
func Fail(c *gin.Context, Code int, msg string, data any) {
    response(c, 0, msg, data)
}

  

//去map内查状态，返回msg
func FailWithCode(c *gin.Context, code int) {

    // 用code码在codeMap内查询，查到了返回错误就
    msg, ok := CodeMap[code]
    //没查到返回默认的msg
    if !ok {
        msg = "服务错误"
    }
    response(c, code, msg, nil)
}
//自定义msg
func FailWithMsg(c *gin.Context, msg string) {
    response(c, 1001, msg, nil)

}
```
## 5.响应html
```go
func main() {
  r := gin.Default()
  //预加载html文件
  router.LoadHTMLGlob("templates/*")
  //这个方法只能预加载单个文件
  router.GET("/index", func(c *gin.Context) {
	//第三个参数传给前端，前端可以采用`{{.title}}`的方式来调用
    c.HTML(http.StatusOK, "index.html", gin.H{
      "title": "我是标题",
    })
  })
  router.Run(":80")
}
```
## 6.部署方式
1. 前后端分别单独部署
		前端单独启动一个系统和端口负责页面展示，后端的系统和端口负责数据
2. 前端先弄好，后端统一部署
		前端当作静态文件在后端中，由后端统一部署。
## 7.响应文件
用于浏览器直接请求找个接口唤起下载
1. 在响应头内设置内容类型，再设置文件名，就能直接唤起浏览器下载
2. 只能是get请求
```go
// 在消息头设置内容类型，表示是文件流，唤起浏览器下载
c.Header("Content-Type", "application/octet-stream")
// 再设置文件名
c.Header("Content-Disposition", "attachment; filename=photo.jpg")
c.File("static/photo.jpg")
```
大部分情况下都是用第二种，前端页面请求后端接口，唤起浏览器下载
```html
<a href="文件路径" download="下载后的文件名">点我下载</a>
```
最好的处理方式其实是
前端请求，后端返回一个临时下载地址，再由前端构造a标签，再请求后端接口下载
## 8.静态文件
```go
// 静态文件路径代替，第一个参数是url路径(别名)，第二个是文件路径
    r.Static("st", "static")
    // 这个只能用于单个文件
    r.StaticFile("hello", "static/hello.txt")
```
静态文件的路径不能再被路由使用，否则会报错
## 9.参数
> 查询参数
`?key=xxx&name=xxx`,这种就是查询参数
查询参数不是get方法的专属

动态参数
表单参数
## 常用功能

- 路由注册：GET、POST、PUT、DELETE、PATCH
- 参数绑定：Query、Path、Body
- JSON 响应：c.JSON()
- 文件上传：c.PostForm()
- 中间件：日志、鉴权、限流、跨域等
- 


