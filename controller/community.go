package controller

import (
	"bluebell/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 社区相关的接口

func CommunityHandler(c *gin.Context){
	// 作用：向前端提供一个列表（slice），内容是 社区的(id和名称)
	// 通常controller会先处理参数等...这里没有，可以直接调用然后开始业务
	list,err := logic.GetCommunityList()
	if err != nil{
		// 获取失败，这是服务端的错误，记录在日志中（不轻易将服务端报错暴露在外面）
		zap.L().Error("logic.GetCommunityList()获取社区列表失败", zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	// 获取成功
	ResponseSuccess(c,list)
}