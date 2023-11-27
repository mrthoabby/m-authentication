package basicAuth

import (
	"com.github/mrthoabby/m-authentication/helpers"
	"com.github/mrthoabby/m-authentication/internal/databases"
	"com.github/mrthoabby/m-authentication/types"
	"com.github/mrthoabby/m-authentication/util"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	respository  databases.IDatabaseConnectionRepository
	reouterGroup *gin.RouterGroup
}

func NewAuthController(respository databases.IDatabaseConnectionRepository, reouterGroup *gin.RouterGroup) *AuthController {
	return &AuthController{
		respository:  respository,
		reouterGroup: reouterGroup,
	}
}

func (controller *AuthController) Register() {
	controller.reouterGroup.POST("/login", controller.login)
}

func (controller *AuthController) login(context *gin.Context) {
	var crendentials types.Credentials
	contentType := context.Request.Header.Get("Content-Type")
	if errorGettingData := helpers.Binder[types.Credentials](contentType, context, &crendentials); errorGettingData != nil {
		util.LoggerHandler().Error("Error getting data", "error", errorGettingData.Error())
		context.JSON(400, gin.H{"message": errorGettingData.Error()})
		return
	}
	controller.respository.ValidAuthentication(crendentials)

	context.JSON(200, gin.H{"message": "login success"})
}
