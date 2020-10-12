package http

import (
	"fmt"
	"time"
	"webhook/pkg/config"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(zap.L(), true))
	v1 := engine.Group("/v1")
	v1.GET("/healthz", healthzHandler)
	v1.POST("whatsapp", whatsappAlertsHandler)
	v1.POST("whatsappw", whatsappAlertsHandlerW)
	v1.POST("whatsappb", whatsappAlertsHandlerB)
	return engine
}

func Run() error {
	zap.L().Info("webhook server statup ...")
	return Router().Run(fmt.Sprintf(":%s", config.NewConfig().Port))
}
