package redis

import (
	"fmt"
	"bluebell/settings"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
        PoolSize: cfg.PoolSize,
	})

    _, err = rdb.Ping().Result()
    return err
}

func Close() {
	_ = rdb.Close()
}