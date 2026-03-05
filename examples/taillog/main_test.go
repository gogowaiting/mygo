package main

import (
	"net/http/httptest"
	"testing"
)

func TestParseOrigins(t *testing.T) {
	origins := parseOrigins("http://a.com, http://b.com ,,")
	if len(origins) != 2 {
		t.Fatalf("expected 2 origins, got %d", len(origins))
	}
	if _, ok := origins["http://a.com"]; !ok {
		t.Fatal("missing origin http://a.com")
	}
	if _, ok := origins["http://b.com"]; !ok {
		t.Fatal("missing origin http://b.com")
	}
}

func TestIsOriginAllowed(t *testing.T) {
	allowed := map[string]struct{}{"http://localhost:3000": {}}
	if !isOriginAllowed("http://localhost:3000", allowed) {
		t.Fatal("expected origin to be allowed")
	}
	if isOriginAllowed("http://evil.com", allowed) {
		t.Fatal("expected origin to be denied")
	}
}

func TestIsAuthorized(t *testing.T) {
	cfg := config{token: "secret"}
	req := httptest.NewRequest("GET", "http://localhost/ws/logs", nil)
	if isAuthorized(req, cfg) {
		t.Fatal("expected unauthorized request")
	}

	req = httptest.NewRequest("GET", "http://localhost/ws/logs?token=secret", nil)
	if !isAuthorized(req, cfg) {
		t.Fatal("expected authorized by query token")
	}

	req = httptest.NewRequest("GET", "http://localhost/ws/logs", nil)
	req.Header.Set("Authorization", "Bearer secret")
	if !isAuthorized(req, cfg) {
		t.Fatal("expected authorized by bearer token")
	}
}
