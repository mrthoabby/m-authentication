package infraestructure

import (
	"dab/own/global/microservices/autentication/pkg/api/handlers"
	"dab/own/global/microservices/autentication/pkg/api/routers"
	"dab/own/global/microservices/autentication/pkg/infraestructure/services"

	"github.com/gin-gonic/gin"
)

func InitialiceDependencyInjection(router *gin.Engine) {
	injectSettingsController(router)
}

func injectSettingsController(router *gin.Engine) {
	service := &services.SettingsServices{}
	controller := handlers.NewSettingsController(service)
	routers.DefineSettingsRoutes(router, controller)
}
