package main

import (
	"go.uber.org/zap"
	_ "webhook/daemon"
	"webhook/http"
)

var Logger *zap.Logger

func main() {
	http.Run()
	Logger.Info("WebHook Server Startup...")
}
