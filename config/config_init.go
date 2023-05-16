package config

// @Title  请填写文件名称（需要改）
// @Description  防止循环引用提取文件
// @Author  请填写自己的真实姓名（需要改）  ${DATE} ${TIME}
// @Update  请填写自己的真实姓名（需要改）  ${DATE} ${TIME}
import (
	"fmt"
	"gee-Init/dao/mysql"
	"gee-Init/util/logger"

	"go.uber.org/zap"
)

// Init 初始化配置项
func Init() error {
	//	使用viper读取配置文件
	err := initWithViper()
	if err != nil {
		return err
	}

	// 设置日志级别
	if err = logger.Init(Conf.LogConfig, Conf.Mode); err != nil {
		fmt.Printf("init logger error : %s \n", err)
		return err
	}
	defer zap.L().Sync()
	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		// zap输出错误日志使用到err
		zap.L().Panic("翻译文件加载失败 %v")
	}

	// 连接数据库
	//model.Database(os.Getenv("MYSQL_DSN"))
	//cache.Redis()

	//初始化数据库
	if err := mysql.Init(Conf.MySQLConfig); err != nil {

	}
	return nil
}
