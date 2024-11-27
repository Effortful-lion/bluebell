package models

const (
	// 默认页码
	DefaultPage = 1
	// 默认每页显示数量
	DefaultSize = 10
	// 按照时间排序
	OrderTime = "time"
	// 按照分数排序
	OrderScore = "score"
)

// 定义请求的参数结构体：tag 定义的名称是请求的名字和结构的映射
type ParamSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// 定义登陆的参数结构体
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

// 接收url格式参数： /api/v1/post2?page=1&size=10&order=time
type ParamPostList struct {
	CommunityID int64 `json:"community_id" form:"community_id"`		// 可以为空，如果为空，则查询所有帖子，否则查询相应社区的帖子
	Page        int64 `json:"page" form:"page"`
	Size        int64 `json:"size" form:"size"`
	Order       string `json:"order" form:"order"`
}

