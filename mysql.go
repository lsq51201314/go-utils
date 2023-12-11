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

type MysqlOptions struct {
	Host    string
	User    string
	Passwd  string
	DbName  string
	Port    int
	MaxOpen int
	MinIdle int
}

type Mysql struct {
	DB *gorm.DB
}

func NewMysql(options MysqlOptions) (m Mysql, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		options.User,
		options.Passwd,
		options.Host,
		options.Port,
		options.DbName)
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
		sqlDB.SetMaxOpenConns(options.MaxOpen)
		sqlDB.SetMaxIdleConns(options.MinIdle)
	}
	return
}
