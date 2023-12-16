package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 连接配置
type MysqlOptions struct {
	User    string //"root"
	Port    int    //3306
	MaxOpen int    //100
	MinIdle int    //20
}

// mysql实例
type Mysql struct {
	DB *gorm.DB
}

// 新建实例
func NewMysql(host, passwd, dbname string, options ...MysqlOptions) (m Mysql, err error) {
	//默认配置
	cfg := MysqlOptions{
		User:    "root",
		Port:    3306,
		MaxOpen: 100,
		MinIdle: 20,
	}
	//自定义配置
	if len(options) > 0 {
		if options[0].User != "" {
			cfg.User = options[0].User
		}
		if options[0].Port > 0 {
			cfg.Port = options[0].Port
		}
		if options[0].MaxOpen > 0 {
			cfg.MaxOpen = options[0].MaxOpen
		}
		if options[0].MinIdle > 0 {
			cfg.MinIdle = options[0].MinIdle
		}
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		passwd,
		host,
		cfg.Port,
		dbname)
	if m.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		CreateBatchSize: 100, //分批插入
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	}); err != nil {
		return
	}
	//连接池
	var sqlDB *sql.DB
	if sqlDB, err = m.DB.DB(); err != nil {
		return
	} else {
		sqlDB.SetMaxOpenConns(cfg.MaxOpen)
		sqlDB.SetMaxIdleConns(cfg.MinIdle)
	}
	return
}
