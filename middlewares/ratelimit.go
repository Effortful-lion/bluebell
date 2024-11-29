package middlewares

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"

    ratelimit1 "github.com/juju/ratelimit" // 令牌桶中间件
    ratelimit2 "go.uber.org/ratelimit"     // 漏桶中间件
)

// 令牌桶中间件
// 限流中间件: 每 fillInterval 的时间，可最多处理 1 个请求
// 特点： 如果一段时间内不请求，token会存储，直到桶的最大容量cap，应对突发多个请求
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
    // 创建令牌桶 ： 参数： 1. 填充时间（单位时间） 2. 容量
    bucket := ratelimit1.NewBucket(fillInterval, cap)
    return func(c *gin.Context) {
        // 判断是否限流(如果取不到 token 就限流)
        if bucket.TakeAvailable(1) <= 0 {
            c.String(http.StatusOK, "你点击的太快了，请慢点点击~")
            // 终止执行
            c.Abort()
            return
        }
        // 取到令牌继续执行
        c.Next()
    }
}

// 漏桶中间件
// 限流中间件：
// 特点：不能应对突发的多个请求
// RateLimitMiddleware 基于漏桶算法的限流中间件
// RateLimitMiddleware 基于漏桶算法的限流中间件，使用go.uber.org/ratelimit库
func RateLimitMiddleware2(rateLimit int) gin.HandlerFunc {
    // 创建一个限流器，这里的rateLimit表示每秒允许的请求数量，单位是每秒多少个请求
    rl := ratelimit2.New(rateLimit)
    return func(c *gin.Context) {
        // 获取当前时间
        now := time.Now()
        // 尝试获取许可，该操作会按照限流器的规则进行阻塞（如果需要等待许可）: 这个虽然有返回值（下一次请求可处理的时间），但是没用，已经阻塞了
        rl.Take()
        // 计算获取许可所花费的时间
        wait := time.Since(now)
        // 如果花费时间大于0，说明进行了等待，可能需要限流处理
        if wait > 0 {
			// 原作者是 等待策略 ，我们这里直接选择丢弃
            c.String(http.StatusOK, "你点击的太快了，请慢点点击~")
            c.Abort()
            return
        }
        c.Next()
    }
}