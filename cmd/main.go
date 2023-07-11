package main

import (
	"dab/own/global/microservices/autentication/pkg/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	router := ServerSetup()
	BootstrapSetup(router)
	router.Run(":3000")
}

func ServerSetup() *gin.Engine {
	router := gin.Default()

	return router
}

func BootstrapSetup(router *gin.Engine) {
	infraestructure.InitialiceDependencyInjection(router)
}
