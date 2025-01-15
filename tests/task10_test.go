package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Router(t *testing.T) {
	t.Parallel()
	router := NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL + "/ping")
	if err != nil {
		t.Fatalf("Error during request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Execpeted Content-Type application/json, got %s", contentType)
	}
	var pong Pong
	if err := json.NewDecoder(resp.Body).Decode(&pong); err != nil {
		t.Fatalf("Error decode %v", err)
	}
	if pong.Message != "Pong" {
		t.Errorf("Expected pong, got %s", pong.Message)
	}
}
