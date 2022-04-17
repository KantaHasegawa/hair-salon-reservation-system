package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSetting(logFile string) {
	logfile, err := os.OpenFile(
		logFile,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Println(err)
	}
	multLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multLogFile)
}
