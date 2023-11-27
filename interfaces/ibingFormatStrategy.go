package interfaces

import (
	"github.com/gin-gonic/gin"
)

type IBindFormatStrategy[T any] interface {
	Bind(context *gin.Context, data *T) error
}
