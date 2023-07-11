package routers

import (
	"dab/own/global/microservices/autentication/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func DefineSettingsRoutes(router *gin.Engine, controller *handlers.SettingsController) {
	router.GET("/settings", controller.InitialiceService)
}
