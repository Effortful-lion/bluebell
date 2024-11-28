package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
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
			ResponseError(c,CodeInvalidParam)
			return
		}
		// 是validator.ValidationErrors类型：validator校验器的错误信息是英文，转为中文
		ResponseErrorWithMsg(c,CodeInvalidParam,removeTopStruct(errs.Translate(trans)))
		return
	}
	// 业务处理 和 返回响应
	if token,err := logic.Login(p);err != nil || token == ""{
		// 登录失败: 用户名或密码错误
		if err == mysql.ErrorUserNotExist{
			ResponseError(c,CodeUserNotExist)
		}
		ResponseError(c,CodeInvalidPassword)
		return
	}else{
		// 登录成功
		ResponseSuccess(c,token)
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
			//是上面反序列化的错误:参数错误
			ResponseError(c,CodeInvalidParam)
			return
		}
		// 是validator.ValidationErrors类型：validator校验器的错误信息是英文，转为中文
		ResponseErrorWithMsg(c,CodeInvalidParam,removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务处理
	if err := logic.SignUp(p);err != nil {
		//注册失败：用户名已存在
		zap.L().Error("logic.Signup failed",zap.Error(err))
		if errors.Is(err,mysql.ErrorUserExist){
			ResponseError(c,CodeUserExist)
			return
		}
		// 数据库插入失败
		ResponseError(c,CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c,nil)
}