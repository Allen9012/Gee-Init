package config_init

// @Title  请填写文件名称（需要改）
// @Description  防止循环引用提取文件
// @Author  请填写自己的真实姓名（需要改）  ${DATE} ${TIME}
// @Update  请填写自己的真实姓名（需要改）  ${DATE} ${TIME}
import (
	"fmt"
	"gee-Init/config"
	"gee-Init/config/conf"
	"gee-Init/dao/mysql"
	"gee-Init/dao/redis"
	"gee-Init/util"
	"gee-Init/util/logger"

	"go.uber.org/zap"
)

// Init 初始化配置项
func Init() error {
	//	使用viper读取配置文件
	err := conf.InitWithViper()
	if err != nil {
		return err
	}

	// 设置日志级别
	if err = logger.Init(conf.Conf.LogConfig, conf.Conf.Mode); err != nil {
		fmt.Printf("config_init logger error : %s \n", err)
		return err
	}
	defer zap.L().Sync()
	// 读取翻译文件
	if err := config.LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		// zap输出错误日志使用到err
		zap.L().Panic("翻译文件加载失败 %v", zap.Error(err))
	}

	// 连接数据库
	//model.Database(os.Getenv("MYSQL_DSN"))
	//cache.Redis()

	//初始化数据库
	if err := mysql.Init(conf.Conf.MySQLConfig); err != nil {
		zap.L().Panic("config_init mysql error : %s \n", zap.Error(err))
		return err
	}
	defer mysql.Close()
	//初始化redis
	if err := redis.Init(conf.Conf.RedisConfig); err != nil {
		zap.L().Panic("config_init redis error : %s \n", zap.Error(err))
		return err
	}
	defer redis.Close()
	// 初始化雪花
	if err := util.Init(conf.Conf.StartTime, conf.Conf.MachineID); err != nil {
		zap.L().Panic("config_init snowflake error : %s \n", zap.Error(err))
		return err
	}

	return nil

}
