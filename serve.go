package main

import (
	"flag"
	"log"
	"net/http"
)

const (
	assetsFolder   = "assets"
	articlesFolder = "assets/articles"
	pagesFolder    = "pages"

	page_404      = "404.html"
	defaultLayout = "layout.html"
)

var root, port string
var CONFIG *siteConfig

func init() {
	flag.StringVar(&root, "root", "", "Site root directory (parent directory of 'assets', 'pages', etc")
	flag.StringVar(&port, "port", ":3000", "http server port")
}

func main() {
	flag.Parse()

	CONFIG = NewSiteConfig(root)

	router := &Router{
		assetsServer: http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))),
	}

	router.add("/", indexHandler)
	router.add("/blog", blogHandler)
	router.add("/blog/:slug", articleHandler)

	err := http.ListenAndServe(port, Mux{router, &SimpleLogger{}})
	if err != nil {
		log.Fatalf("Cannot start server on port: %q", port)
	}
}
