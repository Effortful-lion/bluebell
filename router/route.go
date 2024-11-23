package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
)

//注册路由


func SetupRouter(mode string)(*gin.Engine){

	// 创建路由：使用自定义的日志和panic覆盖中间件

	if mode == gin.ReleaseMode{
		gin.SetMode(gin.ReleaseMode)  //设置gin框架为发布模式
	}
	//gin.SetMode(gin.ReleaseMode)  //设置gin框架为发布模式：不会有gin-debug信息
	r := gin.New()

	// 设置我信任的代理：设置一个具体的值来明确指定可信任的代理(可选)
	trustedProxies := []string{"192.168.1.0/24"}
    r.SetTrustedProxies(trustedProxies)

	r.Use(logger.GinLogger(),logger.GinRecovery(true))

	// 设置路由组，有利于后期的扩展
	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup",controller.SignUpHandler)

	// 登录业务路由
	v1.POST("/login",controller.LoginHandler)

	// 一切访问资源的路由
	// token 放在了一个 名为 Authorization 的 http 请求头中
	
	// 使用中间件
	v1.Use(middlewares.JWTAuthMiddleware(),middlewares.OnlyOneTokenMiddleware())

	// 使用代码块：社区相关路由
	{
		v1.GET("/community",controller.CommunityHandler)
	}

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