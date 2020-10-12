package utils

import (
	"io"
	"log"
	"os"
)

func LogSetup(logPath string, logFile *os.File) (*os.File, error) {
	var err error
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	return logFile, nil
}
