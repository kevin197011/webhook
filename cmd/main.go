package main

import (
	"io"
	"log"
	"os"
	"webhook/config"
	_ "webhook/daemon"
	"webhook/http"
)

func main() {
	logPath := config.NewConfig().LogPath
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.Println("start webhook server...")
	http.InitRouter()
}
