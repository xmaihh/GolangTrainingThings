package util

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	//key = base64.URLEncoding.EncodeToString([]byte("Uf4bGZ53Ds*#"))
	key []byte = []byte("wyr6512@163.com")
)

//产生json web token

func GenToken() string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "Uf4bGZ53Ds*#",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString(key)
	if err != nil {
		fmt.Println("generate token failed.", err)
		return ""
	}
	return str
}

// 检验token是否有效

func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parse with claims failed.", err)
		return false
	}
	return true
}
