package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HelloHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var resp Response
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expectedMessage := "Hello, World!"
	if resp.Message != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, resp.Message)
	}
}
