// +build !appengine

package main

import (
	"net/http"
	"os"
)

func main() {
	envport := os.Getenv("PORT")
	mux.logger = &SimpleLogger{}
	if envport != "" {
		port = envport
	}
	http.ListenAndServe(port, mux)
}
