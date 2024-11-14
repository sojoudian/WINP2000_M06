package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileServer(t *testing.T) {
	fs := http.FileServer(http.Dir("./templates"))
	ts := httptest.NewServer(fs)
	defer ts.Close()
	url := ts.URL + "/"

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 Ok, got: %v", err)
	}

}
