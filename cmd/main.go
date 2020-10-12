package main

import (
	"webhook/pkg/config"
	_ "webhook/pkg/daemon"
	"webhook/pkg/http"
	"webhook/pkg/utils"

	"go.uber.org/zap"
)

func init() {
	utils.SetLogs(zap.DebugLevel, utils.LOGFORMAT_CONSOLE, config.NewConfig().LogPath)
}

func main() {
	if err := http.Run(); err != nil {
		zap.L().Fatal("webhook server start fail!", zap.Error(err))
	}
}
