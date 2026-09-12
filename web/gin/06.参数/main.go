package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	// 1.路径参数 /users?id=123
	// 在浏览器内输入`127.0.0.1?name=张三&age=100&key=234&key=789`
	r.GET("/pathData", func(c *gin.Context) {
		// 拿过来默认为string
		name := c.Query("name")
		// 设置为默认值
		age := c.DefaultQuery("age", "-1")
		// 获取参数数组
		keyList := c.QueryArray("key")

		fmt.Println("name:", name)
		fmt.Println("age:", age)
		fmt.Println("keyList:", keyList)

	})

	// 2.动态参数  /users/:name/:age/*action
	r.GET("/users/:id/*action", func(c *gin.Context) {
		id := c.Param("id")
		age := c.Param("age")

		action := c.Param("action") //这里获取的参数前会有`/`
		fmt.Println(id + "，" + age + "岁，正在" + action)
	})
	//3.表单参数  一般专指form表单
	r.POST("/users", func(c *gin.Context) {
		// 获取表单参数,PostForm分不清前端传没传
		name := c.PostForm("name")
		// GetPostForm会返回bool值，看看传了参数没有,可能用于选择更新
		age, ok := c.GetPostForm("age")
		// 这里处理一下没传的情况,设置一个默认值
		if !ok {
			age = "-1"
		}
		fmt.Println(name, age, ok)

		// 传统的文件上传
		/*
			fileHeader, err := c.FormFile("file")
			if err != nil {
				panic(err)
			}
			fmt.Println(fileHeader.Filename) //文件名
			fmt.Println(fileHeader.Size)     //文件大小

			file, err := fileHeader.Open()
			if err != nil {
				panic(err)
			}
			byteData, _ := io.ReadAll(file)
			// 可以写路径，开头不要加/，相对于项目根目录
			err = os.WriteFile("static/hello.jpg", byteData, 0666)
			fmt.Println(err)
		*/

		// gin的文件上传
		/*
			err = c.SaveUploadedFile(fileHeader, "static/"+fileHeader.Filename)
			fmt.Println(err)
		*/

		// 多文件上传
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)
		}
		// form.File在这里是map[string][]*filerHeader,这个map的value是一个数组，数组内存放着真正的Header
		for _, headers := range form.File {
			// 再次循环这个数组就能得到真正的Header
			for _, header := range headers {
				c.SaveUploadedFile(header, "static/"+header.Filename)
			}
		}

	})

	r.Run(":80")
}
