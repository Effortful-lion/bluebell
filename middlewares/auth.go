package middlewares

import (
	"bluebell/pkg/jwt"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"bluebell/controller"
	"bluebell/dao/redis"
	"fmt"
	//"time" 
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": controller.CodeNeedLogin,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": controller.CodeInvalidToken,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": controller.CodeInvalidToken,
				"msg":  "Token过期: " + err.Error(),
			})
			c.Abort()
			return
		}

		// TODO:尝试加一段代码，判断token过期（发现已经具备判断token过期的功能了）
		// 判断令牌是否过期
        // if mc.ExpiresAt.IsZero() {
        //     c.JSON(http.StatusOK, gin.H{
        //         "code": controller.CodeInvalidToken,
        //         "msg":  "token已过期",
        //     })
        //     c.Abort()
        //     return
        // }
        // now := time.Now()
        // if mc.ExpiresAt.Before(now) {
        //     c.JSON(http.StatusOK, gin.H{
        //         "code": controller.CodeInvalidToken,
        //         "msg":  "token已过期",
        //     })
        //     c.Abort()
        //     return
        // }


		// 将当前请求的userID信息保存到请求的上下文context c中：
		c.Set(controller.ContextUserIDKey, mc.UserID)
		// 后续的处理请求的函数可以用c.Get("userId")来获取当前请求的用户信息
		c.Next() 
	}
}



// TODO: 但是只要用户名和密码正确，就会出现“相互顶号”的情况
// 基于token实现的只能一个用户登录的中间件:token格式正确，判断是否为同一用户（/时间/设备...）
func OnlyOneTokenMiddleware() func(c *gin.Context) {
	
	return func(c *gin.Context) {
		// 从请求头中获取新token
		new_token := c.Request.Header.Get("Authorization")
		new_token = strings.TrimPrefix(new_token, "Bearer ")
		fmt.Println(new_token)
	// 判断新旧token是否一致
		//如果不一致, 返回未登录//TODO:按道理应该有个验证码的业务，来确定使用哪一个token（这里默认可以“相互顶号”）
		userID,_:= c.Get(controller.ContextUserIDKey)
		redis_token,_ := redis.GetUserToken(userID.(int64))
		fmt.Println(redis_token)
		if redis_token != new_token {
			// redis.SetUserToken(new_token,userID.(int64))
			c.JSON(http.StatusOK,gin.H{
				"code":controller.CodeInvalidToken,
				"msg":"账号已在异地登录,请重新登录",
			})
			c.Abort()
			return
		}
		//如果一致放行
		c.Next()
	}
	
}
