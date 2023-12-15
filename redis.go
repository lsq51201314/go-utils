package utils

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisOptions struct {
	Prefix   string
	Host     string
	Port     int32
	Passwd string
	MaxOpen  int
	MinIdle  int
}

type Redis struct {
	prefix string
	DB     *redis.Client
}

func NewRedis(options RedisOptions) (r Redis, err error) {
	r.prefix = options.Prefix
	r.DB = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			options.Host,
			options.Port,
		),
		Username:     "",
		Password:     options.Passwd,
		DB:           0,
		PoolSize:     options.MaxOpen,
		MinIdleConns: options.MinIdle,
	})
	_, err = r.DB.Ping().Result()
	return
}

func (r *Redis) Lock(name string) (bool, error) {
	key := r.prefix + ":redis:lock:" + name
	//锁定10秒
	return r.DB.SetNX(key, "lock", time.Duration(10)*time.Second).Result()
}

func (r *Redis) UnLock(name string) error {
	key := r.prefix + ":redis:lock:" + name
	return r.DB.Del(key).Err()
}
