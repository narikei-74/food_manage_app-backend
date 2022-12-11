package logFile

import (
	"os"
	"log"
	"time"
)

func logStart() {
	t := time.Now()
	logPath := "./logs/" + t.Year() + "/" + t.Month() + "/" + t.Day() + ".go"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	log.SetOutput(file)
}