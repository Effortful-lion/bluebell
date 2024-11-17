package mysql

import (
	"fmt"
	"bluebell/settings"
	_ "github.com/go-sql-driver/mysql" // 不要忘了导入数据库驱动
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// 小写db不能对外暴露
var db *sqlx.DB

func Init(cfg *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
	cfg.User,
	cfg.Password,
	cfg.Host,
	cfg.Port,
	cfg.DB,
)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		// 存错误信息日志
		db.SetMaxIdleConns(cfg.MaxIdleConns)
		db.SetMaxOpenConns(cfg.MaxOpenConns)
		zap.L().Error("connect DB failed.", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxIdleConns)
	db.SetMaxIdleConns(cfg.MaxOpenConns)
	return
}

// 对外暴露一个关闭方法
func Close() {
    if db != nil {
        if err := db.Close(); err != nil {
            zap.L().Info("MySQL close failed: %v", zap.Error(err))
			//log.Fatal("MySQL close failed: %v", err)
        } else {
            zap.L().Info("MySQL close success...")
			//log.Fatal("MySQL close success...")
        }
    } else {
        zap.L().Info("MySQL db is not initialized, nothing to close.")
		//log.Fatal("MySQL db is not initialized, nothing to close.")
    }
	//修改之前关闭服务器的报错：空指针
	//_ = db.Close 
}
