package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPushDocuments(t *testing.T) {
	docs := []Document{
		{"docs/first.md", "# First"},
		{"docs/second.md", "# Second"},
	}

	var receivedRequests [][]string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Unexpected error reading request body: %s", err)
		}
		r.Body.Close()

		receivedRequests = append(receivedRequests, []string{r.Method, string(body)})
		w.WriteHeader(http.StatusOK)
	}))

	PushDocuments(server.URL, docs)

	wantRequests := [][]string{
		{"POST", `{"documents":[{"path":"docs/first.md","content":"# First"},{"path":"docs/second.md","content":"# Second"}]}`},
	}

	if !reflect.DeepEqual(receivedRequests, wantRequests) {
		t.Errorf("want requests %+v, got %+v", wantRequests, receivedRequests)
	}
}
