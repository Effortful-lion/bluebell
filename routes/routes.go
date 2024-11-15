package routes

import (
	"bluebell/logger"
	"github.com/gin-gonic/gin"
)

//注册路由


func Setup()(*gin.Engine){
	// 创建路由：使用自定义的日志和panic覆盖中间件
	//gin.SetMode(gin.ReleaseMode)  //设置为生产环境
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"code":200,
			"message":"hello /",
		})
	})

	r.GET("/index",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"code":200,
			"message":"hello index",
		})
	})

	// 返回
	return r

}