package util

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

const (
	TOKEN_KEY = "http://simple-douyin.com"
)

func GenToken(userId int64, expSec int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    expSec,
		"userId": userId,
	})
	tokenString, err := token.SignedString([]byte(TOKEN_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(TOKEN_KEY), nil
	})
	if err != nil {
		return -1, err
	}
	claims := token.Claims.(jwt.MapClaims)
	err = claims.Valid()
	if err != nil {
		return -1, fmt.Errorf("invalid token")
	}

	return int64(claims["userId"].(float64)), nil
}
