package main

import (
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func noRootSlash(path string) string {
	if strings.HasPrefix(path, "/") {
		return path[1:]
	}
	return path
}

func articlePath(dir, name string) (string, bool) {
	path := filepath.Join(dir, name)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return path, false
		}
	}
	return path, true
}

func recentArticles(dir string) []string {
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return readMarkdowns(files...)
}

func readMarkdowns(files ...string) []string {
	var all []string
	for _, path := range files {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("Cannot read article from %q.", path)
		}
		all = append(all, string(blackfriday.MarkdownCommon(data)))
	}
	return all
}

func makeDate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}
