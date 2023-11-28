package internal

import (
	"net/http"
	"reflect"
	"strconv"
	"time"

	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/internal/databases"
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

	switch currentType := authConfig.GetType(); {
	case configSettings.Service.AuthMethod.Type == globalConfig.AUTH_METHOD_BASIC && reflect.TypeOf(basic.Config{}) == currentType:
		authenticationSettings, isOk := authConfig.(*basic.Config)
		if !isOk {
			util.LoggerHandler().Error("Error validating basic auth config", "error", "Invalid auth method")
			return
		}
		var routerName string
		if routerName = authenticationSettings.Auth.RouterName; routerName[0:1] != "/" {
			routerName = "/" + routerName
		}
		databaseFactory := databases.NewConcreteDatabaseFactory(authenticationSettings.Connection[0])
		pooler := databases.NewDatabasePooler(databaseFactory, &types.TableMapper{
			AuthTable:        authenticationSettings.Auth.Table.Name,
			UserColumn:       authenticationSettings.Auth.Table.User.Column,
			PasswordColumn:   authenticationSettings.Auth.Table.Password.Column,
			DataSourceTables: make([]string, 0),         //Empty for test
			DataSourcColumns: make(map[string][]string), //Empty for test
		})
		if pooler.GetConnectionsOpened() < globalConfig.POOL_SIZE_DATABASE_CONNECTION*0.8 {
			util.LoggerHandler().Error("Error creating database connection", "error", "Not enough connections")
			return
		}
		basicAuth.NewAuthController(pooler.GetConnection(), _router.Group(routerName), *authenticationSettings).Register()

	default:
		util.LoggerHandler().Error("Error validating basic auth config", "error", "Invalid auth method")
	}

	_server.ListenAndServe()
}
