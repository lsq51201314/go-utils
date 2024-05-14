package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtcustom struct {
	Custom interface{} `json:"custom"`
	jwt.StandardClaims
}

// 配置信息
type JwtOptions struct {
	Issuer string //""
	Passwd string //go-utils-jwt
	Expire int64  //7776000秒(90天)
}

// jwt实例
type Jwt struct {
	options JwtOptions
}

// 新建实例
func NewJWT(issuer string, options ...JwtOptions) (j Jwt, err error) {
	j.options = JwtOptions{
		Issuer: issuer,
		Passwd: "go-utils-jwt",
		Expire: 7776000,
	}
	if len(options) > 0 {
		if options[0].Passwd != "" {
			j.options.Passwd = options[0].Passwd
		}
		if options[0].Expire > 0 {
			j.options.Expire = options[0].Expire
		}
	}
	//生成密码
	h := sha1.New()
	if _, err = h.Write([]byte(j.options.Passwd)); err != nil {
		return
	}
	j.options.Passwd = hex.EncodeToString(h.Sum(nil))
	return
}

// 生成凭证
func (j *Jwt) GetToken(custom interface{}) (token string, err error) {
	key := sha1.Sum([]byte(j.options.Passwd))
	p := jwtcustom{
		Custom: custom,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(j.options.Expire) * time.Second).Unix(),
			Issuer:    j.options.Issuer,
			NotBefore: time.Now().Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	return t.SignedString(key[:])
}

// 解析凭证
func (j *Jwt) ParseToken(token string) (custom interface{}, err error) {
	key := sha1.Sum([]byte(j.options.Passwd))
	var t *jwt.Token
	if t, err = jwt.ParseWithClaims(token, &jwtcustom{},
		func(token *jwt.Token) (i interface{}, err error) {
			return key[:], nil
		}); err != nil {
		return
	}
	if info, ok := t.Claims.(*jwtcustom); ok && t.Valid {
		custom = info.Custom
		return
	}
	err = errors.New("无效的凭证")
	return
}
