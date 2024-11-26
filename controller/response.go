package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


// 创建  向前端 响应结构体信息和对应（成功、失败状态码）的响应方法


// 响应结构体
type ResponseData struct {
	Code ResCode    `json:"code"` // 自定义的code
	Msg  any   `json:"msg"`  // 自定义的msg
	Data any 	`json:"data"`// 自定义的数据  ,omitempty可忽略空值不展示  Data any 	`json:"data,omitempty"`
}


//响应错误信息：code+错误信息
func ResponseError(c *gin.Context, code ResCode){
	// map[string]any = gin.H{
	// 	"code":code,
	// 	"msg":getMsg(code),
	// 	"data":nil,
	// }
	// 以上直接使用gin框架返回一个map，或者我们直接使用我们定义好的结构体
	c.JSON(http.StatusOK,&ResponseData{
		Code: code,
		Msg:code.Msg(),
		Data:nil,
	})
}

//错误的信息（较具体的原因，可自定义）
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg any){
	c.JSON(http.StatusOK,&ResponseData{
		Code:code,
		Msg:msg,
		Data:nil,
	})
}

//响应成功信息：code+成功信息+数据
func ResponseSuccess(c *gin.Context, data any){
	// json方法自动将下面的数据装箱为json类型
	c.JSON(http.StatusOK,&ResponseData{
		Code:CodeSuccess,
		Msg:CodeSuccess.Msg(),
		Data:data,
	})
}
