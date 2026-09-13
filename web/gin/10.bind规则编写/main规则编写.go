package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liu0719/go_study/web/gin/utools"

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
	Name  string `json:"name" binding:"required,sensitive" label:"用户名"`
	Email string `json:"email" binding:"required,email" label:"邮箱"`
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
		// 将敏感词函数注册到validator内
		if err := v.RegisterValidation("sensitive", utools.ValidDateSensitiveWord); err != nil {
			panic(err)
		}
		// 向翻译器添加一个名为 "sensitive" 的翻译模板
		// "{0}包含敏感词,换一个试试" 是错误信息的模板
		// {0} 会被自动替换为实际的字段名
		// true 表示这个模板是强制性的
		if err := v.RegisterTranslation("sensitive", trans, func(ut ut.Translator) error {
			return ut.Add("sensitive", "{0}包含敏感词,换一个试试", true)
			// 获取注册的翻译模板，并将字段名 fe.Field() 作为参数传入返回格式化后的错误信
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("sensitive", fe.Field())
			return t
		}); err != nil {
			panic(err)
		}

		// 前端收到的字段名也改为中文
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			// 获取lable字段的值
			label := field.Tag.Get("label")
			if label == "" {
				// 为空就返回默认字段名
				return field.Name
			}
			// 打开下面这一行,所有获取字段名都将是label标签内的中文
			// return label
			// 为了防止对应的json中的key-value中的key也变为中文,这里需要获取json的字段拼接一下
			name := field.Tag.Get("json")
			// fan
			return fmt.Sprintf("%s--%s", name, label)
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
				// errMsg[e.Field()] = e.Translate(trans)
				// 使用通用翻译器翻译错误信息
				msg := e.Translate(trans)
				// 将field分割并map化
				list := strings.Split(msg, "--")
				errMsg[list[0]] = list[1]

			}
			c.JSON(200, gin.H{
				"error": errMsg,
			})
			return
		}

		// 校验成功
		c.String(200, "your name is "+user.Name+",email is "+user.Email)

	})
	r.Run(":80")
}
