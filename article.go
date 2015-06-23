package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var TitleRegex = regexp.MustCompile(`^([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])_([a-zA-Z_]+)-([^.]+)`)

type Article struct {
	Title   string
	Tags    []string
	Date    time.Time
	Content []byte

	src string
}

func NewArticle(src string) *Article {
	a := &Article{src: src}
	a.parseName()
	return a
}

func (a *Article) parseName() {
	m := TitleRegex.FindAllStringSubmatch(a.src, -1)
	if len(m) == 0 {
		log.Fatalf("Can not parse filename: %q", a.src)
	}
	if len(m[0]) != 6 {
		log.Fatalf("Filename not formatted correctly: %q. Got: %v", a.src, m)
	}
	mm := m[0]
	year, _ := strconv.Atoi(mm[1])
	mon, _ := strconv.Atoi(mm[2])
	day, _ := strconv.Atoi(mm[3])

	a.Date = makeDate(year, mon, day)
	a.Tags = strings.Split(mm[4], "_")
	a.Title = strings.Join(strings.Split(mm[5], "-"), " ")
}
