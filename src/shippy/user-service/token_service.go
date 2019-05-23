package main

import (
	"github.com/dgrijalva/jwt-go"
	userProto "shippy/user-service/proto/user"
	"time"
)

type Authable interface {
	Decode(tokenStr string) (*CustomClaims, error)
	Encode(user *userProto.User) (string, error)
}

// 定义加盐哈希密码时所用的盐
// 要保证其生成和保存都足够安全
// 比如使用 md5 来生成
var privateKey = []byte("`xs#a_1-!")

// 自定义的 metadata
// 在加密后作为 JWT 的第二部分返回给客户端
type CustomClaims struct {
	User *userProto.User
	// 使用标准的 payload
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

func (svc *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return privateKey, nil
	})
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (svc *TokenService) Encode(user *userProto.User) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := &CustomClaims{user, jwt.StandardClaims{
		Issuer:    "",
		ExpiresAt: expireTime,
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(privateKey)
}
