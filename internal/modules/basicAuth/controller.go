package basicAuth

import (
	"com.github/mrthoabby/m-authentication/helpers"
	"com.github/mrthoabby/m-authentication/util"
	"github.com/gin-gonic/gin"
)

func AuthController(group *gin.RouterGroup) {
	{
		group.POST("/", login)
	}
}

func login(context *gin.Context) {
	var crendentials Credentials
	contentType := context.Request.Header.Get("Content-Type")
	if errorGettingData := helpers.Binder[Credentials](contentType, context, &crendentials); errorGettingData != nil {
		util.LoggerHandler().Error("Error getting data", "error", errorGettingData.Error())
		context.JSON(400, gin.H{"message": errorGettingData.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "login success"})
}
