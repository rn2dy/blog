package main

import (
	"net/http"
)

// Mux custom http multiplexer
type Mux struct {
	router *Router
	logger Logger
}

// C http request context to store free variables
type C struct {
	vars map[string]string
}

func (m Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &C{make(map[string]string)}
	if m.logger != nil {
		m.logger.log(r)
	}
	m.router.route(c, w, r)
}
