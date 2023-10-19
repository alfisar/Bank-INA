package hashing

import (
	"Bank-INA/domain"
	"Bank-INA/internal/errorhandler"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePass(data string) (string, domain.ErrorData) {
	pwd := []byte(data)
	pwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", errorhandler.Failed(domain.FailedHashingPassCode, domain.FailedHashingPass, err)
	}
	response := string(pwd[:])
	return response, domain.ErrorData{}
}
