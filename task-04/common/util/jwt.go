package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/endymion/go-base/task-04/common/setting"
	"time"
)

var jwtSetting = setting.JWTSetting

// jwtSetting.AccessTokenSecret 转为[]byte
var accessSecret = []byte(jwtSetting.AccessTokenSecret)
var refreshSecret = []byte(jwtSetting.RefreshTokenSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(username string) (string, string, error) {
	var expireTime = time.Now().Add(time.Duration(jwtSetting.ExpireTime) * time.Hour)
	accessClaims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "task-04",
		},
	}

	refreshClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    "task-04",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessString, err := accessToken.SignedString(accessSecret)
	refreshString, err := refreshToken.SignedString(refreshSecret)

	return accessString, refreshString, err
}

// ParseToken parse jwt token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// IsTokenExpired check if the token is expired
func IsTokenExpired(err error) bool {
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return true
			}
		}
	}
	return false
}

// RefreshToken 刷新token
func RefreshToken(tokenString string) (string, error) {
	var expireTime = time.Now().Add(time.Duration(jwtSetting.ExpireTime) * time.Hour)
	standardClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    "gin-blog",
	}
	token, err := jwt.ParseWithClaims(tokenString, &standardClaims, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	claims := token.Claims.(*jwt.StandardClaims)
	newAccessToken, _, _ := GenerateToken(claims.Subject)
	return newAccessToken, nil
}
