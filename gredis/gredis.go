package gredis

import (
	"fmt"

	"github.com/go-redis/redis"
)

// 新建连接
func New(host string, port int, passwd string, db ...int) (*redis.Client, error) {
	dbi := 0
	if len(db) > 0 && db[0] > 0 {
		dbi = db[0]
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Password:     passwd,
		DB:           dbi,
		PoolSize:     100,
		MinIdleConns: 10,
	})
	if _, err := rdb.Ping().Result(); err != nil {
		return nil, err
	}
	return rdb, nil
}
