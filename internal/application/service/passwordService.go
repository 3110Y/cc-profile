package service

import (
	"crypto/sha512"
	"encoding/hex"
)

type PasswordService struct {
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (p PasswordService) Encode(password string) (passwordHash string, err error) {
	passwordSHA := sha512.Sum512([]byte(password))
	return hex.EncodeToString(passwordSHA[:]), nil
}
