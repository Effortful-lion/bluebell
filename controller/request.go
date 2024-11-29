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
	return userId,nil
}

// 从请求参数中获取分页参数
// func getPageInfo(c *gin.Context) (page,size int64) {
// 	// 获取分页参数: limit页容量， offset页码
// 	pageStr := c.Query("page")
// 	sizeStr := c.Query("size")
// 	page, err := strconv.ParseInt(pageStr, 10, 64)
// 	if err != nil {
// 		page = 1  // 默认第一页
// 	}
// 	size, err = strconv.ParseInt(sizeStr, 10, 64)
// 	if err != nil {
// 		size = 10 // 默认每页显示10条记录
// 	}
// 	return
// }