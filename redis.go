package utils

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

// 连接配置
type RedisOptions struct {
	Prefix   string //""
	UserName string //""
	Port     int32  //6379
	DB       int    //0
	MaxOpen  int    //100
	MinIdle  int    //20
	LockTime int    //5秒
}

// redis实例
type Redis struct {
	prefix string
	DB     *redis.Client
	lock   int
}

// 新建实例
func NewRedis(host, passwd string, options ...RedisOptions) (r Redis, err error) {
	//默认配置
	cfg := RedisOptions{
		Prefix:   "",
		UserName: "",
		Port:     6379,
		DB:       0,
		MaxOpen:  100,
		MinIdle:  20,
		LockTime: 5,
	}
	//自定义配置
	if len(options) > 0 {
		if options[0].Prefix != "" {
			cfg.Prefix = options[0].Prefix
		}
		if options[0].UserName != "" {
			cfg.UserName = options[0].UserName
		}
		if options[0].Port > 0 {
			cfg.Port = options[0].Port
		}
		if options[0].DB > 0 {
			cfg.DB = options[0].DB
		}
		if options[0].MaxOpen > 0 {
			cfg.MaxOpen = options[0].MaxOpen
		}
		if options[0].MinIdle > 0 {
			cfg.MinIdle = options[0].MinIdle
		}
		if options[0].LockTime > 0 {
			cfg.LockTime = options[0].LockTime
		}
	}
	r.prefix = cfg.Prefix
	r.lock = cfg.LockTime
	r.DB = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			host,
			cfg.Port,
		),
		Username:     cfg.UserName,
		Password:     passwd,
		DB:           cfg.DB,
		PoolSize:     cfg.MaxOpen,
		MinIdleConns: cfg.MinIdle,
	})
	_, err = r.DB.Ping().Result()
	return
}

// 锁定
func (r *Redis) Lock(name string) (bool, error) {
	key := r.prefix + ":redis:lock:" + name
	return r.DB.SetNX(key, "lock", time.Duration(r.lock)*time.Second).Result()
}

// 解锁
func (r *Redis) UnLock(name string) error {
	key := r.prefix + ":redis:lock:" + name
	return r.DB.Del(key).Err()
}
