package server

import (
	"gatewayH/internal/gateway/services"
	"gatewayH/pkg/middlewares"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// initRouter 初始化路由
func initRouter(s *services.Services) *gin.Engine {
	if s.Cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(middlewares.Cors())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	return router
}
