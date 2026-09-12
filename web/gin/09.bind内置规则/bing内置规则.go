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
			// 1.检验为空
			Name    string `json:"name" binding:"required"`
			Age     int    `json:"age"`
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
	r.Run(":80")

	// 不能为空,且不能没有这个字段
	`bind:"required"`
	
	// 针对字符串的长度
	min 最小长度，如：`bind:"min=5"`
	max 最大长度，如：`bind:"max=10"`
	len 长度，如：`bind:"len=6"`

	// 针对数字大小
	eq 等于 ：`binding:"eq=3"`
	ne 不等于：
	gt 大于：
	gte 大于等于：
	lt 小于：
	lte 小于等于：

	// 针对同级字段

}
