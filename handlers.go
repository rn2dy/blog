package main

import (
	"net/http"
	"path/filepath"
	"time"

  "google.golang.org/appengine"
  "google.golang.org/appengine/datastore"
)

// Handler a adapter for my own convenient
type Handler func(*C, http.ResponseWriter, *http.Request)

func (h Handler) serveHTTP(c *C, w http.ResponseWriter, r *http.Request) {
	h(c, w, r)
}

func indexHandler(_ *C, w http.ResponseWriter, r *http.Request) {
	w.Write(renderPage("index.html", struct {
		Title      string
		SkipFooter bool
		Version    string
	}{"", true, appVersion}))
}

func blogHandler(c *C, w http.ResponseWriter, r *http.Request) {
	articles := LoadArticles()
	w.Write(renderPage("articles.html", struct {
		Title      string
		Articles   Articles
		SkipFooter bool
		Version    string
	}{"Articles", articles, true, appVersion}))
}

func articleHandler(c *C, w http.ResponseWriter, r *http.Request) {
	if _, ok := c.vars["slug"]; !ok {
		http.Redirect(w, r, filepath.Join(config.pagesDir, page404), http.StatusNotFound)
		return
	}
	article := FindArticle(c.vars["slug"])
	if article == nil {
		http.Redirect(w, r, filepath.Join(config.pagesDir, page404), http.StatusNotFound)
		return
	}
	byDate, byTag := GetArchiveList()
	w.Write(renderPage("article.html", struct {
		Title         string
		Article       *Article
		ArchiveByDate []time.Time
		ArchiveByTag  tagSet
		SkipFooter    bool
		Version       string
	}{article.Title, article, byDate, byTag, true, appVersion}))
}

func subscribeHandler(c *C, w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	r.ParseForm()

	if r.Form.Get("email") != "" {
		http.Error(w, "Subscriber must provide an email.", http.StatusBadRequest)
		return
	}

	sub := Subscriber{r.Form.Get("name"), r.Form.Get("email"), time.Now()}

	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Sub", subscribersKey(ctx)), &sub)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
