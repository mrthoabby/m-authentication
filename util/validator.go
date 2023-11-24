package util

import (
	"com.github/mrthoabby/m-authentication/globalConfig"
	"github.com/go-playground/validator/v10"
)

var instance *validator.Validate

func GetValidator() *validator.Validate {
	globalConfig.OneTime.Do(func() {
		instance = validator.New(validator.WithRequiredStructEnabled())
	})
	return instance
}
