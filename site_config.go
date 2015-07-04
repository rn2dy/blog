package main

import (
	"log"
	"os"
	"path/filepath"
)

// SiteConfig holds global configs, mostly paths
type SiteConfig struct {
	rootDir, assetsDir, pagesDir, articlesDir string
}

// NewSiteConfig initialize global site config
func NewSiteConfig(root string) *SiteConfig {
	s := &SiteConfig{}
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
