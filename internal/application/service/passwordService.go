package service

import (
	"crypto/sha512"
	"encoding/hex"
	utlits "github.com/3110Y/cc-utlits"
)

type PasswordService struct {
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (p PasswordService) Encode(password string) (passwordHash *string, err error) {
	passwordSHA := sha512.Sum512([]byte(password))
	return utlits.Pointer(hex.EncodeToString(passwordSHA[:])), nil
}
