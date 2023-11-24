package helpers

import (
	"com.github/mrthoabby/m-authentication/types"
	"github.com/gin-gonic/gin"
)

// Binder is a function that binds the request body to a given type.
// It takes the content type, the context, and a pointer to the result type as parameters.
// The function returns an error if the binding fails.
func Binder[R any](contentType string, context *gin.Context, result *R) error {

	var binder types.BinderStrategy[R]

	switch contentType {
	case "application/json":
		binder.SetStrategy(&types.JSONBinder[R]{})
	case "application/xml":
		binder.SetStrategy(&types.XMLBinder[R]{})
	case "application/x-www-form-urlencoded":
		binder.SetStrategy(&types.FORMBinder[R]{})
	default:
		return types.NewCustomError("Content type not supported")
	}

	if errorBinding := binder.Bind(context, result); errorBinding != nil {
		return errorBinding
	}
	return nil
}
