package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webhook/config"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/healthz", healthzHandler)
	v1.POST("whatsapp", whatsappAlertsHandler)
	return r
}

func Run() {
	_ = Router().Run(fmt.Sprintf(":%s", config.NewConfig().Port))
}
