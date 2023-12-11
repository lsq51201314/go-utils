# go-utils

```go
package main

import (
	"errors"
	"fmt"

	"github.com/lsq51201314/go-utils"
	"go.uber.org/zap"
)

func main() {
	//配置文件
	utils.YamlLoad(Configs)
	fmt.Println(Configs.App.Mode)
	fmt.Println(Configs.Log.Level)
	fmt.Println(Configs.MySQL.DbName)
	fmt.Println(Configs.Redis.Prefix)
	fmt.Println(Configs.Token.Issuer)
	//获取验证
	var vc utils.VerifyCode
	id, code, _, _ := vc.GetCode()
	fmt.Println(id, code)
	//压缩解压
	var zlib utils.Zlib
	c, _ := zlib.Compress([]byte("Hello World"))
	u, _ := zlib.UnCompress(c)
	fmt.Println(string(u))
	//随机索引
	var sfid utils.Snowflake
	sfid, _ = utils.NewSnowflake("2023-12-01", 1)
	fmt.Println(sfid.GetID())
	//redis
	var redis utils.Redis
	redis, _ = utils.NewRedis(utils.RedisOptions{
		Prefix:  Configs.Redis.Prefix,
		Host:    Configs.Redis.Host,
		Port:    Configs.Redis.Port,
		Passwd:  Configs.Redis.Passwd,
		MaxOpen: Configs.Redis.MaxOpen,
		MinIdle: Configs.Redis.MinIdle,
	})
	fmt.Println(redis.Lock("localhost"))
	defer redis.UnLock("localhost")
	//mysql
	var mysql utils.Mysql
	mysql, _ = utils.NewMysql(utils.MysqlOptions{
		Host:    Configs.MySQL.Host,
		User:    Configs.MySQL.User,
		Passwd:  Configs.MySQL.Passwd,
		DbName:  Configs.MySQL.DbName,
		Port:    Configs.MySQL.Port,
		MaxOpen: Configs.Redis.MaxOpen,
		MinIdle: Configs.Redis.MinIdle,
	})
	fmt.Println(mysql.DB.Migrator().HasTable("test"))
	//zap
	utils.ZapInit(utils.ZapOptions{
		Level:      Configs.Log.Level,
		FileName:   Configs.Log.FileName,
		MaxSize:    Configs.Log.MaxSize,
		MaxAge:     Configs.Log.MaxAge,
		MaxBackups: Configs.Log.MaxBackups,
	})
	zap.L().Error("test error")
	//http
	var http utils.Http
	http.Get("http://www.baidu.com")
	http.Post("http://www.baidu.com", nil)
	//const
	var constant utils.Const
	constant.AutoBuild(&ErrorCode)
	fmt.Println(ErrorCode.LoginChangeError)
	//aliyunoss
	var aliyunoss utils.AliyunOSS
	aliyunoss, _ = utils.NewAliyunOSS(utils.AliyunOSSOptions{
		AccessKeyId:     Configs.YunOSS.AccessKeyId,
		AccessKeySecret: Configs.YunOSS.AccessKeySecret,
		OssEndpoint:     Configs.YunOSS.OssEndpoint,
		OssBucketName:   Configs.YunOSS.OssBucketName,
	})
	name, _ := aliyunoss.Upload([]byte("hello world"))
	fmt.Println(name)
	ossbuf, _ := aliyunoss.Download("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
	fmt.Println(string(ossbuf))
	fmt.Println(aliyunoss.Delete("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"))
	//yunfile
	var yunfile utils.YunFile
	yunfile, _ = utils.NewYunFile()
	name, _ = yunfile.Upload([]byte("hello world"))
	fmt.Println(name)
	ossbuf, _ = yunfile.Download("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
	fmt.Println(string(ossbuf))
	fmt.Println(yunfile.Delete("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"))
	//jwt
	var jwt utils.Jwt
	jwt, _ = utils.NewJWT(utils.JwtOptions{
		Issuer: Configs.Token.Issuer,
		Passwd: Configs.Token.Passwd,
		Expire: Configs.Token.Expire,
	})
	str, _ := jwt.GetToken("hello world")
	fmt.Println(str)
	fmt.Println(jwt.ParseToken(str))
	//logger
	var logger utils.Logger
	logger.Info("Hello world")
	logger.Error("Hello world",errors.New("Hello world"))
}
```