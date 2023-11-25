package util

import (
	"github.com/go-playground/validator/v10"
)

var instance *validator.Validate

func GetValidator() *validator.Validate {

	if instance == nil {
		instance = validator.New(validator.WithRequiredStructEnabled())
	}
	return instance
}
