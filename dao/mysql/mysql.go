package mysql

import (
	"fmt"
	"gee-Init/config"
	"gee-Init/model"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg *config.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 默认情况下，GORM会在每个操作上启动一个新事务，如果该操作已经位于事务中，则不会启动新事务。
		SkipDefaultTransaction: true,
	})
	if err != nil {
		zap.L().Error("[dao mysql Init] connect mysql error ", zap.Error(err))
		return
	}
	//	自动迁移
	err = db.AutoMigrate(&model.User{}, &model.Post{})
	if err != nil {
		zap.L().Info("[dao mysql Init] create table failed ", zap.Error(err))
		return err
	}

	conn, err := db.DB()
	if err != nil {
		zap.L().Info("[dao mysql Init] get sql instance failed ", zap.Error(err))
		return err
	}
	conn.SetMaxOpenConns(cfg.MaxOpenConn)
	conn.SetMaxIdleConns(cfg.MaxIdleConn)
	return
}

func Close() {
	conn, err := db.DB()
	zap.L().Info("[dao mysql Close] get sql instance failed ", zap.Error(err))
	err = conn.Close()
	zap.L().Info("[dao mysql Close] close the mysql connect failed ", zap.Error(err))
}
