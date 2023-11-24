package middlewares

import (
	"time"

	"com.github/mrthoabby/m-authentication/util"
	"github.com/gin-gonic/gin"
)

func RequestTracer() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentTime := time.Now()
		util.LoggerHandler().Info("Request received", "method", c.Request.Method, "path", c.Request.URL.Path, "time", currentTime.Format("2006-01-02 15:04:05"))

		c.Next()

		latencyTime := time.Since(currentTime)
		util.LoggerHandler().Info("Request finished", "method", c.Request.Method, "path", c.Request.URL.Path, "time", currentTime.Format("2006-01-02 15:04:05"), "latency", latencyTime)
	}

}
