package main

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"

	//gin的绑定机制
	"github.com/gin-gonic/gin/binding"
	// 中文语言包,提供中文环境
	"github.com/go-playground/locales/zh"
	//通用翻译器,将验证错误消息翻译为不同语言
	ut "github.com/go-playground/universal-translator"
	//验证器
	"github.com/go-playground/validator/v10"
	// 中文翻译,有预定义的中文错误消息
	// 将英文默认错误翻译为中文
	zh_translate "github.com/go-playground/validator/v10/translations/zh"
)

type User struct {
	Ip string `json:"ip" binding:"fip=345" label:"IP地址"`
}

// 声明全局变量，来当翻译器
var trans ut.Translator

// init作用是：
// 1.创建一个中文翻译器
// 2.将翻译器注册到 Gin 的验证引擎中
// 3.这样后续的所有验证错误都会自动使用中文提示
func init() {
	// 创建翻译器
	uni := ut.New(zh.New())            //使用中文环境创建通用翻译器
	trans, _ = uni.GetTranslator("zh") // 获取中文翻译器

	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 注册默认中文翻译
		_ = zh_translate.RegisterDefaultTranslations(v, trans)

		// 前端收到的字段名也改为中文
		// v.RegisterTagNameFunc(func(field reflect.StructField) string {
		// 	// 获取lable字段的值
		// 	label := field.Tag.Get("label")
		// 	if label == "" {
		// 		// 为空就返回默认字段名
		// 		return field.Name
		// 	}
		// 	// 打开下面这一行,所有获取字段名都将是label标签内的中文
		// 	// return label
		// 	// 为了防止对应的json中的key-value中的key也变为中文,这里需要获取json的字段拼接一下
		// 	name := field.Tag.Get("json")
		// 	// fan
		// 	return fmt.Sprintf("%s--%s", name, label)
		// })

		v.RegisterValidation("fip", func(fl validator.FieldLevel) bool {
			// 字段值
			fmt.Println("fl.Field()", fl.Field())
			// 处理后的字段名称
			fmt.Println("fl.FieldName()", fl.FieldName())
			// 原本的字段名
			fmt.Println("fl.StructFieldName()", fl.StructFieldName())
			// 父级
			fmt.Println("fl.Parent()", fl.Parent())
			// 祖先级,顶层
			fmt.Println("fl.Top()", fl.Top())
			// 参数
			fmt.Println("fl.Param()", fl.Param())

			ip, ok := fl.Field().Interface().(string)
			if ok && ip != "" {
				if ipObj := net.ParseIP(ip); ipObj != nil {
					return true
				}

			}
			// 设置为true,不填就不会去校验
			return true
		})
	}

}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			// 类型断言，进行验证失败类型
			errs, ok := err.(validator.ValidationErrors)
			// 不是校验错误就直接返回给前端，是其他错误
			if !ok {

				c.JSON(200, gin.H{
					"error": err.Error(), // 对于JSON解析错误等，直接返回原始错误
				})
				return
			}
			// 校验失败
			errMsg := make(map[string]string)
			for _, e := range errs {
				// 为了显示显示中文信息,在上面通过v.RegisterTagNameFunc()方法将所有获取字段改为了中文,下面这个就不能用了
				errMsg[e.Field()] = e.Translate(trans)
				// 使用通用翻译器翻译错误信息
				// msg := e.Translate(trans)
				// // 将field分割并map化
				// list := strings.Split(msg, "--")
				// errMsg[list[0]] = list[1]

			}
			c.JSON(400, gin.H{
				"error": errMsg,
			})
			return
		}
		fmt.Println(user.Ip)
		// 校验成功
		c.JSON(200, gin.H{
			"ip":   user.Ip,
			"code": "成功",
		})

	})
	r.Run(":80")
}
