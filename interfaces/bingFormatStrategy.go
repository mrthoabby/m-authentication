package interfaces

import (
	"github.com/gin-gonic/gin"
)

type BindFormatStrategy[T any] interface {
	Bind(context *gin.Context, data *T) error
}
