package http

import (
	"fmt"
	"net/http"
	"sort"
	"webhook/config"
	"webhook/providers"
	"webhook/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func healthzHandler(c *gin.Context) {
	zap.L().Info("webhook interface is healthz!")

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": "healthz",
	})
}

func whatsappAlertsHandler(c *gin.Context) {
	var messages HookMessage
	var emoji string

	if err := c.Bind(&messages); err != nil {
		zap.L().Error("Bind json data fail!", zap.Error(err))
		return
	}
	zap.L().Info("Bind json data sucess.", zap.Any("data", messages))

	if messages.Status == "firing" {
		emoji = "❌"
	} else {
		emoji = "✅"
	}

	for _, alert := range messages.Alerts {
		var msg string
		msg = fmt.Sprintf("*Status:* %s %s\\n", messages.Status, emoji)
		msg += "*Labels:*\\n"
		keys := make([]string, 0, len(alert.Labels))
		for k := range alert.Labels {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			msg += fmt.Sprintf("  %s = %s\\n", key, alert.Labels[key])
		}
		msg += fmt.Sprintf("*Annotations:*\\n  %s\\n", alert.Annotations["message"])
		timeVal, err := utils.TimeFormat(alert.StartsAt)
		if err != nil {
			zap.L().Error("Time format error!", zap.Error(err))
		}
		msg += fmt.Sprintf("*TimeAt:* %s", timeVal)
		zap.L().Info("send alert data.", zap.Any("msg", msg))
		providers.SendMsg(providers.NewWhatsappOpt("123456", config.NewConfig().GroupName), msg)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": "success!",
	})
}
