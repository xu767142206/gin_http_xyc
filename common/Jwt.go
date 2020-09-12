package common

import (
	"fmt"
	"gin/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username, password string, uid int) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":      uid,
		"username": username,
		"password": password,
		"create":   time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 6).Unix(),
	})
	token, err := at.SignedString([]byte(config.Conf.Authorization))
	if err != nil {
		return "", err
	}
	return token, nil
}

//解密
func ParseHStoken(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.Authorization), nil
	})
	if err != nil {
		fmt.Println("HS256的token解析错误，err:", err)
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("ParseHStoken:claims类型转换失败")
		return nil
	}
	return claims
}
