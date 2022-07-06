package dal

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

var (
	QueryCtx = context.Background()
)

func Md5(str []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(str, bcrypt.DefaultCost)
	if err != nil {
		return string(str), err
	}

	return string(hash), nil
}
