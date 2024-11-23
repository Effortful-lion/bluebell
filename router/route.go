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
	// 注册业务路由
	r.POST("/signup",controller.SignUpHandler)

	// 登录业务路由
	r.POST("/login",controller.LoginHandler)

	// 一切访问资源的路由
	// token 放在了一个 名为 Authorization 的 http 请求头中
	r.GET("/ping", middlewares.JWTAuthMiddleware(),middlewares.OnlyOneTokenMiddleware(),func(c *gin.Context){
		// 在中间件中进行 判断请求头中是否有 有效的jwt，有则继续，没有则驳回请求
		// 正常访问全部资源
		c.JSON(200,gin.H{
			"msg":"访问到了所有资源",
		})

	})

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