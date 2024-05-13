package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type Document struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func GetDocuments(f fs.FS) ([]Document, error) {
	var documents []Document
	err := fs.WalkDir(f, "docs", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		ext := filepath.Ext(d.Name())
		if strings.ToLower(ext) != ".md" {
			return nil
		}

		contents, err := fs.ReadFile(f, path)
		if err != nil {
			return fmt.Errorf("error reading %s: %w", path, err)
		}

		document := Document{path, string(contents)}
		documents = append(documents, document)
		return nil
	})

	if err != nil {
		return documents, err
	}

	return documents, nil
}
