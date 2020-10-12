package http

import (
	"fmt"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"webhook/config"
)

func Router() *gin.Engine {
	r := gin.New()
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	v1 := r.Group("/v1")
	v1.GET("/healthz", healthzHandler)
	v1.POST("whatsapp", whatsappAlertsHandler)
	return r
}

func Run() {
	_ = Router().Run(fmt.Sprintf(":%s", config.NewConfig().Port))
}
