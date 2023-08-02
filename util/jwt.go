package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/securemist/douyin-mini/model/constant"
	"time"
)

var jwtSecret = []byte(constant.Secret)

type Claims struct {
	Id int64 `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(userId int64) (string, error) {
	nowTime := time.Now()
	// 文档中没有体现token过期的情况
	expireTIme := nowTime.Add(time.Hour * 24 * 30)

	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expireTIme.Unix(),
			Issuer:    "douyin1562",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
