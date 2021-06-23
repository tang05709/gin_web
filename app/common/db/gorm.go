package db

import (
	"festival/app/common/cfg"
	"festival/app/common/log"

	// "festival/app/model/module"
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	mainDb   *gorm.DB
	onceInit sync.Once
)

var DbList = make([]interface{}, 0)

//初始化数据操作 driver为数据库类型
func Global() *gorm.DB {
	onceInit.Do(func() {
		var db *gorm.DB
		var err error
		dsn := cfg.Instance().Database.Master
		fmt.Println(dsn)
		mysqlConfig := mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         200,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据版本自动配置
		}
		if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig(cfg.Instance().Database.Debug)); err != nil {
			os.Exit(0)
		} else {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(time.Second * 300)
		}
		mainDb = db
	})
	return mainDb
}

// 自动同步数据库表
func AutoMigrate() {
	err := mainDb.AutoMigrate(DbList...)
	if err != nil {
		fmt.Println("同步数据库失败")
		log.Global().Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	log.Global().Info("register table success")
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(mod bool) *gorm.Config {
	if cfg.Instance().Database.LogMode {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
	if mod {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
}
