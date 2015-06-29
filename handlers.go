package main

import (
	"net/http"
	"path/filepath"
	"time"
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
	articles := LoadArticles()
	w.Write(renderPage("articles.html", struct {
		Title      string
		Articles   Articles
		SkipFooter bool
	}{"Articles", articles, true}))
}

func articleHandler(c *C, w http.ResponseWriter, r *http.Request) {
	if _, ok := c.vars["slug"]; !ok {
		http.Redirect(w, r, filepath.Join(CONFIG.pagesDir, page_404), http.StatusNotFound)
		return
	}
	article := FindArticle(c.vars["slug"])
	byDate, byTag := GetArchiveList()
	w.Write(renderPage("article.html", struct {
		Title         string
		Article       *Article
		ArchiveByDate []time.Time
		ArchiveByTag  tagSet
		SkipFooter    bool
	}{article.Title, article, byDate, byTag, true}))
}
