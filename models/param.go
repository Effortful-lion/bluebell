package models

// 定义请求的参数结构体：tag 定义的名称是请求的名字和结构的映射
type ParamSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}