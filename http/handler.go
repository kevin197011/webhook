package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sort"
	"webhook/config"
	"webhook/providers"
	"webhook/utils"
)

func healthzHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": "healthz",
	})
}

func whatsappAlertsHandler(c *gin.Context) {
	var messages HookMessage
	var emoji string
	cf := config.NewConfig()

	if err := c.Bind(&messages); err != nil {
		log.Printf("Bind json data err: %s \n", err)
		return
	}

	log.Printf("Bind json data: %#v \n", messages)

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
			log.Printf("Time format err: %v\n", err)
		}
		msg += fmt.Sprintf("*TimeAt:* %s", timeVal)
		log.Printf("send alert data: %v\n", msg)
		providers.SendMsg(providers.NewWhatsappOpt("123456", cf.GroupName), msg)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": "success!",
	})
}
