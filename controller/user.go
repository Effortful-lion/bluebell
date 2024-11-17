package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new (models.ParamSignUp)
	if err := c.ShouldBindJSON(p);err != nil {
		// 请求参数错误
		zap.L().Error("SignUp with invalid param",zap.Error(err))
		c.JSON(http.StatusOK,gin.H{
			//"msg": "请求参数错误",
			"msg": err.Error(),//详细的报错信息
		})
		return
	}
	fmt.Println(p)
	// 2. 业务处理
	logic.SignUp(p)
	// 3. 返回响应
	c.JSON(200, gin.H{
		"msg": "success",
	})
}