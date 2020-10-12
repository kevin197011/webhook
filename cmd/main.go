package main

import (
	"log"
	"os"
	"webhook/config"
	_ "webhook/daemon"
	"webhook/http"
	"webhook/utils"
)

var (
	logFile *os.File
	err     error
)

func main() {
	logPath := config.NewConfig().LogPath
	logFile, err = utils.LogSetup(logPath, logFile)
	if err != nil {
		log.Println(err)
		return
	}
	defer logFile.Close()
	http.Run()
	log.Println("start webhook server...")
}
