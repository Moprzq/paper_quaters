package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoCacheHeaders(t *testing.T) {
	handler := noCache(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	request := httptest.NewRequest(http.MethodGet, "/paper_quarters.wasm", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if got := response.Header().Get("Cache-Control"); got != "no-store, no-cache, must-revalidate, max-age=0" {
		t.Fatalf("Cache-Control = %q", got)
	}
	if got := response.Header().Get("Pragma"); got != "no-cache" {
		t.Fatalf("Pragma = %q", got)
	}
	if got := response.Header().Get("Expires"); got != "0" {
		t.Fatalf("Expires = %q", got)
	}
}
