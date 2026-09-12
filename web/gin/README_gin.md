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
## 7.下载文件
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
## 9.参数，文件上传
1. 查询参数
`?key=xxx&name=xxx&name=yyy`,这种就是查询参数
查询参数不是get方法的专属
```go
// 拿过来默认为string
name := c.Query("name")
// 设置为默认值
age := c.DefaultQuery("age", "-1")
// 获取参数数组
keyList := c.QueryArray("key")

fmt.Println("name:", name)
fmt.Println("age:", age)
fmt.Println("keyList:", keyList)
```
输出为
```json
name: hello
age: 100
keyList: [234 789]
```
2. 动态参数
	用户个人信息页面，他的路径：
	```go
	/users?id=123   //查询参数格式
	/users/123      //动态参数
	```
3. 表单参数
```go
// 获取表单参数,PostForm分不清前端传没传
        name := c.PostForm("name")
        // GetPostForm会返回bool值，看看传了参数没有,可能用于选择更新
        age, ok := c.GetPostForm("age")
        // 这里处理一下没传的情况,设置一个默认值
        if !ok {
            age = "-1"
        }
        fmt.Println(name, age, ok)

```
4. 文件上传
	传统的文件上传处理
```go
fileHeader, err := c.FormFile("file")
        if err != nil {
            panic(err)
        }
        fmt.Println(fileHeader.Filename) //文件名
        fmt.Println(fileHeader.Size)     //文件大小

        file, err := fileHeader.Open()
        if err != nil {
            panic(err)
        }
        byteData, _ := io.ReadAll(file)
        // 可以写路径，开头不要加/，相对于项目根目录
        err = os.WriteFile("static/hello.jpg", byteData, 0666)
        fmt.Println(err)
```
	gin的文件上传处理
```go
err = c.SaveUploadedFile(fileHeader, "static/"+fileHeader.Filename)
        fmt.Println(err)
```
5. 多文件上传
```go
form, err := c.MultipartForm()
        if err != nil {
            fmt.Println(err)
        }
        // form.File在这里是map[string][]*filerHeader
        // 这个map的value是一个数组，数组内存放着真正的Header
        for _, headers := range form.File {
            // 再次循环这个数组就能得到真正的Header
            for _, header := range headers {
                c.SaveUploadedFile(header, "static/"+header.Filename)
            }
        }
```
## 10.数据传输方式的原始内容
不同请求体对应的原始内容
### 解决Body阅后即焚
```go
// 这个body本质上是从网络中拿的，阅后即焚，再读就读不到了
byteData, _ := io.ReadAll(c.Request.Body)
fmt.Println(string(byteData))
// 解决Body阅后即焚的问题
// 只有将body的内容加载到到内存,再用bytes的NewReader读取内存中的bytedata
// 再用io.NopCloser()将*reader包装成关闭空操作的同类型，再赋值给Body，就能实现反复读取
c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
```
### form-data格式
对应的消息头
```go
map[Accept:[*/*] 
Accept-Encoding:[gzip, deflate, br]
Connection:[keep-alive]
Content-Length:[269]
Content-Type:[multipart/form-data; boundary=--------------------------378516665593888438278638]
User-Agent:[Apifox/1.0.0 (https://apifox.com)]]
```
消息体
```go
----------------------------378516665593888438278638
Content-Disposition: form-data; name="name"

张三
----------------------------378516665593888438278638
Content-Disposition: form-data; name="age"

11
----------------------------378516665593888438278638--
```

对应的分隔符就是`----------------------------378516665593888438278638`来分隔开每个form-data的参数。
想要提取就需要按字符串分割，再按`Content-Disposition:`切割，就可以拿到想要的参数
### x-www-form-url格式
消息头中只有指定类型不同
```go
Content-Type:[application/x-www-form-urlencoded]
```
消息体
```go

name=%E5%BC%A0%E4%B8%89&age=11
```
获取参数
```go
//调用url包内的ParseQuery()
values,err:=url.ParseQuery("?")
        if err!=nil{
            fmt.Println(err)
        }
        // 得到想要的参数
        for k, v := range values {
            fmt.Println(k,v)
        }
```
### json格式
json格式的消息头中只有指定的传输类型不同
```go
Content-Type:[application/json]
```
消息体
```json
{
    "name":"张三",
    "age":34,
    "friend":["Tom","Jack","LiMing"],
    "family":{
        "dad":"liu",
        "mom":"111",
        "brother":"李四",
    }
}
```
## 11.bind参数绑定
#### 查询参数
`/user?id=xxx`
```go
type User struct {
// 一定要大写，不然获取不到
	Name string `form:"name"`
	Age  int    `form:"age"`
}
var user User
// 查询参数Query也会完成数据的校验
err := c.ShouldBindQuery(&user)
fmt.Println(user, err)
```
### 动态参数
`user/:id/*action`
```go
type User struct {
// 一定要大写，不然获取不到
	Name   string `uri:"name"`
	Id     int    `uri:"id"`
	Action string `uri:"action"`
}

var user User
// 动态参数用uri
err := c.ShouldBindUri(&user)
fmt.Println(user, err)
```

### 表单参数
```go
type User struct {
// 一定要大写，不然获取不到
	Name string `form:"name"`
	Age  int    `form:"age"`
}

var user User
// 也会完成数据的校验
err := c.ShouldBind(&user)
fmt.Println(user, err)
```
> 注意，不能解析x-www-form-urlencoded格式
### json参数
```go
type User struct {
// 一定要大写，不然获取不到
	Name    string `json:"name"`
	Age     int    `json:"age"`
	address string `json:"address"`
}

var user User
// 也会完成数据的校验
err := c.ShouldBindJSON(&user)

fmt.Println("json数据为：", user, err)
```
### header参数
```go
// 一定要大写，不然获取不到
        type Header struct {
        // 这里注意一定要和前端发来的数据名一样，有的中间有-
            UserAgent string `header:"User-Agent"`
            ContentType string `header:"Content-Type"`
        }
        var header Header
        // 也会完成数据的校验
        err := c.ShouldBindHeader(&header)

        fmt.Println("json数据为：", header, err)
```
## 12.binding内置规则
```go
r.POST("json", func(c *gin.Context) {
	// 一定要大写，不然获取不到
	type User struct {

		// bind内置规则
		// 1.检验为空,和json的tag相同在后面加上bind的指定字段，来限制获取数据
		Name    string `json:"name" binding:"required"`

		Age     int    `json:"age"`

		address string `json:"address"`

	}

	var user User

	// 也会完成数据的校验

	err := c.ShouldBindJSON(&user)

	if err != nil {

		c.String(200, "出现错误：", err.Error())

		return

	}

	c.JSON(200, user)

})
```
> 常用bind字段
```go

```

## 常用功能

- 路由注册：GET、POST、PUT、DELETE、PATCH
- 参数绑定：Query、Path、Body
- JSON 响应：c.JSON()
- 文件上传：c.PostForm()
- 中间件：日志、鉴权、限流、跨域等
- 


