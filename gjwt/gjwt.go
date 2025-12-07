package gjwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lsq51201314/go-utils/gsha"
)

type customClaims struct {
	UserID int64 `json:"uid,string"`
	jwt.RegisteredClaims
}

type Token struct {
	key     []byte
	expires time.Duration
}

// 默认5分钟
func New(key string, expires ...int) *Token {
	obj := Token{
		key:     gsha.Sha1([]byte(key)),
		expires: time.Duration(300) * time.Second, //5分钟
	}
	if len(expires) > 0 && expires[0] > 0 {
		obj.expires = time.Duration(expires[0]) * time.Second
	}
	return &obj
}

// 生成凭证
func (t *Token) Generate(userId int64) (string, error) {
	n := time.Now()
	claims := &customClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(n.Add(t.expires)),
			IssuedAt:  jwt.NewNumericDate(n),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.key)
}

// 验证凭证
func (t *Token) Validate(token string) (int64, error) {
	tk, err := jwt.ParseWithClaims(token, &customClaims{},
		func(token *jwt.Token) (any, error) {
			return t.key, nil
		})
	if err != nil {
		return 0, err
	}
	if claims, ok := tk.Claims.(*customClaims); ok && tk.Valid {
		return claims.UserID, nil
	}
	return 0, errors.New("凭证无效")
}
