package res

import "github.com/gin-gonic/gin"

// json封装，目的是为了实现业务更方便，直接调用方法，清晰一点

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var CodeMap = map[int]string{
	1001: "权限错误",
	1002: "角色错误",
}

// 封装响应方法，response是真正干活的，在response上层调用不同的方法，清晰点
func response(c *gin.Context, code int, msg string, data any) {
	c.JSON(200, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 成功响应全部
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

//
func FailWithCode(c *gin.Context, code int) {
	// 用code码在codeMap内查询，查到了返回错误就行
	msg, ok := CodeMap[code]
	// 没查到返回默认的msg
	if !ok {
		msg = "服务错误"
	}
	response(c, code, msg, nil)
}
func FailWithMsg(c *gin.Context, msg string) {
	response(c, 1001, msg, nil)
}
