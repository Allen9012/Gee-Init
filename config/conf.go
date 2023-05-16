package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	Port         int    `mapstructure:"port"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Dbname      string `mapstructure:"dbname"`
	Port        int    `mapstructure:"port"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func initWithViper() (err error) {
	//TODO 设置默认值
	viper.SetDefault("app.Name", "Gee-Init")
	viper.SetDefault("app.port", 8080)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("REDIS_ADDR", "127.0.0.1:6379")
	//	设置读取的配置文件
	viper.SetConfigFile("config.yaml")
	//	设置读取的配置文件路径
	viper.AddConfigPath(".")
	//	读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	//	将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(&Conf); err != nil {
		zap.L().Error("viper.Unmarshal failed, err : %v", zap.Error(err))
		return
	}
	//	监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(&Conf); err != nil {
			zap.L().Error("viper.Unmarshal failed, err : %v", zap.Error(err))
			return
		}
	})
	return
}
