package main

import (
	"flag"
	"net/http"
)

const (
	assetsFolder   = "assets"
	articlesFolder = "assets/articles"
	pagesFolder    = "pages"

	page404       = "404.html"
	defaultLayout = "layout.html"
)

var (
	root, port string
	config     *SiteConfig
	mux        *Mux
	appVersion string
)

func init() {
	flag.StringVar(&root, "root", "", "Site root directory (parent directory of 'assets', 'pages', etc")
	flag.StringVar(&port, "port", ":3000", "http server port")
	flag.Parse()

	config = NewSiteConfig(root)

	appVersion = version()

	router := &Router{
		assetsServer: http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))),
	}
	router.add("/", indexHandler)
	router.add("/blog", blogHandler)
	router.add("/blog/:slug", articleHandler)
	router.add("/subscribe", subscribeHandler)

	mux = &Mux{router, nil}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	})
}
