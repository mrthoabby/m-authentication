package types

import (
	"com.github/mrthoabby/m-authentication/interfaces"
	"com.github/mrthoabby/m-authentication/util"
	"github.com/gin-gonic/gin"
)

type JSONBinder[T any] struct{}

func (json *JSONBinder[T]) Bind(context *gin.Context, data *T) error {
	return context.ShouldBindJSON(data)
}

type XMLBinder[T any] struct{}

func (b *XMLBinder[T]) Bind(context *gin.Context, data *T) error {
	return context.ShouldBindXML(data)
}

type FORMBinder[T any] struct{}

func (b *FORMBinder[T]) Bind(context *gin.Context, data *T) error {
	return context.ShouldBind(data)
}

type BinderStrategy[T any] struct {
	Strategy interfaces.BindFormatStrategy[T]
}

func (b *BinderStrategy[T]) SetStrategy(strategy interfaces.BindFormatStrategy[T]) {
	b.Strategy = strategy
}

func (b *BinderStrategy[T]) Bind(context *gin.Context, data *T) error {
	if b.Strategy == nil {
		util.LoggerHandler().Error("Error binding data", "error", "Strategy not set")
	}
	return b.Strategy.Bind(context, data)
}
