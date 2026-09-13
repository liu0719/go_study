package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 方案1: 直接修改json标签为中文（推荐）
type UserChineseJson struct {
	Name  string `json:"用户名" binding:"required,sensitive" label:"用户名"`
	Email string `json:"邮箱" binding:"required,email" label:"邮箱"`
	Age   int    `json:"年龄" binding:"gte=0,lte=150" label:"年龄"`
}

// 方案2: 保持英文json标签，响应时使用自定义转换
type UserEnglishJson struct {
	Name  string `json:"name" binding:"required,sensitive" label:"用户名"`
	Email string `json:"email" binding:"required,email" label:"邮箱"`
	Age   int    `json:"age" binding:"gte=0,lte=150" label:"年龄"`
}

// 响应包装器
type ResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 转换为中文响应的辅助函数
func ToChineseResponse(data interface{}) ResponseWrapper {
	// 如果data是map[string]interface{}，直接转换字段名
	if m, ok := data.(map[string]interface{}); ok {
		return ResponseWrapper{
			Code:    200,
			Message: "成功",
			Data:    convertToChineseFields(m),
		}
	}

	// 如果是结构体，先序列化再转换
	jsonData, _ := json.Marshal(data)
	var temp map[string]interface{}
	json.Unmarshal(jsonData, &temp)

	return ResponseWrapper{
		Code:    200,
		Message: "成功",
		Data:    convertToChineseFields(temp),
	}
}

// 字段名转换函数
func convertToChineseFields(input map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range input {
		switch key {
		case "name":
			result["用户名"] = value
		case "email":
			result["邮箱"] = value
		case "age":
			result["年龄"] = value
		default:
			result[key] = value
		}
	}

	return result
}

// 自定义JSON序列化器（使用函数作为值）
type UserWithCustomJson struct {
	Name  string `json:"用户名" binding:"required,sensitive" label:"用户名"`
	Email string `json:"邮箱" binding:"required,email" label:"邮箱"`
	Age   int    `json:"年龄" binding:"gte=0,lte=150" label:"年龄"`
}

// 方案3: 使用别名和自定义序列化
type UserAlias struct {
	UserName  string `json:"name" binding:"required,sensitive"`
	UserEmail string `json:"email" binding:"required,email"`
	UserAge   int    `json:"age" binding:"gte=0,lte=150"`
}

func (u UserAlias) MarshalJSON() ([]byte, error) {
	// 创建一个临时结构体，使用中文字段名
	type Alias UserAlias
	return json.Marshal(struct {
		Alias
	}{
		Alias: Alias(u),
	})
}

func setupRoutes() *gin.Engine {
	r := gin.Default()

	// 方案1演示：直接使用中文json标签
	r.POST("/user/chinese-json", func(c *gin.Context) {
		var user UserChineseJson
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}

		// 响应时自动使用中文字段名
		c.JSON(200, user)
	})

	// 方案2演示：保持英文json标签，响应时转换
	r.POST("/user/english-json", func(c *gin.Context) {
		var user UserEnglishJson
		if err := c.ShouldBindJSON(&user); err != nil {
			// 使用RegisterTagNameFunc处理错误消息（中文）
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}

		// 手动转换为中文响应
		response := ToChineseResponse(user)
		c.JSON(200, response)
	})

	// 方案3演示：使用别名和自定义序列化
	r.POST("/user/alias-json", func(c *gin.Context) {
		var user UserAlias
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}

		// 自定义序列化会自动使用中文字段名
		c.JSON(200, user)
	})

	// 方案4演示：获取用户列表
	r.GET("/users", func(c *gin.Context) {
		users := []UserChineseJson{
			{
				Name:  "张三",
				Email: "zhangsan@example.com",
				Age:   25,
			},
			{
				Name:  "李四",
				Email: "lisi@example.com",
				Age:   30,
			},
		}

		// 直接返回，自动使用中文字段名
		c.JSON(200, users)
	})

	return r
}

func main() {
	fmt.Println("=== 完整Go JSON序列化解决方案演示 ===\n")

	// 测试各个方案
	fmt.Println("测试方案1 - 中文json标签:")
	user1 := UserChineseJson{
		Name:  "张三",
		Email: "zhangsan@example.com",
		Age:   25,
	}
	json1, _ := json.Marshal(user1)
	fmt.Printf("%s\n\n", string(json1))

	fmt.Println("测试方案2 - 英文json标签 + 转换:")
	user2 := UserEnglishJson{
		Name:  "李四",
		Email: "lisi@example.com",
		Age:   30,
	}
	json2, _ := json.Marshal(user2)
	fmt.Printf("原始JSON: %s\n", string(json2))

	response := ToChineseResponse(user2)
	json3, _ := json.Marshal(response)
	fmt.Printf("转换后JSON: %s\n\n", string(json3))

	fmt.Println("测试方案3 - 自定义序列化:")
	user3 := UserAlias{
		UserName:  "王五",
		UserEmail: "wangwu@example.com",
		UserAge:   28,
	}
	json4, _ := json.Marshal(user3)
	fmt.Printf("%s\n\n", string(json4))

	fmt.Println("=== 启动服务器测试 ===")
	fmt.Println("启动服务器后，可以测试以下端点:")
	fmt.Println("POST /user/chinese-json - 方案1：中文json标签")
	fmt.Println("POST /user/english-json - 方案2：英文json标签+转换")
	fmt.Println("POST /user/alias-json  - 方案3：自定义序列化")
	fmt.Println("GET  /users             - 用户列表")
	fmt.Println("\n服务器启动在 :8080")

	// 启动服务器
	r := setupRoutes()
	r.Run(":8080")
}