package settings

import (
	"com.github/mrthoabby/m-authentication/global"
	"com.github/mrthoabby/m-authentication/utils"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Service Service `xml:"service" validate:"required"`
	Server  Server  `xml:"server" validate:"required"`
}

type Service struct {
	AuthMethod AuthMethod `xml:"authMethod" validate:"required"`
}

type AuthMethod struct {
	Output string `xml:"output,attr" validate:"validOutputAuthMethod"`
	Type   string `xml:"type,attr" validate:"validTypeAuthMethod"`
}

func isAnValidOutputAuthMethod(fl validator.FieldLevel) bool {
	validOutputs := global.CurrentOutputTypes
	for _, validOutput := range validOutputs {
		if fl.Field().String() == validOutput {
			return true
		}
	}
	return false
}

func isAnValidTypeAuthMethod(fl validator.FieldLevel) bool {
	validTypes := global.CurrentAuthMethods
	for _, validType := range validTypes {
		if fl.Field().String() == validType {
			return true
		}
	}
	return false
}

type Server struct {
	Port int `xml:"port"`
}

func init() {
	validator := utils.GetValidator()
	global.Once.Do(func() {
		validator.RegisterValidation("validOutputAuthMethod", isAnValidOutputAuthMethod)
		validator.RegisterValidation("validTypeAuthMethod", isAnValidTypeAuthMethod)
	})
}
