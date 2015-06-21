package main

import (
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
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

func articleHandler(c *C, w http.ResponseWriter, r *http.Request) {
	if _, ok := c.vars["slug"]; !ok {
		http.Redirect(w, r, filepath.Join(site.pagesDir, page_404), http.StatusNotFound)
		return
	}
	var title = c.vars["slug"]

	path, found := articlePath(site.articlesDir, title+".md")
	if !found {
		http.Redirect(w, r, filepath.Join(site.pagesDir, page_404), http.StatusNotFound)
		return
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot read article from %q.", path)
	}
	out := blackfriday.MarkdownCommon(data)
	w.Write(renderPage("article.html", struct {
		Title      string
		Content    string
		SkipFooter bool
	}{title, string(out), false}))
}
