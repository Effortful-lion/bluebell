package router

import (
	"bluebell/controller"
	"bluebell/logger"

	"github.com/gin-gonic/gin"
)

//注册路由


func SetupRouter()(*gin.Engine){

	// 创建路由：使用自定义的日志和panic覆盖中间件
	//gin.SetMode(gin.ReleaseMode)  //设置为生产环境
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	// 注册业务路由
	r.POST("/signup",controller.SignUpHandler)

	// 没有路由
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,gin.H{
			"code":404,
			"message":"404 not found",
		})
	})
	
	// 返回
	return r

}