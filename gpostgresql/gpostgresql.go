package gpostgresql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 新建连接
func New(host string, port int, user, passwd, dbname string) (*gorm.DB, error) {
	pdb, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, user, passwd, dbname, port),
	}), &gorm.Config{
		TranslateError:  true, //开启重复插入错误生效
		CreateBatchSize: 100,  //分批插入
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Error, //日志等级，Info可输出SQL语句
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	})
	if err != nil {
		return nil, err
	}
	//连接池
	sqlDB, err := pdb.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	return pdb, nil
}
