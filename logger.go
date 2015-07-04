package main

import (
	"log"
	"net/http"
)

// Logger custom logger
type Logger interface {
	log(*http.Request)
}

// SimpleLogger default logger
type SimpleLogger struct{}

func (l *SimpleLogger) log(r *http.Request) {
	log.Printf("Request: %q", r.URL.Path)
}
