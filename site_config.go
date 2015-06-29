package main

import (
	"log"
	"os"
	"path/filepath"
)

type siteConfig struct {
	rootDir, assetsDir, pagesDir, articlesDir string
}

func NewSiteConfig(root string) *siteConfig {
	s := &siteConfig{}
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
	return s
}
