package main

import (
  "log"
  "net/http"
  tpl "html/template"
  "os"
  "path/filepath"
  "flag"
  "bytes"
  "io/ioutil"
)

var (
  root = flag.String("root", "", "Website root (parent of 'static', 'articles', and 'html")
  port = flag.String("port", ":3000", "http server port")

  layoutHtml, errorHtml *tpl.Template
)

func main() {
  flag.Parse()

  // check if root is a valid directory
  if *root == "" {
    var err error
    *root, err = os.Getwd()
    if err != nil {
      log.Fatalf("Failed to getwd: %v", err)
    }
  }

  // handlers
  mux := http.DefaultServeMux
  mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(*root, "static")))))
	mux.Handle("/favicon.ico", http.FileServer(http.Dir(filepath.Join(*root, "static"))))
  mux.HandleFunc("/blog", blogHandler)
  mux.HandleFunc("/projects", projectHandler)
  mux.HandleFunc("/show", showHandler)
  mux.HandleFunc("/error", errorHandler)
  mux.HandleFunc("/", mainHandler)

  // start server
  errc := make(chan error)
  go func() {
    errc <- http.ListenAndServe(*port, nil)
  }()
  log.Fatalf("Serve error: %v", <-errc)
}

func mainHandler(rw http.ResponseWriter, req *http.Request) {
  if len(req.URL.Path) > 1 {
    log.Printf("Unsupported request: %q\n", req.URL.Path)
    http.Redirect(rw, req, "/error", http.StatusFound)
    return
  }
  data := struct {
    Title string
    SkipFooter bool
  }{
    "",
    true,
  }
  servePage(rw, "index.html", data)
}

func blogHandler(rw http.ResponseWriter, req *http.Request) {
  rw.WriteHeader(200)
}

func projectHandler(rw http.ResponseWriter, req *http.Request) {
  rw.WriteHeader(200)
}

func showHandler(rw http.ResponseWriter, req *http.Request) {
  rw.WriteHeader(200)
}

func errorHandler(rw http.ResponseWriter, req *http.Request) {
  html, err := ioutil.ReadFile(filepath.Join(*root, "html", "error.html"))
  if err != nil {
    log.Println(err)
  }
  rw.Write(html)
}

func servePage(rw http.ResponseWriter, pageName string, data interface{}) {
  content := renderPage(pageName, data)
  rw.Write(content)
}

func renderPage(pageName string, data interface{}) []byte {
  var buf bytes.Buffer

  contentPath := filepath.Join(*root, "html", pageName)
  layoutPath := filepath.Join(*root, "html", "layout.html")

  t, err := tpl.ParseFiles(layoutPath, contentPath)
  if err != nil {
    log.Printf("ParseFile: %s, %s", layoutPath, contentPath)
  }

  if err := t.Execute(&buf, data); err != nil {
    log.Printf("%s.Execute: %s", t.Name(), err)
  }

  return buf.Bytes()
}
