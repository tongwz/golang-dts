package util

import (
	"golang-dts/pkg/setting"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecrect)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成用户名密码的token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	// 初始化结构体
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 验证token过程 没看懂
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if Claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return Claims, nil
		}
	}
	return nil, err
}
