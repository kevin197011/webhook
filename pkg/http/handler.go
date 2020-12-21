package http

import (
	"fmt"
	"net/http"
	"sort"
	"webhook/pkg/config"
	"webhook/pkg/providers"
	"webhook/pkg/utils"

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
		zap.L().Error("Bind json data fail", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "bad",
			"data": "failed!",
		})
		return
	}
	zap.L().Info("Bind json data success", zap.Any("data", messages))

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
		if message, ok := alert.Annotations["message"]; ok {
			msg += fmt.Sprintf("*注解:*\\n  %s\\n", message)
		} else {
			msg += fmt.Sprintf("*注解:*\\n  %s\\n", alert.Annotations["summary"])
		}
		timeVal, err := utils.TimeFormat(alert.StartsAt)
		if err != nil {
			zap.L().Error("Time format error!", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "bad",
				"data": "failed!",
			})
			return
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

func whatsappAlertsHandlerW(c *gin.Context) {
	var messages HookMessage
	var emoji string

	if err := c.Bind(&messages); err != nil {
		zap.L().Error("Bind json data fail", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "bad",
			"data": "failed!",
		})
		return
	}
	zap.L().Info("Bind json data success", zap.Any("data", messages))

	if messages.Status == "firing" {
		emoji = "❌"
	} else {
		emoji = "✅"
	}

	for _, alert := range messages.Alerts {
		var msg string
		msg = fmt.Sprintf("*状态:* %s %s\\n", messages.Status, emoji)
		msg += "*标签:*\\n"
		keys := make([]string, 0, len(alert.Labels))
		for k := range alert.Labels {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			msg += fmt.Sprintf("  %s = %s\\n", key, alert.Labels[key])
		}
		if message, ok := alert.Annotations["message"]; ok {
			msg += fmt.Sprintf("*注解:*\\n  %s\\n", message)
		} else {
			msg += fmt.Sprintf("*注解:*\\n  %s\\n", alert.Annotations["summary"])
		}
		timeVal, err := utils.TimeFormat(alert.StartsAt)
		if err != nil {
			zap.L().Error("Time format error!", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "bad",
				"data": "failed!",
			})
			return
		}
		msg += fmt.Sprintf("*时间:* %s", timeVal)
		zap.L().Info("send alert data.", zap.Any("msg", msg))
		providers.SendMsg(providers.NewWhatsappOpt("123456", config.NewConfig().GroupNameW), msg)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": "success!",
	})
}

func whatsappAlertsHandlerB(c *gin.Context) {
	var messages HookMessage
	var emoji string

	if err := c.Bind(&messages); err != nil {
		zap.L().Error("Bind json data fail", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "bad",
			"data": "failed!",
		})
		return
	}
	zap.L().Info("Bind json data success", zap.Any("data", messages))

	if messages.Status == "firing" {
		emoji = "❌"
	} else {
		emoji = "✅"
	}

	for _, alert := range messages.Alerts {
		var msg string
		msg = fmt.Sprintf("*状态:* %s %s\\n", messages.Status, emoji)
		msg += "*标签:*\\n"
		keys := make([]string, 0, len(alert.Labels))
		for k := range alert.Labels {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			msg += fmt.Sprintf("  %s = %s\\n", key, alert.Labels[key])
		}
		if message, ok := alert.Annotations["message"]; ok {
			msg += fmt.Sprintf("*注解:*\\n  %s\\n", message)
		} else {
			msg += fmt.Sprintf("*注解:*\\n  %s\\n", alert.Annotations["summary"])
		}
		timeVal, err := utils.TimeFormat(alert.StartsAt)
		if err != nil {
			zap.L().Error("Time format error!", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "bad",
				"data": "failed!",
			})
			return
		}
		msg += fmt.Sprintf("*时间:* %s", timeVal)
		zap.L().Info("send alert data.", zap.Any("msg", msg))
		providers.SendMsg(providers.NewWhatsappOpt("123456", config.NewConfig().GroupNameB), msg)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": "success!",
	})
}


func whatsappAlertsHandlerC(c *gin.Context) {
	var messages HookMessage
	var emoji string

	if err := c.Bind(&messages); err != nil {
		zap.L().Error("Bind json data fail", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "bad",
			"data": "failed!",
		})
		return
	}
	zap.L().Info("Bind json data success", zap.Any("data", messages))

	if messages.Status == "firing" {
		emoji = "❌"
	} else {
		emoji = "✅"
	}

	for _, alert := range messages.Alerts {
		var msg string
		msg = fmt.Sprintf("*状态:* %s %s\\n", messages.Status, emoji)
		msg += "*标签:*\\n"
		keys := make([]string, 0, len(alert.Labels))
		for k := range alert.Labels {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			msg += fmt.Sprintf("  %s = %s\\n", key, alert.Labels[key])
		}
		if message, ok := alert.Annotations["message"]; ok {
			msg += fmt.Sprintf("*注解:*\\n  %s\\n", message)
		} else {
			msg += fmt.Sprintf("*注解:*\\n  %s\\n", alert.Annotations["summary"])
		}
		timeVal, err := utils.TimeFormat(alert.StartsAt)
		if err != nil {
			zap.L().Error("Time format error!", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "bad",
				"data": "failed!",
			})
			return
		}
		msg += fmt.Sprintf("*时间:* %s", timeVal)
		zap.L().Info("send alert data.", zap.Any("msg", msg))
		providers.SendMsg(providers.NewWhatsappOpt("123456", config.NewConfig().GroupNameC), msg)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": "success!",
	})
}

