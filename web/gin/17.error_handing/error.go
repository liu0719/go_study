package main

import (
	"fmt"
	"net/http"
)

// 1.和前端约定好错误码，用数字表示状态和前端约定好
// 不用string或者msg，如果用字符串匹配如果修改会瘫痪
const (
	CodeOK = 0 // 成功。code=0 表示成功是 Go / 前端社区的常见约定

	// 1xxx —— 客户端 / 参数问题：用户的锅，改一下请求就能好
	CodeParamInvalid  = 1001 // 参数不合法
	CodeUserNotFound  = 1002 // 查无此人
	CodeAlreadyExists = 1003 // 重复创建

	// 2xxx —— 认证与授权：、（和 16 章 JWT 那一套接得上）
	CodeNotLogin     = 2001 //没登录
	CodeTokenExpired = 2002 //token过期
	CodeForbidden    = 2003 //没权限

	// 3xxx —— 业务规则：系统没坏，是业务不允许
	CodeBalanceNotEnough = 3001 //业务不允许
	CodeDivByZero        = 3002 // 除以零
	CodeNameTaken        = 3003 //名称占用

	// 4xxx —— 资源不存在（HTTP 404 那一类）
	CodeNotFound = 4001 //资源不存在

	// 5xxx —— 系统错误：服务器的锅
	CodeInternal = 5000 //网络错误
	CodeDBDown   = 5001 //数据库错误
)

// 2. 自己定义的错误类型
// 传给前端的错误信息，进入日志的错误信息，都需要靠这个错误类型来传递和规范
type AppError struct {
	Code       int    //传给前端的状态码，前端用它来做分支，在200，404的基础上再自己设置一层
	HTTPStatus int    // HTTP 状态码，决定浏览器 / 网关 / 监控看到什么，就是标准的200，404，等等
	UserMsg    string // 传给前端的错误信息，要去掉服务器的具体细节，比如哪个端口或者模块，都要藏起来，把前端当成攻击者视角
	Detail     string // 内部细节，只进日志，这部分要包含报错的完整信息，日志是内部查错用的
}

// 把错误封装能被go的errors包内的error.Is()和error.As()识别
// 要想让自定义的错误类型实现errors包的机制，就必须有Error方法
func (e *AppError) Error() string {
	return fmt.Sprintf("Code:[%d],UserMsg:[%s],Detail:[%s]", e.Code, e.UserMsg, e.Detail)
}

// 下面三个函数，是把“错误属于哪个类型”弄成直接调用的方法
// 这样方便，而且不会写错

// 业务错误
func newBizErr(code int, usermsg string) *AppError {
	return &AppError{Code: code, UserMsg: usermsg, HTTPStatus: http.StatusBadRequest}
}

// 权限错误，这个需要手动设置401 | 403
func newAuthErr(code int, usermsg string, status int) *AppError {

}

// 3.
