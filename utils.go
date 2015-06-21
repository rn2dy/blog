package main

import (
	"os"
	"path/filepath"
	"strings"
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
