package main

import (
	"fmt"
	"os"
)

func main() {
	orgSlug, foundEnv := os.LookupEnv("ORG_SLUG")
	if !foundEnv {
		fmt.Fprintln(os.Stdout, "An organization must be specified with the ORG_SLUG environment variable.")
		os.Exit(1)
	}

	f := os.DirFS(".")
	documents, err := GetDocuments(f)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Error loading documents: %s", err)
		os.Exit(1)
	}

	url := fmt.Sprintf("http://%s.localhost:3000/api/documents", orgSlug)
	fmt.Printf("Pushing %d documents to %s", len(documents), url)
	PushDocuments(url, documents)
}
