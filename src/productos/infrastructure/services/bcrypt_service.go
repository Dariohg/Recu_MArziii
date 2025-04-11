package services

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct{}

func NewBcryptService() *BcryptService {
	return &BcryptService{}
}

func (bs *BcryptService) Encrypt(plainText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (bs *BcryptService) Verify(hashedText, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))
	return err == nil
}
