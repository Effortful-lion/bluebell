package models

// 定义请求的参数结构体：tag 定义的名称是请求的名字和结构的映射
type ParamSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 投票数据
type ParamVoteData struct {
	//UserID      从请求中获取当前用户
	PostID string `json:"post_id" binding:"required"`	// 帖子id   ：`json:"post_id,string" binding:"required"` 这种有问题！！！
	Direction int8 `json:"direction,string" binding:"oneof=-1 0 1"` // 赞成票(1)还是反对票(-1)还是取消投票(0)
}