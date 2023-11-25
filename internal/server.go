package internal

import (
	"net/http"
	"reflect"
	"strconv"
	"time"

	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/internal/middlewares"
	"com.github/mrthoabby/m-authentication/internal/modules/basicAuth"
	"com.github/mrthoabby/m-authentication/types"
	"com.github/mrthoabby/m-authentication/types/basic"
	"com.github/mrthoabby/m-authentication/util"
	"github.com/gin-gonic/gin"
)

func ServerInit() {
	_configs := types.NewConfigBuilder()
	configSettings := _configs.GetConfigs()
	errorBuildingConfig, authConfig := _configs.BuildAuthConfig()
	if errorBuildingConfig != nil {
		util.LoggerHandler().Error("Error building config", "error", errorBuildingConfig.Error())
	}

	_router := gin.Default()
	_router.Use(middlewares.RequestTracer())

	_server := &http.Server{
		Addr:           ":" + strconv.Itoa(configSettings.Server.Port),
		Handler:        _router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	auth, isOk := authConfig.(*basic.Config)
	switch currentType := authConfig.GetType(); {
	case reflect.TypeOf(basic.Config{}) == currentType && isOk && configSettings.Service.AuthMethod.Type == globalConfig.AUTH_METHOD_BASIC:

		var routerName string
		if routerName = auth.Auth.RouterName; routerName[0:1] != "/" {
			routerName = "/" + routerName
		}
		basicAuth.AuthController(_router.Group(routerName))
	default:
		util.LoggerHandler().Error("Error validating basic auth config", "error", "Invalid auth method")
	}

	_server.ListenAndServe()
}
