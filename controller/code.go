package controller


// 定义状态码 + 定义对应状态码的信息 + 定义获取状态码的方法 


// 作用： 1.起别名  2.实际上可以为int64类型的数据写方法（基本类型不能写方法）
type ResCode int64 

// 定义一组常量——状态码：前后端商量定义好的一系列业务状态码
const(
	// iota 从0开始，每次加1
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

// 状态码对应的msg
var codeMsgMap = map[ResCode]string{
	CodeSuccess: "success",
	CodeInvalidParam: "请求参数错误",
	CodeUserExist: "用户已存在",
	CodeUserNotExist: "用户不存在",
	CodeInvalidPassword: "用户或密码错误",
	CodeServerBusy: "服务繁忙",
}

// 取得code对应的msg
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		// 无法获取，服务器服务繁忙（内部有问题，查看日志问题）
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}



