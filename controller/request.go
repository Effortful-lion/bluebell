package controller

// 请求的预处理，取出context的东西，获取重要信息，为后面controller的处理请求准备调用

import (
	"errors"

	"github.com/gin-gonic/gin"
)


var ErrorUserNotLogin = errors.New("用户未登录")
var ContextUserIDKey = "userID"

// 获取当前登录的用户的ID
func getCurrentUserId(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin 
		return 
	}
	userId,ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return 
	}
	return
}