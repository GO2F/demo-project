package base

import (
	"github.com/gin-gonic/gin"
)

// Success 操作成功
func Success(data interface{}) (response gin.H) {
	return Result(data, "操作成功", 0, "success", "")
}

// Failed 操作失败
func Failed(errorMsg string, errCode int) (response gin.H) {
	return Result(map[string]string{}, errorMsg, errCode, "alert", "")
}

// Jump 跳转
func Jump(errorMsg string, errCode int, url string) (response gin.H) {
	return Result(map[string]string{}, errorMsg, errCode, "redirect", url)
}

// Login 跳转至登录页
func Login() (response gin.H) {
	return Result(map[string]string{}, "请先登录", 200, "login", "/")
}

// Result 标准响应
func Result(data interface{}, msg string, code int, action string, url string) (response gin.H) {
	return gin.H{
		"code":   code,
		"data":   data,
		"msg":    msg,
		"url":    url,
		"action": action,
	}
}
