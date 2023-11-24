package settings

import (
	globalConfig "com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/util"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Service service `xml:"service" validate:"required"`
	Server  server  `xml:"server" validate:"required"`
}

type service struct {
	AuthMethod authMethod `xml:"authMethod" validate:"required"`
}

type authMethod struct {
	Output string `xml:"output,attr" validate:"validOutputAuthMethod"`
	Type   string `xml:"type,attr" validate:"validTypeAuthMethod"`
}

func isAnValidOutputAuthMethod(fl validator.FieldLevel) bool {
	validOutputs := globalConfig.CurrentOutputTypes
	for _, validOutput := range validOutputs {
		if fl.Field().String() == validOutput {
			return true
		}
	}
	return false
}

func isAnValidTypeAuthMethod(fl validator.FieldLevel) bool {
	validTypes := globalConfig.CurrentAuthMethods
	for _, validType := range validTypes {
		if fl.Field().String() == validType {
			return true
		}
	}
	return false
}

type server struct {
	Port int `xml:"port"`
}

func init() {
	validator := util.GetValidator()
	globalConfig.OneTime.Do(func() {
		validator.RegisterValidation("validOutputAuthMethod", isAnValidOutputAuthMethod)
		validator.RegisterValidation("validTypeAuthMethod", isAnValidTypeAuthMethod)
	})
}
