package main

import "net/http"

func init() {
	LoggingSettings(Logfile)
}

func StartMainServer() error {
	http.HandleFunc("/todos", HandleIndex)
	return http.ListenAndServe(":8080", nil)
}

func main() {
	StartMainServer()
}
