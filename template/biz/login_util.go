package biz

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	TOKEN_KEY = "GOCODEGEN20220808"
)

func CreatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 相同密码原文，每次调用产生的字符串都不同，需要使用CompareHashAndPassword才能判断是否相同
	if err != nil {
		return password, err
	}

	return string(hash), nil
}

// passwordHashInDb: pasword hash str stored in db
// passwordLogin: password from login input
func IsPasswordRight(passwordHashInDb string, passwordLogin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashInDb), []byte(passwordLogin))
	if err != nil {
		return false
	} else {
		return true
	}
}

func CreateToken(tokenPairs map[string]interface{}) (string, error) {
	mapClaims := jwt.MapClaims{}
	for k, v := range tokenPairs {
		mapClaims[k] = v
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	token, err := t.SignedString([]byte(TOKEN_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeToken(tokenStr string) (map[string]interface{}, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
		}
		return []byte(TOKEN_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		tokenPairs := map[string]interface{}{}
		for k, v := range claims {
			tokenPairs[k] = v
		}

		return tokenPairs, nil
	} else {
		return nil, errors.New("token验证失败")
	}
}
