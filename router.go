package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

var pathVarReg = regexp.MustCompile(`(:[^?/()]+)`)

// Matcher my custom path matcher
type Matcher struct {
	names   []string
	reg     *regexp.Regexp
	handler Handler
}

func (m Matcher) matchPath(path string) (map[string]string, bool) {
	matches := m.reg.FindAllStringSubmatch(path, -1)
	if len(matches) != 1 {
		return nil, false
	}
	mm := matches[0][1:]
	if len(mm) != len(m.names) {
		return nil, false
	}
	vars := make(map[string]string)
	for i, v := range mm {
		vars[m.names[i]] = v
	}
	return vars, true
}

// Router register matchers from a pattern and also does static file serving
type Router struct {
	assetsServer http.Handler
	matchers     []Matcher
}

func (rter *Router) match(c *C, r *http.Request) (interface{}, error) {
	// check static assets request first
	if strings.HasPrefix(noRootSlash(r.URL.Path), noRootSlash(assetsFolder)) {
		log.Printf("Serve static file from %q.", r.URL.Path)
		return rter.assetsServer, nil
	}
	for _, m := range rter.matchers {
		if vars, ok := m.matchPath(r.URL.Path); ok {
			for k, v := range vars {
				c.vars[k] = v
			}
			return m.handler, nil
		}
	}
	return nil, fmt.Errorf("Unmatched route: %q\n", r.URL.Path)
}

func (rter *Router) route(c *C, w http.ResponseWriter, r *http.Request) {
	handler, err := rter.match(c, r)
	if err != nil {
		http.Redirect(w, r, filepath.Join(config.pagesDir, page404), http.StatusNotFound)
		return
	}
	switch h := handler.(type) {
	case http.Handler:
		h.ServeHTTP(w, r)
	case Handler:
		h.serveHTTP(c, w, r)
	default:
		log.Fatalf("Unknown handler type %v", h)
	}
}

func (rter *Router) add(path string, h Handler) {
	index := pathVarReg.FindAllStringIndex(path, -1)
	var reg = ""
	if len(index) == 0 {
		reg = fmt.Sprintf("%s", path)
	}
	for i, j := 0, 0; i < len(index); j, i = i, i+1 {
		if i == 0 {
			reg = reg + fmt.Sprintf("%s([^/?]+)", path[:index[i][0]])
		} else {
			reg = reg + fmt.Sprintf("%s([^/?]+)", path[index[j][1]:index[i][0]])
		}
	}
	namesMatches := pathVarReg.FindAllStringSubmatch(path, -1)
	var names = make([]string, len(namesMatches))
	for i, n := range namesMatches {
		names[i] = n[0][1:]
	}
	rter.matchers = append(rter.matchers, Matcher{names, regexp.MustCompile(reg + "$"), h})
}
