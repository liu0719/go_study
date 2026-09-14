package utools

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var sensitiveword = []string{"操", "傻逼", "妈", "fuck", "shit", "脑残"}

// 自定义屏蔽敏感词函数,init函数时注册到validator内
func ValidDateSensitiveWord(fl validator.FieldLevel) bool {
	
	value := fl.Field().String()
	for _, word := range sensitiveword {
		if strings.Contains(value, word) {
			return false
		}
	}
	return true
}
