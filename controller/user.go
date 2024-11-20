package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)


// 用户登录
func LoginHandler(c *gin.Context) {
	// 1. 获取参数和参数校验(输入的数据格式核验)
	p := new (models.ParamLogin)
	if err := c.ShouldBindJSON(p);err != nil {
		// 请求参数错误json格式（反序列化失败）
		zap.L().Error("Login with invalid param",zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型（数据的个数或某些数据之间的基础联系（比如原密码和再次输入密码需要相同）错误）
		errs, ok := err.(validator.ValidationErrors)
		if !ok{
			//是上面反序列化的错误
			c.JSON(http.StatusOK,gin.H{
				//"msg": "请求参数错误",
				"msg": err.Error(),//详细的报错信息
			})
			return
		}
		// 是validator.ValidationErrors类型：validator校验器的错误信息是英文，转为中文
		c.JSON(http.StatusOK,gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //详细的报错信息(翻译错误)
		})
		return
	}
	// 开发时调试信息：打印参数信息
	fmt.Println(p)
	// 业务处理 和 返回响应
	if user,err := logic.Login(p);err != nil || user == nil{
		// 登录失败
		c.JSON(http.StatusOK,gin.H{
			"msg": "登录失败",
		})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"msg": "登录成功",
		})
		return
	}
}


// 用户注册
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new (models.ParamSignUp)
	if err := c.ShouldBindJSON(p);err != nil {
		// 请求参数错误
		zap.L().Error("SignUp with invalid param",zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok{
			//是上面反序列化的错误
			c.JSON(http.StatusOK,gin.H{
				//"msg": "请求参数错误",
				"msg": err.Error(),//详细的报错信息
			})
			return
		}
		// 是validator.ValidationErrors类型：validator校验器的错误信息是英文，转为中文
		c.JSON(http.StatusOK,gin.H{
			//"msg": "请求参数错误",
			"msg": removeTopStruct(errs.Translate(trans)), //详细的报错信息(翻译错误)
		})
		return
	}
	// 开发时调试信息
	fmt.Println(p)
	// 2. 业务处理
	if err := logic.SignUp(p);err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败!",
		})
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}