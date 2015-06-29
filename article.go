package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

var TitleRegex = regexp.MustCompile(`^([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])_([a-zA-Z_]+)-([^.]+)`)

// Article is the model of each posts
type Article struct {
	Title        string
	Tags         []string
	Date         time.Time
	ShortContent string
	FullContent  string

	src string
}

// Articles is a collection of Article
type Articles []*Article

// to void loading articles on every request
type contentCache struct {
	articles Articles
}

var cache = new(contentCache)

// Loads all articles in srcDir
func LoadArticles() Articles {
	if len(cache.articles) > 0 {
		return cache.articles
	}
	var srcs []string
	filepath.Walk(CONFIG.articlesDir, func(path string, info os.FileInfo, err error) error {
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

	cache.articles = articles
	return articles
}

func LoadArticle(src string, skipContent bool) *Article {
	a := &Article{src: src}
	a.Date, a.Tags, a.Title = parseFilename(src)
	if !skipContent {
		full := string(readMarkdowns(src)[0])
		k := len(full)
		if len(full) > 255 {
			k = 255
		}
		a.ShortContent = full[:k]
		a.FullContent = full
	}
	return a
}

func FindArticle(slug string) *Article {
	title := strings.Replace(slug, "-", " ", -1)
	var articles Articles
	if len(cache.articles) > 0 {
		articles = cache.articles
	} else {
		articles = LoadArticles()
	}
	for _, art := range articles {
		if strings.ToLower(art.Title) == title {
			return art
		}
	}
	return nil
}

func GetArchiveList() (byDate []time.Time, byTag tagSet) {
	var articles Articles
	if len(cache.articles) > 0 {
		articles = cache.articles
	} else {
		articles = LoadArticles()
	}
	byDate = make([]time.Time, len(articles))
	for i, art := range articles {
		byDate[i] = art.Date
		for _, t := range art.Tags {
			byTag.add(t)
		}
	}
	return
}

func (arts Articles) Len() int           { return len(arts) }
func (arts Articles) Swap(i, j int)      { arts[i], arts[j] = arts[j], arts[i] }
func (arts Articles) Less(i, j int) bool { return arts[i].Date.Before(arts[j].Date) }
