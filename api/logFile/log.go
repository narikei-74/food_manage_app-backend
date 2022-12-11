package logFile

import (
	"os"
	"log"
	"time"
)

func LogStart() *os.File{
	t := time.Now()
	s := t.Format("2006-01-02")
	logPath := "./logFile/logs/" + s + ".log"
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}