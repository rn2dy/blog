package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	tpl "text/template"
)

const (
	assetsFolder   = "assets"
	articlesFolder = "assets/articles"
	pagesFolder    = "pages"

	page_404      = "404.html"
	defaultLayout = "layout.html"
)

var root, port string
var site *Site

func init() {
	flag.StringVar(&root, "root", "", "Website root (parent of 'assets', 'pages")
	flag.StringVar(&port, "port", ":3000", "http server port")
}

type Site struct {
	assetsDir, pagesDir, articlesDir string
}

func (s Site) initialize() *Site {
	// check root directory
	if root == "" {
		_, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to getwd: %v", err)
		}
	}
	log.Printf("Site root is: %q\n", root)

	s.assetsDir = filepath.Join(root, assetsFolder)
	s.pagesDir = filepath.Join(root, pagesFolder)
	s.articlesDir = filepath.Join(root, articlesFolder)
	return &s
}

func main() {
	flag.Parse()
	site = Site{}.initialize()

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

func renderPage(pageName string, data interface{}) []byte {
	var buf bytes.Buffer
	contentPath := filepath.Join(site.pagesDir, pageName)
	layoutPath := filepath.Join(site.pagesDir, defaultLayout)

	t, err := tpl.ParseFiles(layoutPath, contentPath)
	if err != nil {
		log.Printf("ParseFile: %s, %s", layoutPath, contentPath)
	}
	if err := t.Execute(&buf, data); err != nil {
		log.Printf("%s.Execute: %s", t.Name(), err)
	}
	return buf.Bytes()
}
