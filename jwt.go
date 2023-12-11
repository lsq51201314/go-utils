package utils

import (
	"crypto/sha1"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtcustom struct {
	Custom interface{} `json:"custom"`
	jwt.StandardClaims
}

type JwtOptions struct {
	Issuer string
	Passwd string
	Expire int64
}

type Jwt struct {
	options JwtOptions
}

func NewJWT(options JwtOptions) (j Jwt, err error) {
	j.options = options
	return
}

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
	err = errors.New("invalid  token")
	return
}
