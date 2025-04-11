package core

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptText(plainText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
