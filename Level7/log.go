package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var Logfile = "err.log"

func LoggingSettings(logfilePath string) {
	file, err := os.OpenFile(Logfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return

	MultiLogfile := io.MultiWriter(os.Stderr, file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(MultiLogfile)
}
