package handlers

import (
	"dab/own/global/microservices/autentication/pkg/domain/interfaces/services"
	"dab/own/global/microservices/autentication/scripts"

	"github.com/gin-gonic/gin"
)

type SettingsController struct {
	services services.ISettingsService
}

func NewSettingsController(services services.ISettingsService) *SettingsController {
	return &SettingsController{
		services: services,
	}
}

func (controller *SettingsController) InitialiceService(context *gin.Context) {
	if scripts.IsSettingsSetteds() {
		context.IndentedJSON(404, gin.H{
			"message": "Not found",
		})
	} else {
		context.File("./static/index.html")
	}

}
