package main

//go web 开发较通用脚手架模板
//一个web项目中，后端代码部分：
// 1.读取配置文件
// 2.初始化日志
// 3.初始化数据库连接（mysql、redis）
// 4.注册路由
// 5.启动服务（优雅关机）

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/settings"
	//"context"
	"fmt"
	//"net/http"
	//"os"
	//"os/signal"
	//"syscall"
	//"time"

	"go.uber.org/zap"
)

//go中，包下的方法调用时，只需要导入包就可以，至于包中需要的其他依赖，在那个包中已经被声明过了

func main(){
//初始化错误，全部return

	// 1.读取配置文件
	if err := settings.Init();err != nil{
		fmt.Printf("init settings failed:%v\n",err)
		return
	}
	// 2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig,settings.Conf.Mode); err != nil{
		fmt.Printf("init logger failed:%v\n",err)
		return
	}
	//通常用于将日志缓冲区中的内容（如果有的话）刷新到对应的输出目标（比如文件、标准输出等）
	//避免出现日志丢失等情况，保证日志记录的完整性
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	// 3.初始化mysql
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil{
		fmt.Printf("init mysql failed:%v\n",err)
		return
	}
	defer mysql.Close()
	// 初始化redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil{
		fmt.Printf("init redis failed:%v\n",err)
		return
	}
	defer redis.Close()
	// 初始化雪花算法
	if err := snowflake.Init(settings.Conf.AppConfig.StartTime, int64(settings.Conf.AppConfig.MachineID)); err != nil{
		fmt.Printf("init snowflake failed:%v\n",err)
		return
	}

	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil{
		fmt.Printf("init validator trans failed,err:%v\n",err)
		return
	}

	// 4.注册路由
	r := router.SetupRouter(settings.Conf.Mode)

	if err := r.Run(fmt.Sprintf("127.0.0.1:%d",settings.Conf.AppConfig.Port)); err != nil{
		fmt.Printf("server startup failed,err:%v\n",err)
		return
	}


	// 5.启动服务（优雅关机）
	// srv := &http.Server{
	// 	Addr:    fmt.Sprintf("127.0.0.1:%d",settings.Conf.AppConfig.Port),
	// 	Handler: r,
	// }

	// go func() {
	// 	// 开启一个goroutine启动服务
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		zap.L().Fatal("listen: %s\n", zap.Error(err))
	// 	}
	// }()

	// // 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	// quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// // kill 默认会发送 syscall.SIGTERM 信号
	// // kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// // kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// // signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)  // 此处不会阻塞
	// <-quit  // 阻塞在此，当接收到上述两种信号时才会往下执行
	// zap.L().Info("Shutdown Server ...")
	// // 创建一个5秒超时的context
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// // 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	// if err := srv.Shutdown(ctx); err != nil {
	// 	zap.L().Fatal("Server Shutdown", zap.Error(err))
	// }

	// zap.L().Info("Server exiting")
}