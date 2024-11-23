package settings

import (
	"fmt"
	"time"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//作用： 管理配置文件

//这里还可以继续进行修改，使得settings用结构体保存配置文件信息：代码清晰、易读性好
//注意：如果使用全局变量的结构体存储配置文件信息，那么在配置文件发生实时变化的时候就需要利用回调函数进行再一次的反序列化，更新结构体

// Conf 全局变量，用来保存程序的所有配置信息
// 这里的mapstructure标签，简单的说是一个映射：配置文件名 -> 结构体名
var Conf = new(Config)

type Config struct {
	*AppConfig       `mapstructure:"app"`
	*LogConfig       `mapstructure:"log"`
	*MysqlConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port int `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID int `mapstructure:"machine_id"`
	*LogConfig `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize int `mapstructure:"max_size"`
	MaxAge int `mapstructure:"max_age"`
	MaxBackups int `mapstructure:"max_backups"`
}

type MysqlConfig struct {
    Host string `mapstructure:"host"`
    User string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DB string `mapstructure:"dbname"`
    Port int `mapstructure:"port"`
    MaxOpenConns int `mapstructure:"max_open_conns"`
    MaxIdleConns int `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port string `mapstructure:"port"`
	DB int `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
}

func Init()(err error){
	
	//指定配置文件
	//viper.SetConfigFile("./settings/config.yaml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./settings/")
	err = viper.ReadInConfig()		//从配置文件中读取配置项
	if err != nil{
		// 文件未找到，读取失败
		fmt.Printf("配置文件读取失败:%v",err)
		return
	}
	// 反序列化到配置变量的结构体中
	if err := viper.Unmarshal(Conf);err != nil {
		fmt.Printf("viper.Unmarshal failed,err: %v",err)
	}

	viper.WatchConfig()	 // 监控配置文件
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改")
		time.Sleep(1 * time.Second) // 延迟1秒，可根据实际情况调整
		if err := viper.Unmarshal(Conf); err!= nil {
			fmt.Printf("viper.Unmarshal change failed, err: %v", err)
			// 可以在这里添加更完善的错误处理逻辑，比如尝试重新读取等操作
		}
	})
	return
}
