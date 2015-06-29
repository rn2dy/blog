package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
)

var TitleRegex = regexp.MustCompile(`^([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])_([a-zA-Z_]+)-([^.]+)`)

// Article is the model of each posts
type Article struct {
	Title   string
	Tags    []string
	Date    time.Time
	Content string

	src string
}

// Articles is a collection of Article
type Articles []*Article

// Loads all articles in srcDir
func LoadArticles(srcDir string) Articles {
	var srcs []string
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			srcs = append(srcs, path)
		}
		return nil
	})
	articles := make([]*Article, len(srcs))
	for i, src := range srcs {
		articles[i] = LoadArticle(src, false)
	}
	sort.Sort(Articles(articles))

	return articles
}

func LoadArticle(src string, skipContent bool) *Article {
	a := &Article{src: src}
	a.Date, a.Tags, a.Title = parseFilename(src)
	if !skipContent {
		a.Content = string(readMarkdowns(src)[0])
	}
	return a
}

func (arts Articles) Len() int           { return len(arts) }
func (arts Articles) Swap(i, j int)      { arts[i], arts[j] = arts[j], arts[i] }
func (arts Articles) Less(i, j int) bool { return arts[i].Date.Before(arts[j].Date) }
