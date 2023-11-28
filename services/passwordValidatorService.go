package services

import (
	"crypto/sha256"
	"encoding/hex"

	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/types/basic"
)

type PasswordValidatorService struct {
	encrypt  basic.Encrypt
	password string
}

func NewPasswordValidatorService(encrypt basic.Encrypt, password string) *PasswordValidatorService {
	return &PasswordValidatorService{
		encrypt:  encrypt,
		password: password,
	}
}

func (p *PasswordValidatorService) IsAnValidPassword(hash string) bool {
	switch p.encrypt.Algorithm {
	case globalConfig.ALGORITHM_ENCRYPTION_TYPE_HS256:
		return hash == p.NewSHA256()
	default:
		return false
	}
}

func (p *PasswordValidatorService) NewSHA256() string {
	hashCreator := sha256.New()
	hashCreator.Write([]byte(p.password))
	return hex.EncodeToString(hashCreator.Sum(nil))
}
