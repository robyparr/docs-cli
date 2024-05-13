package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DocumentCreationRequest struct {
	Documents []Document `json:"documents"`
}

func PushDocuments(url string, docs []Document) error {
	jsonStr, err := json.Marshal(DocumentCreationRequest{docs})
	if err != nil {
		return fmt.Errorf("error encoding API request body: %w", err)
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return fmt.Errorf("error making API request: %w", err)
	}

	return nil
}
