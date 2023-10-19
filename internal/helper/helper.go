package helper

import (
	"Bank-INA/domain"
	"Bank-INA/internal/errorhandler"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateUser(data domain.User) (err domain.ErrorData) {
	errData := validation.ValidateStruct(&data,
		validation.Field(&data.Email, validation.Required, is.Email),
	)

	if errData != nil {
		err = errorhandler.Failed(domain.FailedValidationControllCode, domain.FailedValidationControllCode, errData)
	}

	return
}
