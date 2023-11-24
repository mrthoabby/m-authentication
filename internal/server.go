package internal

import (
	"net/http"
	"strconv"
	"time"

	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/internal/middlewares"
	"com.github/mrthoabby/m-authentication/internal/modules/basicAuth"
	"com.github/mrthoabby/m-authentication/types/settings"
	"github.com/gin-gonic/gin"
)

func ServerInit(configs *settings.Config) {
	_router := gin.Default()
	_router.Use(middlewares.RequestTracer())

	_server := &http.Server{
		Addr:           ":" + strconv.Itoa(configs.Server.Port),
		Handler:        _router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if configs.Service.AuthMethod.Type == globalConfig.AUTH_METHOD_BASIC {
		basicAuth.AuthController(_router.Group("/auth"))
	}

	_server.ListenAndServe()
}
