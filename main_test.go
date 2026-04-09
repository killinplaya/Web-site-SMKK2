package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleCompanyReturnsJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/company", nil)
	rr := httptest.NewRecorder()

	handleCompany(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Fatalf("expected JSON content type, got %q", ct)
	}
	if !strings.Contains(rr.Body.String(), "\"legal_name\"") {
		t.Fatalf("expected response to include legal_name, body=%q", rr.Body.String())
	}
}

func TestHandleCompanyMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/company", strings.NewReader("{}"))
	rr := httptest.NewRecorder()

	handleCompany(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rr.Code)
	}
}

func TestSecurityHeadersMiddleware(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	handler := withSecurityHeaders(next)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	expectedHeaders := []string{
		"X-Content-Type-Options",
		"X-Frame-Options",
		"Referrer-Policy",
		"Content-Security-Policy",
		"Permissions-Policy",
		"Cross-Origin-Opener-Policy",
		"Cross-Origin-Resource-Policy",
		"X-Permitted-Cross-Domain-Policies",
	}

	for _, h := range expectedHeaders {
		if rr.Header().Get(h) == "" {
			t.Fatalf("expected header %s to be set", h)
		}
	}
}
