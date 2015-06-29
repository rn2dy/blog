package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/russross/blackfriday"
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

func readMarkdowns(files ...string) [][]byte {
	var all = make([][]byte, len(files))
	for i, path := range files {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("Cannot read article from %q.", path)
		}
		all[i] = blackfriday.MarkdownCommon(data)
	}
	return all
}

func parseFilename(src string) (date time.Time, tags []string, title string) {
	name := filepath.Base(src)
	m := TitleRegex.FindAllStringSubmatch(name, -1)
	if len(m) == 0 {
		log.Fatalf("Can not parse filename: %q", name)
	}
	if len(m[0]) != 6 {
		log.Fatalf("Filename not formatted correctly: %q. Got: %v", name, m)
	}
	mm := m[0]
	year, _ := strconv.Atoi(mm[1])
	mon, _ := strconv.Atoi(mm[2])
	day, _ := strconv.Atoi(mm[3])

	date = makeDate(year, mon, day)
	tags = strings.Split(mm[4], "_")
	title = strings.Join(strings.Split(mm[5], "-"), " ")
	return
}

func makeDate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

func simpleDate(date time.Time) string {
	return date.Format("01/02/2006")
}

func toLink(title string) string {
	return strings.Replace(strings.ToLower(title), " ", "-", -1)
}
