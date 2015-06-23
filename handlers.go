package main

import (
	"net/http"
	"path/filepath"
)

type Handler func(*C, http.ResponseWriter, *http.Request)

func (h Handler) serveHTTP(c *C, w http.ResponseWriter, r *http.Request) {
	h(c, w, r)
}

func indexHandler(_ *C, w http.ResponseWriter, r *http.Request) {
	w.Write(renderPage("index.html", struct {
		Title      string
		SkipFooter bool
	}{"", true}))
}

func blogHandler(c *C, w http.ResponseWriter, r *http.Request) {
	// get most recent blog entries
	articles := recentArticles(site.articlesDir)
	w.Write(renderPage("article.html", struct {
		Articles []string
	}{articles}))
}

func articleHandler(c *C, w http.ResponseWriter, r *http.Request) {
	if _, ok := c.vars["slug"]; !ok {
		http.Redirect(w, r, filepath.Join(site.pagesDir, page_404), http.StatusNotFound)
		return
	}
}
