package basicAuth

import (
	"com.github/mrthoabby/m-authentication/helpers"
	"com.github/mrthoabby/m-authentication/internal/databases"
	"com.github/mrthoabby/m-authentication/services"
	"com.github/mrthoabby/m-authentication/types"
	"com.github/mrthoabby/m-authentication/types/basic"
	"com.github/mrthoabby/m-authentication/util"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	respository  databases.IDatabaseConnectionRepository
	reouterGroup *gin.RouterGroup
	authSettings basic.Config
}

func NewAuthController(respository databases.IDatabaseConnectionRepository, reouterGroup *gin.RouterGroup, authSettings basic.Config) *AuthController {
	return &AuthController{
		respository:  respository,
		reouterGroup: reouterGroup,
		authSettings: authSettings,
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
		context.JSON(400, gin.H{"message": "login failed"})
		return
	}
	passwordHash, errorGettingHash := controller.respository.GetPasswordHash(crendentials)
	if errorGettingHash != nil {
		util.LoggerHandler().Error("Error getting password hash", "error", errorGettingHash.Error())
		context.JSON(404, gin.H{"message": "login failed"})
		return
	}

	passWordService := services.NewPasswordValidatorService(controller.authSettings.Auth.Table.Password.Encrypt, crendentials.Password)
	if passWordService.IsAnValidPassword(passwordHash) {
		context.JSON(200, gin.H{"message": "login success"})
	} else {
		context.JSON(401, gin.H{"message": "login failed"})
	}
	return
}
