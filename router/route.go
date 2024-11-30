package router

import (
	"bluebell/controller"
	docs "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"
	//"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	gs "github.com/swaggo/gin-swagger"     // gin-swagger middleware

	"github.com/gin-contrib/pprof"

	"github.com/gin-contrib/cors"		// 一个用于处理跨域的三方库
)

//注册路由
func SetupRouter(mode string)(*gin.Engine){

	// 创建路由：使用自定义的日志和panic覆盖中间件

	if mode == gin.ReleaseMode{
		gin.SetMode(gin.ReleaseMode)  //设置gin框架为发布模式
	}
	//gin.SetMode(gin.ReleaseMode)  //设置gin框架为发布模式：不会有gin-debug信息
	r := gin.New()

	// 一种 “前后端文件冗杂一起的部署模式：” 这里前端页面和资源也是在go中进行加载的
	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")

	docs.SwaggerInfo.BasePath = "/api/v1"

	//配置跨域中间件: 否则会导致 前端发送：Options的预检请求
    // r.Use(func(c *gin.Context) {
    //     c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    //     c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    //     c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
    //     if c.Request.Method == "OPTIONS" {
    //         c.AbortWithStatus(http.StatusNoContent)
    //         return
    //     }
    //     c.Next()
    // })
	
	// 这里也可以直接使用github上的一个github.com/gin-contrib/cors库，直接使用
	r.Use(cors.Default())

	// 设置我信任的代理：设置一个具体的值来明确指定可信任的代理(可选)
	trustedProxies := []string{"192.168.1.0/24"}
    r.SetTrustedProxies(trustedProxies)

	// gin-swagger同时还提供了DisablingWrapHandler函数，方便我们通过设置某些环境变量来禁用Swagger。例如：
	// r.GET("/swagger/*any", gs.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))


	// 对全网的请求限速：可以对r；或者可以对api进行限速
	//r.Use(logger.GinLogger(),logger.GinRecovery(true),middlewares.RateLimitMiddleware(2 * time.Second, 1))
	//r.Use(logger.GinLogger(),logger.GinRecovery(true),middlewares.RateLimitMiddleware2(1))
	r.Use(logger.GinLogger(),logger.GinRecovery(true))

	//加上这一段，会有什么效果？：设置api的根路径
	//docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",nil)
	})
	
	// 设置路由组，有利于后期的扩展
	v1 := r.Group("/api/v1")

	{
		// 注册业务路由
		v1.POST("/signup",controller.SignUpHandler)
		// 登录业务路由
		v1.POST("/login",controller.LoginHandler)
	}

	{
	// GET 允许不登录查看
		// 根据帖子分数/创建时间，查询帖子列表(所有帖子)
		v1.GET("/post2",controller.GetPostListHandler2)
		// 查询帖子详情
		v1.GET("/post/:id",controller.GetPostDetailHandler)
		// 使用代码块：社区相关路由
	{
		// 获取社区分类列表
		v1.GET("/community",controller.CommunityHandler)
		// 获取社区分类详情（路径参数）
		v1.GET("/community/:id",controller.CommunityDetailHandler)
	}

	}
	

	// 一切访问资源的路由
	// token 放在了一个 名为 Authorization 的 http 请求头中
	
	// 使用中间件
	v1.Use(middlewares.JWTAuthMiddleware(),middlewares.OnlyOneTokenMiddleware())
	
	

	// 帖子相关路由
	{
		// 创建帖子
		v1.POST("/post",controller.CreatePostHandler)
		// 查询贴子列表【没用了】
		//v1.GET("/post",controller.GetPostListHandler)
	}

	// 帖子细节(分数等)相关路由
	{
		// 帖子投票
		v1.POST("/vote",controller.PostVoteHandler)
	}

	// 没有路由
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,gin.H{
			"code":404,
			"message":"404 not found",
		})
	})

	r.GET("/swagger/*any",  gs.WrapHandler(swaggerfiles.Handler))

	pprof.Register(r)	// 注册pprof相关路由
	
	// 返回
	return r

}

