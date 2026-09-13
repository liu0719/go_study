package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("json", func(c *gin.Context) {
		// 一定要大写，不然获取不到
		type User struct {
			// bind内置规则
			// 1.检验为空,最大最小长度
			Name string `json:"name" binding:"required,min=5,max=10"`
			// 2.最大最小范围
			Age int `json:"age" binding:"gte=0,lte=150"`
			// 忽略字段
			Nothing string `json:"-"`
			Pwd     string `json:"pwd"`
			// 3.是否和其他字段相同
			RePwd string `json:"repwd" binding:"eqfield=Pwd"`
			// 4.枚举，要符合列出的其中一个
			Color string `json:"color" binding:"oneof=red green blue"`
			// 5.字符串开头结尾，包含和不包含
			Address string `json:"address" binding:"endswith=bb"`
			// 6.ip和和数组
			Ip []string `json:"ip" binding:"dive,ip"`
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
	r.Run(":80")

}
