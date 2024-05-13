package main

import (
	"reflect"
	"sort"
	"testing"
	"testing/fstest"
)

func TestGetDocuments(t *testing.T) {
	fs := fstest.MapFS{
		"docs/first.md":       {Data: []byte("# First")},
		"docs/second.md":      {Data: []byte("# Second")},
		"docs/nested/file.md": {Data: []byte("# Nested")},
		"README.MD":           {Data: []byte("# Readme")},
	}
	docs, err := GetDocuments(fs)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	wantDocs := []Document{
		{"docs/first.md", "# First"},
		{"docs/second.md", "# Second"},
		{"docs/nested/file.md", "# Nested"},
	}

	sort.Slice(wantDocs, func(i, j int) bool {
		return wantDocs[i].Path < wantDocs[j].Path
	})
	if !reflect.DeepEqual(docs, wantDocs) {
		t.Fatalf("Want %+v, got %+v", wantDocs, docs)
	}
}
