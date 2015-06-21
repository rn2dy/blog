package main

import (
	"log"
	"net/http"
)

type Logger interface {
	log(*http.Request)
}
type SimpleLogger struct{}

func (l *SimpleLogger) log(r *http.Request) {
	log.Printf("Request: %q", r.URL.Path)
}
