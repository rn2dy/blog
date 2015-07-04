// +build !appengine

package main

import "net/http"

func main() {
	mux.logger = &SimpleLogger{}
	http.ListenAndServe(port, mux)
}
