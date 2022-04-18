package service

import (
	"back/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
)

type sJWT struct {
}

var (
	insJWT = sJWT{}
	jwtKey = []byte("secret")
)

func JWT() *sJWT {
	return &insJWT
}

type Claims struct {
	Context model.JWTContext
	jwt.StandardClaims
}

func (j sJWT) Generate(jwtCtx model.JWTContext) (string, error) {
	// 在这里声明令牌的到期时间，我们将其保留为5分钟
	expirationTime := time.Now().Add(24 * time.Hour)
	// 创建JWT声明，其中包括用户名和有效时间
	claims := &Claims{
		Context: jwtCtx,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 使用用于签名的算法和令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 创建JWT字符串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j sJWT) Parse(token string) (model.JWTContext, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return model.JWTContext{}, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims.Context, nil
		}
	}
	return model.JWTContext{}, gerror.New("token error")
}
